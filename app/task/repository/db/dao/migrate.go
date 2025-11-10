package dao

import "go-micro-todoList/app/task/repository/db/model"

// mapping user struct into DB table
func migration() {
	_db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Task{})
}
