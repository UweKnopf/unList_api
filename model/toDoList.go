package model

import (
	"unList_api/database"

	"gorm.io/gorm"
)

type toDoList struct {
	gorm.Model
	Content string     `gorm:"type:text" json:"content"`
	Items   []toDoItem `gorm:"foreignkey:toDoItemID"`
	UserID  uint
}

func (entry *Entry) CreateItemList() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
