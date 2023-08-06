package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"" json:"title"`
	IsActive        *bool          `gorm:"default:true" json:"is_active"`
	Priority        string         `gorm:"default:'very-high'" json:"priority"`
	CreatedAt       *time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ActivityGroupId uint           `json:"activity_group_id"`
}

type Todos []Todo

func (t *Todos) List() error {
	return db.Find(&t).Error
}
func (t *Todos) ListByActivity(id int) error {
	return db.Find(&t, "activity_group_id=?", id).Error
}

func (t *Todo) Get(id int) error {
	if err := db.First(&t, "id=?", id).Error; err != nil {
		return fmt.Errorf("Todo with ID %d Not Found", id)
	}
	return nil
}

func (t *Todo) Create() error {
	if len(t.Title) < 1 {
		return fmt.Errorf("title cannot be null")
	}
	if t.ActivityGroupId < 1 {
		return fmt.Errorf("activity_group_id cannot be null")
	}

	return db.Create(&t).Error
}

func (t *Todo) Update(id int) error {
	temp := &Todo{}
	*temp = *t

	if err := t.Get(id); err != nil {
		return err
	}
	return db.Model(&t).Updates(&temp).Error
}

func (t *Todo) Delete(id int) error {
	if err := t.Get(id); err != nil {
		return err
	}
	return db.Delete(&t, "id=?", id).Error
}
