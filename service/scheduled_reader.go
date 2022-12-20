package service

import (
	"context"
	"log"
	"math"
	"promotions/config"
	"promotions/model"
	"time"

	"github.com/procyon-projects/chrono"
)

type PromotionProcessor interface {
	UpsertAll(promotions []model.Promotion)
}

type PromotionParser interface {
	Parse(csvString string) (model.Promotion, error)
}

type HistoryProcessor interface {
	GetAfter(processedAfter time.Time) []model.ProcessedFile
	Save(processedFile model.ProcessedFile)
}

type Storage interface {
	Walk(process func(file model.FileData))
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
	timeAfter := now.AddDate(0, 0, -daysBefore-1)

	importedFiles := reader.history.GetAfter(timeAfter)
	pathToProcessedFile := make(map[string]model.ProcessedFile)

	for i := 0; i < len(importedFiles); i += 1 {
		pathToProcessedFile[importedFiles[i].Path] = importedFiles[i]
	}

	reader.storage.Walk(func(file model.FileData) {
		modificationDate := file.ModificationDate().UTC()
		if modificationDate.After(timeAfter) {
			_, exists := pathToProcessedFile[file.Path()]
			if !exists {
				reader.importSingleFile(file, batchSize)
			}
		}
	})
}

func (reader ScheduledReader) importSingleFile(file model.FileData, batchSize int) {
	log.Printf("importing %s file", file.Path())

	liner, err := file.Content()
	if err != nil {
		log.Printf("File [%s] could not be imported. Error: %s", file.Path(), err.Error())
		return
	}
	defer liner.Close()

	reader.importFileContent(liner, batchSize)

	reader.history.Save(
		model.ProcessedFile{
			Path:           file.Path(),
			ProcessingDate: time.Now().UTC(),
		},
	)
}

func (reader ScheduledReader) importFileContent(liner model.FileLiner, batchSize int) {
	batchPointer := 0
	batch := make([]model.Promotion, batchSize)

	for i := 0; liner.ReadNext(); i++ {
		line := liner.NextLine()

		if len(line) > 0 {
			promotion, err := reader.promotionParser.Parse(line)
			if err == nil {
				batch[batchPointer] = promotion
				batchPointer += 1
			} else {
				log.Printf("line [%s] could not be parsed. Error: %s", line, err.Error())
			}
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
}
