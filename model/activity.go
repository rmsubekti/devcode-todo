package model

import (
	"fmt"
	"reflect"
	"time"
)

type Activity struct {
	ID         uint       `gorm:"primaryKey" json:"id" `
	Title      string     `gorm:"type:varchar(150)" json:"title" `
	Email      string     `gorm:"type:varchar(150)" json:"email" `
	CreatedAt  *time.Time `json:"created_at" time_format:"sql_date"`
	UpdateddAt *time.Time `json:"updated_at" time_format:"sql_date"`
	DeletedAt  *time.Time `json:"deleted_at" time_format:"sql_date"`
	Todos      Todos      `gorm:"foreignkey:ActivityGroupId" json:"-" `
}

type Activities []Activity

func (a *Activity) IsEmpty() bool {
	return reflect.DeepEqual(&Activity{}, a)
}

func (a *Activity) IsExist(id int) bool {
	return db.First(&Activity{}, "id=?", id).RowsAffected > 0
}

func (a *Activities) List() error {
	return db.Find(&a).Error
}

func (a *Activity) Get(id int) error {
	return db.First(&a, "id=?", id).Error
}

func (a *Activity) Create() error {
	if a.IsEmpty() {
		return fmt.Errorf("no activity created")
	}
	return db.Create(&a).Error
}

func (a *Activity) Update(id int) error {

	if !a.IsExist(id) {
		return fmt.Errorf("Activity with ID %d Not Found", id)
	}
	if a.IsEmpty() {
		return fmt.Errorf("no activity updated")
	}
	return db.Model(&Activity{ID: uint(id)}).Updates(&a).Error
}

func (a *Activity) Delete(id int) error {
	if !a.IsExist(id) {
		return fmt.Errorf("Activity with ID %d Not Found", id)
	}
	if err := db.Delete(&Todo{}, "ActivityGroupId=?", id).Error; err != nil {
		return err
	}
	return db.Delete(&a, "id=?", id).Error
}
