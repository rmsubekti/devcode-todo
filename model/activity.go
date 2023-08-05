package model

import (
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        uint           `gorm:"primaryKey" json:"id" `
	Title     string         `gorm:"type:varchar(150)" json:"title" `
	Email     string         `gorm:"type:varchar(150)" json:"email" `
	CreatedAt *time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Todos     Todos          `gorm:"foreignkey:ActivityGroupId" json:"-" `
}

type Activities []Activity

func (a *Activity) IsEmpty() bool {
	return reflect.DeepEqual(&Activity{}, a)
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

	if len(a.Title) < 1 {
		return fmt.Errorf("Activity title is required")
	}

	return db.Create(&a).Error
}

func (a *Activity) Update(id int) error {
	if a.IsEmpty() {
		return fmt.Errorf("no activity updated")
	}
	temp := &Activity{}
	*temp = *a

	if err := a.Get(id); err != nil {
		return fmt.Errorf("Activity with ID %d Not Found", id)
	}
	return db.Model(&a).Updates(&temp).Error
}

func (a *Activity) Delete(id int) error {
	if err := a.Get(id); err != nil {
		return fmt.Errorf("Activity with ID %d Not Found", id)
	}

	if err := db.Delete(&Todo{}, "activity_group=?", id).Error; err != nil {
		return err
	}
	return db.Delete(&a, "id=?", id).Error
}
