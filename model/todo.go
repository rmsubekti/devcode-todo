package model

import (
	"fmt"
	"reflect"
	"time"
)

type Todo struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Title           string    `gorm:"" json:"title"`
	IsActive        bool      `gorm:"default:false" json:"is_active"`
	Priority        string    `gorm:"" json:"priority"`
	CreatedAt       time.Time `time_format:"sql_date" json:"created_at"`
	UpdatedAt       time.Time `time_format:"sql_date" json:"updated_at"`
	DeletedAt       time.Time `time_format:"sql_date" json:"deleted_at"`
	ActivityGroupId uint      `json:"activity_group_id"`
}

type Todos []Todo

func (t *Todo) IsEmpty() bool {
	return reflect.DeepEqual(&Todo{}, t)
}

func (a *Todo) IsExist(id int) bool {
	return db.First(&Todo{}, "id=?", id).RowsAffected > 0
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
	return db.Create(&t).Error
}

func (t *Todo) Update(id int) error {
	if t.IsEmpty() {
		return fmt.Errorf("no todo updated")
	}
	temp := &Todo{}
	*temp = *t

	if err := t.Get(id); err != nil {
		return fmt.Errorf("Todo with id %d is not found", id)
	}
	temp.UpdatedAt = time.Now()
	return db.Model(&t).Updates(&temp).Error
}

func (t *Todo) Delete(id int) error {
	if !t.IsExist(id) {
		return fmt.Errorf("Todo with id %d is not found", id)
	}
	return db.Delete(&t, "id=?", id).Error
}
