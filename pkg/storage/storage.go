package storage

import (
	"epiphanius_bot/pkg/models"
	"epiphanius_bot/pkg/types"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(d *gorm.DB) {
	db = d
}

func GetHolidaysForTelebot() ([]models.Holiday, error) {
	var holidays []models.Holiday
	err := db.Find(&holidays).Error
	return holidays, err
}

func CreateHoliday(request models.Holiday) error {
	return db.Create(&request).Error
}

func CreateUser(users models.Users) error {
	var exists bool
	if err := db.Model(users).
		Select("count(*) > 0").
		Where("user_id = ?", users.UserID).
		Find(&exists).
		Error; err != nil {
		log.Info("user existence check error")
	}

	if exists == false {
		if err := db.Create(&users).Error; err != nil {
			log.Info("error create new user")
		}
	}
	return nil

}

func DeleteHoliday(id int) error {

	var holiday models.Holiday

	return db.Delete(&holiday, id).Error

}

func GetUsers() ([]models.Users, error) {

	var users []models.Users
	err := db.Find(&users).Error
	return users, err
}

func GetHoliday(id int) (models.Holiday, error) {
	var holiday models.Holiday
	err := db.First(&holiday, id).Error
	if err != nil {
		log.Infof("error get holiday: %v", id)
	}

	return holiday, err
}

func GetHolidays() ([]models.Holiday, error) {
	var holidays []models.Holiday
	err := db.Find(&holidays).Error
	return holidays, err
}

func UpdateHoliday(id int, r models.Holiday) error {
	var holiday models.Holiday

	if err := db.First(&holiday, id).Error; err != nil {
		log.Infof("holiday with id:%v - not found", id)
		return err
	}

	holiday.Name = r.Name
	holiday.Description = r.Description
	holiday.Date = r.Date

	return db.Save(&holiday).Error
}

func NextHoliday() ([]types.Message, error) {

	const getMessagesQuery = `SELECT
	name,
	description,
	date
	FROM holidays
	WHERE date > now()::date
	ORDER BY date
	LIMIT 1;`

	var messages []types.Message

	err := db.Raw(getMessagesQuery).Scan(&messages).Error
	return messages, err
}

func GetMessages() ([]types.Message, error) {

	const getMessagesQuery = `SELECT
		name,
		description,
		date,
		CASE WHEN date = now()::date then 'today' else 'in_advance' END as type
	FROM holidays
	WHERE date BETWEEN now()::date AND now()::date + '1d'::interval;`

	var messages []types.Message

	err := db.Raw(getMessagesQuery).Scan(&messages).Error
	return messages, err
}
