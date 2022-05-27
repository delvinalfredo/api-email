package Models

import "gorm.io/gorm"

func GetAllUser(user *[]User, db *gorm.DB) (err error) {
	if err = db.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(username string, db *gorm.DB) (err error) {
	if err = db.Where("username = ?", username).Error; err != nil {
		return err
	}
	return nil
}

func Register(user *User, db *gorm.DB) (err error) {
	if err = db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string, db *gorm.DB) (err error) {
	if err = db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id string, db *gorm.DB) (err error) {
	db.Where("id = ?", id).Delete(user)
	return nil
}
