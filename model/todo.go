package model

import (
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"" json:"title"`
	IsActive        bool           `gorm:"default:true" json:"is_active"`
	Priority        string         `gorm:"default:'very-high'" json:"priority"`
	CreatedAt       *time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ActivityGroupId uint           `json:"activity_group_id"`
}

type Todos []Todo

func (t *Todo) IsEmpty() bool {
	return reflect.DeepEqual(&Todo{}, t)
}

func (t *Todos) List() error {
	return db.Find(&t).Error
}
func (t *Todos) ListByActivity(id int) error {
	return db.Find(&t, "activity_group_id=?", id).Error
}

func (t *Todo) Get(id int) error {
	return db.First(&t, "id=?", id).Error
}

func (t *Todo) Create() error {
	if t.IsEmpty() {
		return fmt.Errorf("no todo created")
	}
	if len(t.Title) < 1 {
		return fmt.Errorf("Todo title is required")
	}
	if t.ActivityGroupId < 1 {
		return fmt.Errorf("ActivityGroupId is required")
	}

	return db.Create(&t).Error
}

func (t *Todo) Update(id int) error {
	if t.IsEmpty() {
		return fmt.Errorf("no todo updated")
	}
	temp := &Todo{}
	*temp = *t

	if err := t.Get(id); err != nil {
		return fmt.Errorf("Todo with ID %d Not Found", id)
	}
	return db.Model(&t).Updates(&temp).Error
}

func (t *Todo) Delete(id int) error {
	if err := t.Get(id); err != nil {
		return fmt.Errorf("Todo with ID %d Not Found", id)
	}
	return db.Delete(&t, "id=?", id).Error
}
