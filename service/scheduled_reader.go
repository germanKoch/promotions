package service

import (
	"context"
	"math"
	"promotions/config"
	"promotions/model"
	"promotions/service/storage"
	"time"

	"github.com/procyon-projects/chrono"
)

type PromotionProcessor interface {
	UpsertAll(promotions []model.Promotion)
}

type PromotionParser interface {
	Parse(csvString string) model.Promotion
}

type HistoryProcessor interface {
	GetAfter(processedAfter time.Time) []model.ProcessedFile
	Save(processedFile model.ProcessedFile)
}

type Storage interface {
	Walk(process func(file storage.FileData))
}

type ScheduledReader struct {
	config          config.SchedulerConfig
	promotions      PromotionProcessor
	history         HistoryProcessor
	storage         Storage
	promotionParser PromotionParser
}

func GetScheduledReader(schedulerConfig config.SchedulerConfig, saver PromotionProcessor, history HistoryProcessor, promotionParser PromotionParser, storage Storage) ScheduledReader {
	return ScheduledReader{
		config:          schedulerConfig,
		promotions:      saver,
		storage:         storage,
		history:         history,
		promotionParser: promotionParser,
	}
}

func (reader ScheduledReader) ScheduleJob() {
	taskScheduler := chrono.NewDefaultTaskScheduler()
	taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		reader.importPromotions()
	}, time.Duration(reader.config.Period*int64(math.Pow(10, 6))))
}

func (reader ScheduledReader) importPromotions() {
	daysBefore := reader.config.DaysDelta
	batchSize := reader.config.BatchSize

	now := time.Now().UTC()
	timeAfter := now.AddDate(0, 0, -daysBefore)

	importedFiles := reader.history.GetAfter(timeAfter)
	pathToProcessedFile := make(map[string]model.ProcessedFile)

	for i := 0; i < len(importedFiles); i += 1 {
		pathToProcessedFile[importedFiles[i].Path] = importedFiles[i]
	}

	reader.storage.Walk(func(file storage.FileData) {
		modificationDate := file.ModificationDate().UTC()
		if modificationDate.After(timeAfter) {
			_, exists := pathToProcessedFile[file.Path()]
			if !exists {
				reader.importSingleFile(file, batchSize)
			}
		}
	})
}

func (reader ScheduledReader) importSingleFile(file storage.FileData, batchSize int) {
	liner := file.Content()
	batchPointer := 0
	batch := make([]model.Promotion, batchSize)

	for i := 0; liner.ReadNext(); i++ {
		line := liner.NextLine()

		if len(line) > 0 {
			batch[batchPointer] = reader.promotionParser.Parse(line)
			batchPointer += 1
		}

		if batchPointer > 0 && batchPointer%batchSize == 0 {
			reader.promotions.UpsertAll(batch)
			batchPointer = 0
			batch = make([]model.Promotion, batchSize)
		}
	}

	if batchPointer != 0 {
		reader.promotions.UpsertAll(batch[:batchPointer])
	}

	reader.history.Save(
		model.ProcessedFile{
			Path:           file.Path(),
			ProcessingDate: time.Now().UTC(),
		},
	)

	defer liner.Close()
}
