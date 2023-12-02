package model

import (
	"unList_api/database"

	"gorm.io/gorm"
)

type toDoItem struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	Tick    bool   `gorm:"type:bool" json:"tick"`
	UserID  uint
}

func (item *toDoItem) CreateItem() (*toDoItem, error) {
	err := database.Database.Create(&item).Error
	if err != nil {
		return &toDoItem{}, err
	}
	return item, nil
}
