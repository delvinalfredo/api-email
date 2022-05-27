package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GetAllTemplate(template *[]EmailTemplate, db *gorm.DB) (err error) {
	if err = db.Find(template).Error; err != nil {
		return err
	}
	return nil
}

func CreateTemplate(template *EmailTemplate, db *gorm.DB) (err error) {
	if err = db.Create(template).Error; err != nil {
		return err
	}
	return nil
}

func GetTemplateByID(template *EmailTemplate, id uint, db *gorm.DB) (err error) {
	if err = db.Where("id = ?", id).First(template).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTemplate(template *EmailTemplate, id string, db *gorm.DB) (err error) {
	fmt.Println(template)
	db.Save(template)
	return nil
}

func DeleteTemplate(template *EmailTemplate, id string, db *gorm.DB) (err error) {
	db.Where("id = ?", id).Delete(template)
	return nil
}
