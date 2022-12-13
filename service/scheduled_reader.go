package service

import (
	"context"
	"log"
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

func GetScheduledReader(saver PromotionProcessor, history HistoryProcessor, promotionParser PromotionParser) ScheduledReader {
	return ScheduledReader{
		promotions:      saver,
		history:         history,
		promotionParser: promotionParser,
	}
}

// TODO: job config
func (reader ScheduledReader) ScheduleJob() {
	taskScheduler := chrono.NewDefaultTaskScheduler()
	taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		log.Print("One-Shot Task")
	}, 10000)
}

func (reader ScheduledReader) ImportPromotions() {
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
		if file.ModificationDate().UTC().After(timeAfter) {
			_, exists := pathToProcessedFile[file.Path()]
			if !exists {
				reader.importSingleFile(file, batchSize)
			}
		}
	})
}

func (reader ScheduledReader) importSingleFile(file storage.FileData, batchSize int) {
	batch := make([]model.Promotion, batchSize)
	liner := file.Content()

	for i := 0; liner.HasNext(); i++ {
		for j := 0; j < batchSize && liner.HasNext(); j++ {
			batch[j] = reader.promotionParser.Parse(liner.NextLine())
		}
		reader.promotions.UpsertAll(batch)
	}

	reader.history.Save(
		model.ProcessedFile{
			Path:           file.Path(),
			ProcessingDate: time.Now().UTC(),
		},
	)

	defer liner.Close()
}
