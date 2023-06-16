package model

import (
	"personal/TODO-list/database"
	"personal/TODO-list/helper"
	"time"
)

type ToDo struct {
	Id        int
	Item      string
	Marked    bool
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}

func CreateTodoItem(item, createdBy string) error {
	toDo := ToDo{
		Item:      item,
		Marked:    false,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}

	err := database.Db.Create(&toDo).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodoItem(id int) error {
	err := database.Db.Where("id = ?", id).Delete(ToDo{}).Error
	if err != nil {
		return err
	}

	return nil
}

type ToDoList struct {
	Id        int    `json:"id"`
	Item      string `json:"item"`
	Marked    int    `json:"marked"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

func ListTodoItem(pagination *helper.Pagination) (toDoList []ToDoList, err error) {
	// query builder
	queryBuilder := database.Db.Table("to_do").Select("*, DATE_FORMAT(created_at,'%Y-%m-%d %T') AS created_at")

	// Count
	var count int64
	err = queryBuilder.Count(&count).Error
	if err != nil {
		return nil, err
	}
	pagination.TotalRecord = count

	// Result
	offset := (pagination.Page - 1) * pagination.Record
	err = queryBuilder.
		Limit(pagination.Record).
		Offset(offset).
		Order(pagination.Sort + " " + pagination.Order).
		Find(&toDoList).Error
	if err != nil {
		return nil, err
	}

	return toDoList, nil
}

func UpdateMarkTodoItem(id int, updatedBy string) error {
	var toDo ToDo

	err := database.Db.Where("id = ?", id).First(&toDo).Error
	if err != nil {
		return err
	}

	var marked bool
	if !toDo.Marked {
		marked = true
	} else {
		marked = false
	}

	updateTodo := map[string]interface{}{
		"marked":     marked,
		"updated_by": updatedBy,
		"updated_at": time.Now(),
	}

	err = database.Db.Model(&toDo).Updates(updateTodo).Error
	if err != nil {
		return err
	}

	return nil
}
