package database

import (
	"errors"
	"github.com/spf13/viper"
	"go-telebot-init/pkg/database/dbservice"
	"go-telebot-init/pkg/database/models"
	"go-telebot-init/pkg/helpers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
)

type DBImpl struct {
	DBS dbservice.DBService
}

func NewDB(db *gorm.DB) *DBImpl {
	return &DBImpl{
		DBS: dbservice.NewDBService(db),
	}
}

func Init() (*DBImpl, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = setupDatabase(db)
	if err != nil {
		return nil, err
	}

	// Delete or comment in production
	db = db.Debug()

	database := NewDB(db)
	return database, nil
}

func setupDatabase(db *gorm.DB) error {
	if err := migrateIfNeeded(db); err != nil {
		return err
	}

	if err := createAdminIfNotExists(db); err != nil {
		return err
	}

	slog.Info("Database is ready")
	return nil
}

func migrateIfNeeded(db *gorm.DB) error {
	if pingDatabase(db) && healthCheck(db) {
		return nil
	}

	slog.Info("Auto-migrating database")
	return db.AutoMigrate(&models.User{})
}

func healthCheck(db *gorm.DB) bool {
	tables := map[string]interface{}{
		"Users": &models.User{},
	}

	var missingTables []string
	allTablesExist := true

	for tableName, table := range tables {
		if !db.Migrator().HasTable(table) {
			missingTables = append(missingTables, tableName)
			allTablesExist = false
		}
	}

	if !allTablesExist {
		slog.Warn("The following tables are missing: ", "tables", missingTables)
	}
	return allTablesExist
}

func pingDatabase(db *gorm.DB) bool {
	sqlDB, err := db.DB()
	if err != nil || sqlDB.Ping() != nil {
		slog.Error("Database connection/ping failed: %v", err)
		return false
	}
	return true
}

func createAdminIfNotExists(db *gorm.DB) error {
	suChatID := viper.GetInt64("SUPERUSER_ID")
	var user models.User

	err := db.Where("chat_id = ?", suChatID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = models.User{
			ChatId:    suChatID,
			Username:  viper.GetString("SUPERUSER_NAME"),
			FirstName: viper.GetString("SUPERUSER_FIRSTNAME"),
			LastName:  viper.GetString("SUPERUSER_LASTNAME"),
			IsAdmin:   helpers.BoolToPointer(true),
		}
		if err = db.Create(&user).Error; err != nil {
			return err
		}
		slog.Info("Superuser created")
	}
	return err
}
