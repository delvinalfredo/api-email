package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GetAllEmail(account *[]EmailAccount, db *gorm.DB) (err error) {
	if err = db.Find(account).Error; err != nil {
		return err
	}
	return nil
}


func CreateEmail(account *EmailAccount, db *gorm.DB) (err error) {
	if err = db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func GetEmailByID(account *EmailAccount, id uint, db *gorm.DB) (err error) {
	if err = db.Where("id = ?", id).First(account).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEmail(account *EmailAccount, id string, db *gorm.DB) (err error) {
	fmt.Println(account)
	db.Save(account)
	return nil
}

func DeleteEmail(account *EmailAccount, id string, db *gorm.DB) (err error) {
	db.Where("id = ?", id).Delete(account)
	return nil
}
