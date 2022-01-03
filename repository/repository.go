package repository

import (
	"database/sql"
	"fmt"

	"github.com/vanshajg/go-play/config"
	"github.com/vanshajg/go-play/logger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	Close() error
	Transaction(fc func(r Repository) error) (err error)
	Select(query interface{}, args ...interface{}) *gorm.DB
	AutoMigrate(interface{}) error
	DropTableIfExists(value interface{}) error
	ScanRows(rows *sql.Rows, result interface{}) error
	Raw(sql string, values ...interface{}) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

type commentRepository struct {
	*repository
}

func NewCommentRepository(logger *logger.Logger, config *config.Config) Repository {
	logger.GetZapLogger().Infof("Setting up connection to DB")
	db, err := connectDatabase(logger, config)
	if err != nil {
		logger.GetZapLogger().Errorf("cannot connect to database. %s", err)
	}
	logger.GetZapLogger().Infof("successfully connected to database %s:%s", config.Database.Host, config.Database.Port)
	return &commentRepository{&repository{db: db}}
}

const (
	SQLITE   = "sqlite3"
	POSTGRES = "postgres"
)

func connectDatabase(logger *logger.Logger, config *config.Config) (*gorm.DB, error) {
	var dsn string
	gormConfig := &gorm.Config{Logger: logger}

	if config.Database.Dialect == POSTGRES {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Dbname, config.Database.Password)
		return gorm.Open(postgres.Open(dsn), gormConfig)
	}
	return gorm.Open(sqlite.Open(config.Database.Host), gormConfig)
}

func (rep *repository) Close() error {
	sqlDB, _ := rep.db.DB()
	return sqlDB.Close()
}

func (rep *repository) Transaction(fc func(r Repository) error) (err error) {
	panicked := true
	txn := rep.db.Begin()
	defer func() {
		if panicked || err != nil {
			txn.Rollback()
		}
	}()

	txnrep := &repository{}
	txnrep.db = txn
	err = fc(txnrep)
	if err == nil {
		err = txn.Commit().Error
	}
	panicked = false
	return
}

// Select specify fields that you want to retrieve from database when querying, by default, will select all fields;
func (rep *repository) Select(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Select(query, args...)
}

// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
func (rep *repository) AutoMigrate(value interface{}) error {
	return rep.db.AutoMigrate(value)
}

// DropTableIfExists drop table if it is exist
func (rep *repository) DropTableIfExists(value interface{}) error {
	return rep.db.Migrator().DropTable(value)
}

func (rep *repository) ScanRows(rows *sql.Rows, result interface{}) error {
	return rep.db.ScanRows(rows, result)
}

func (rep *repository) Raw(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Raw(sql, values...)
}
