package service

import (
	"context"
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
	promotions      PromotionProcessor
	history         HistoryProcessor
	storage         Storage
	promotionParser PromotionParser
}

func GetScheduledReader(saver PromotionProcessor, history HistoryProcessor, promotionParser PromotionParser, storage Storage) ScheduledReader {
	return ScheduledReader{
		promotions:      saver,
		storage:         storage,
		history:         history,
		promotionParser: promotionParser,
	}
}

// TODO: job config
func (reader ScheduledReader) ScheduleJob() {
	taskScheduler := chrono.NewDefaultTaskScheduler()
	taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		reader.importPromotions()
	}, 100000)
}

func (reader ScheduledReader) importPromotions() {
	daysBefore := 7
	batchSize := 100

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

	for i := 0; liner.HasNext(); i++ {
		batchPointer := 0
		batch := make([]model.Promotion, batchSize)

		for j := 0; batchPointer < batchSize && liner.HasNext(); j++ {
			line := liner.NextLine()
			if len(line) > 0 {
				//TODO error processing
				batch[batchPointer] = reader.promotionParser.Parse(line)
				batchPointer += 1
			}
		}
		if len(batch) != 0 {
			reader.promotions.UpsertAll(batch)
		}
	}

	reader.history.Save(
		model.ProcessedFile{
			Path:           file.Path(),
			ProcessingDate: time.Now().UTC(),
		},
	)

	defer liner.Close()
}
