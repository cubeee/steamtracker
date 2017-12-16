package config

import (
	"fmt"

	"github.com/cubeee/steamtracker-updater/shared/db"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rakanalh/scheduler/storage"
)

type TaskAttributes struct {
	Id          int64 `orm:"auto;index"`
	Name        string
	Params      string
	Duration    string
	LastRun     string `gorm:"column:last_run"`
	NextRun     string `gorm:"column:next_run"`
	IsRecurring string `gorm:"column:is_recurring"`
	Hash        string
}

func (TaskAttributes) TableName() string {
	return "task_store"
}

type PostgresConfig struct {
	DbName string
}

type PostgresStorage struct {
	storage.TaskStore
	details db.ConnectDetails
	db      *gorm.DB
}

func NewPostgresStorage(connectDetails db.ConnectDetails) PostgresStorage {
	return PostgresStorage{
		details: connectDetails,
	}
}

func (postgres *PostgresStorage) Connect() error {
	details := postgres.details
	args := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", details.Host, details.Db, details.User, details.Pass)
	if details.Additional != "" {
		args = args + " " + details.Additional
	}
	gormDb, err := gorm.Open("postgres", args)
	if err != nil {
		return err
	}
	postgres.db = gormDb
	return nil
}

func (postgres *PostgresStorage) Close() error {
	return postgres.db.Close()
}

func (postgres *PostgresStorage) Initialize() error {
	postgres.db.AutoMigrate(&TaskAttributes{})
	return postgres.db.Error
}

func (postgres PostgresStorage) Add(attr storage.TaskAttributes) error {
	task := postgres.createTask(attr)
	errors := postgres.db.Where("hash = ?", task.Hash).
		FirstOrCreate(&task).
		GetErrors()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func (postgres PostgresStorage) Remove(attr storage.TaskAttributes) error {
	postgres.db.Delete(TaskAttributes{}, "hash = ?", attr.Hash)
	err := postgres.db.Error
	if err != nil {
		return fmt.Errorf("Error while deleting task: %s", err)
	}
	return nil
}

func (postgres PostgresStorage) Fetch() ([]storage.TaskAttributes, error) {
	var tasks []TaskAttributes
	postgres.db.Find(&tasks)
	return postgres.createTaskAttributes(tasks), postgres.db.Error
}

func (postgres *PostgresStorage) insert(task TaskAttributes) error {
	postgres.db.Create(&task)
	err := postgres.db.Error
	if err != nil {
		return fmt.Errorf("Error while inserting task: %s", err)
	}
	return nil
}

func (postgres *PostgresStorage) createTask(attr storage.TaskAttributes) TaskAttributes {
	return TaskAttributes{
		Hash:        attr.Hash,
		Name:        attr.Name,
		LastRun:     attr.LastRun,
		NextRun:     attr.NextRun,
		Duration:    attr.Duration,
		IsRecurring: attr.IsRecurring,
		Params:      attr.Params,
	}
}

func (postgres *PostgresStorage) createTaskAttribute(task TaskAttributes) storage.TaskAttributes {
	return storage.TaskAttributes{
		Hash:        task.Hash,
		Name:        task.Name,
		LastRun:     task.LastRun,
		NextRun:     task.NextRun,
		Duration:    task.Duration,
		IsRecurring: task.IsRecurring,
		Params:      task.Params,
	}
}

func (postgres *PostgresStorage) createTaskAttributes(tasks []TaskAttributes) []storage.TaskAttributes {
	var attrs []storage.TaskAttributes
	for _, task := range tasks {
		attrs = append(attrs, postgres.createTaskAttribute(task))
	}
	return attrs
}
