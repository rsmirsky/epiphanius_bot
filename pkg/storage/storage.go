package storage

import (
	"epiphanius_bot/pkg/models"
	"epiphanius_bot/pkg/types"
	"fmt"

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
	//return  db.Create(&users).Error

	//  fmt.Println("==1=1=1=")
	//  fmt.Println(exists)
	//  fmt.Println("==1=1=1=")

	//  return err
	//  if exists == false {
	// 	return  db.Create(&users).Error
	//  }
	//   return nil
	// r := db.Find(&users)
	// exists := r.RowsAffected > 0
	// if exists == true {
	// 	return nil
	// }
	// return db.Create(&users).Error

}

// 	result := db.First(&users)

// 	if result.RowsAffected != nil {
// 		return nil
// 	} else {
// 	return db.Create(&users).Error
// 	}
// }

func DeleteHoliday(id int) error {

	// TODO: не обязательно доставать запись, можно сразу удалить по айди

	var holiday models.Holiday
	if err := db.First(&holiday, id).Error; err != nil {
		return err
	}

	return db.Delete(&holiday).Error
}

func GetUsers() ([]models.Users, error) {
	//var messages []types.Message
	var users []models.Users
	err := db.Find(&users).Error
	return users, err
}

func GetHoliday(id int) (models.Holiday, error) {
	var holiday models.Holiday
	fmt.Println("=id=id=id=")
	fmt.Println(id)
	fmt.Println("=id=id=id=")
	err := db.First(&holiday, id).Error
	fmt.Println("=Getholiday=Getholiday=Getholiday=")
	fmt.Println(holiday)
	fmt.Println("=Getholiday=Getholiday=Getholiday=")
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
	fmt.Println("---------------------")
	fmt.Println(messages)
	fmt.Println("---------------------")
	return messages, err 
}

func GetMessages() ([]types.Message, error) {

	// var messages []types.Message
	// err := db.Find(&messages).Error
	// return messages, err

	const getMessagesQuery = `SELECT
		name,
		description,
		date,
		CASE WHEN date = now()::date then 'today' else 'in_advance' END as type
	FROM holidays
	WHERE date BETWEEN now()::date AND now()::date + '1d'::interval;`

	var messages []types.Message

	err := db.Raw(getMessagesQuery).Scan(&messages).Error
	fmt.Println("---------------------")
	fmt.Println(messages)
	fmt.Println("---------------------")
	return messages, err
}
