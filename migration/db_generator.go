package migration

import (
	"github.com/vanshajg/go-play/container"
	"github.com/vanshajg/go-play/models"
)

func CreateDatabase(container container.Container) {
	if !container.GetConfig().Database.Migration {
		return
	}
	db := container.GetRepository()
	var err error
	err = db.DropTableIfExists(&models.Comment{})
	if err != nil {
		container.GetLogger().GetZapLogger().Errorf(err.Error())
	}
	err = db.AutoMigrate(&models.Comment{})
	if err != nil {
		container.GetLogger().GetZapLogger().Errorf(err.Error())
	}

	container.GetLogger().GetZapLogger().Infof("Migration complete")
}
