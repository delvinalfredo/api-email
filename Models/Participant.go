package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GetAllParticipant(participant *[]PublisherParticipant, db *gorm.DB) (err error) {
	if err = db.Find(participant).Error; err != nil {
		return err
	}
	return nil
}

func CreateParticipant(participant *PublisherParticipant, db *gorm.DB) (err error) {
	if err = db.Create(participant).Error; err != nil {
		return err
	}
	return nil
}

// func CreateParticipantExcel(participant *PublisherParticipant,id uint, db *gorm.DB) (err error) {
// 	if err = db.Create(participant).Where("publisher_id = ?", id).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetParticipantByID(participant *[]PublisherParticipant, id uint, db *gorm.DB) (err error) {
	if err = db.Where("publisher_id = ?", id).Find(participant).Error; err != nil {
		return err
	}
	return nil
}

func GetParticipantID(participant *PublisherParticipant, id uint, db *gorm.DB) (err error) {
	if err = db.Where("id = ?", id).First(participant).Error; err != nil {
		return err
	}
	return nil
}

func GetParticipantIDList(participant *[]PublisherParticipant, id uint, db *gorm.DB) (err error) {
	if err = db.Where("publisher_id = ?", id).Find(participant).Error; err != nil {
		return err
	}
	return nil
}

func UpdateParticipant(participant *PublisherParticipant, id string, db *gorm.DB) (err error) {
	fmt.Println(participant)
	db.Save(participant)
	return nil
}

func DeleteParticipant(participant *PublisherParticipant, id string, db *gorm.DB) (err error) {
	db.Where("id = ?", id).Delete(participant)
	return nil
}
