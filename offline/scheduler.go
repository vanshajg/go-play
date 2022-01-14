package offline

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/vanshajg/go-play/container"
)

func Init(container container.Container) {
	s := gocron.NewScheduler(time.UTC)
	logger := container.GetLogger()
	logger.GetZapLogger().Infof("starting scheduled jobs")
	s.Every(5).Seconds().Do(func() {
		logger.GetZapLogger().Infof("scheduled log")
	})

	fetchComments(container)

	s.StartAsync()
}
