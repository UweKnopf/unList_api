package model

import (
	"unList_api/database"

	"gorm.io/gorm"
)

type ToDoItem struct {
	gorm.Model
	Content    string `gorm:"type:text" json:"content"`
	Tick       bool   `gorm:"type:bool" json:"tick"`
	ToDoListID uint
}

func (item *ToDoItem) CreateItem() (*ToDoItem, error) {
	err := database.Database.Create(&item).Error
	if err != nil {
		return &ToDoItem{}, err
	}
	return item, nil
}

func (item *ToDoItem) ToggleTick() (*ToDoItem, error) {
	if item.Tick {
		err := database.Database.Model(&item).Update("Tick", false).Error
		if err != nil {
			return &ToDoItem{}, err
		}
		return item, nil
	} else {
		err := database.Database.Model(&item).Update("Tick", true).Error
		if err != nil {
			return &ToDoItem{}, err
		}
		return item, nil
	}
}
