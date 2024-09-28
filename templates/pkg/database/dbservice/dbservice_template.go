package dbservice

import (
	"go-telebot-init/pkg/database/models"
	"gorm.io/gorm"
)

type DBSImpl struct {
	DBS *gorm.DB
}

func NewDBService(db *gorm.DB) DBService {
	return &DBSImpl{DBS: db}
}

type DBService interface {
	Create(model interface{}) error
	Update(model interface{}) error
	GetByID(id uint, model interface{}) error
	GetAll(models interface{}) error
	GetUserByChatID(chatID int64) (*models.User, error)
	Delete(id uint, model interface{}) error
	CloseConnection() error
}

func (d *DBSImpl) Create(model interface{}) error {
	return d.DBS.Create(model).Error
}

func (d *DBSImpl) Update(model interface{}) error {
	return d.DBS.Save(model).Error
}

func (d *DBSImpl) GetByID(id uint, model interface{}) error {
	return d.DBS.First(model, id).Error
}

func (d *DBSImpl) GetAll(models interface{}) error {
	return d.DBS.Find(models).Error
}

func (d *DBSImpl) GetUserByChatID(chatID int64) (*models.User, error) {
	var user *models.User
	err := d.DBS.Where("chat_id = ?", chatID).First(&user).Error
	return user, err
}

func (d *DBSImpl) Delete(id uint, model interface{}) error {
	return d.DBS.Delete(model, id).Error
}

func (d *DBSImpl) CloseConnection() error {
	sqlDB, err := d.DBS.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
