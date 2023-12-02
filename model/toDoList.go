package model

import (
	"unList_api/database"

	"gorm.io/gorm"
)

type ToDoList struct {
	gorm.Model
	Items  []ToDoItem `json:"items"`
	UserID uint
}

func (entry *ToDoList) CreateItemList() (*ToDoList, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &ToDoList{}, err
	}
	return entry, nil
}

func FindToDoListByID(id uint) (ToDoList, error) {
	var toDoList ToDoList
	err := database.Database.Preload("Entries").Where("ID=?", id).Find(&toDoList).Error
	if err != nil {
		return ToDoList{}, err
	}
	return toDoList, nil
}
