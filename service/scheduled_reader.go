package service

import (
	"context"
	"log"
	"promotions/model"

	"github.com/procyon-projects/chrono"
)

type PromotionSaver interface {
	Save(model.Promotion)
}

type ScheduledReader struct {
	saver PromotionSaver
}

func GetScheduledReader(saver PromotionSaver) ScheduledReader {
	return ScheduledReader{
		saver: saver,
	}
}

// TODO: job config
func (reader ScheduledReader) ScheduleJob() {
	taskScheduler := chrono.NewDefaultTaskScheduler()
	taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		log.Print("One-Shot Task")
	}, 10000)
}
