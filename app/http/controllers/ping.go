package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"

	model "confetti-framework/app/models"
	"confetti-framework/config"
)

func GetTodo(_ inter.Request) inter.Response {
	db := config.GetDBInstance()
	todo := []model.Todo{}
	if err := db.Find(&todo).Error; err != nil {
		return outcome.Json(err).Status(400)
	}

	return outcome.Json(todo)
}
func GetSingleTodo(request inter.Request) inter.Response {
	id := request.Parameter("id").String()
	db := config.GetDBInstance()
	todo := []model.Todo{}
	if err := db.Find(&todo, id).Error; err != nil {
		return outcome.Json(err).Status(400)
	}

	return outcome.Json(todo)
}
