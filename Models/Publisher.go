package Models

import (
	"fmt"

	"gorm.io/gorm"
	"mail.blast/query_model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllPublisher(publisher *[]query_model.PublisherQuery, db *gorm.DB) (err error) {
	if err = db.Raw(`select e.email , t.name ,p.*
	from publishers p
	left join email_accounts e on p.email_account_id = e.id 
	left join email_templates t on p.email_template_id = t.id  
	`).Scan(publisher).Error; err != nil {
		return err
	}
	return nil
}

func CreatePublisher(publisher *Publisher, db *gorm.DB) (err error) {
	if err = db.Create(publisher).Error; err != nil {
		return err
	}
	return nil
}

func GetPublisherByID(publisher *Publisher,  db *gorm.DB, id uint,) (err error) {
	if err = db.Where("id = ?", id).First(publisher).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePublisher(publisher *Publisher, id string, db *gorm.DB) (err error) {
	fmt.Println(publisher)
	db.Save(publisher)
	return nil
}

func DeletePublisher(publisher *Publisher, id string, db *gorm.DB) (err error) {
	db.Where("id = ?", id).Delete(publisher)
	return nil
}
