package controllers

import (
	"crud-revel/app"
	"crud-revel/app/models"
	"strconv"

	"github.com/revel/revel"
)

type Users struct {
	*revel.Controller
}

func (u Users) GetAllUsers() revel.Result {
	var users []models.User

	query := "SELECT id, name, age, address FROM users"
	rows, err := app.DB.Query(query)

	if err != nil {
		return u.RenderJSON(sendResponse(400, err.Error(), nil))
	}

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			revel.AppLog.Error(err.Error())
		} else {
			users = append(users, user)
		}
	}

	return u.RenderJSON(sendResponse(200, "Success", users))
}

func (u Users) InsertUser(user models.User) revel.Result {
	err := u.Request.ParseForm()
	if err != nil {
		return u.RenderJSON(sendResponse(400, err.Error(), nil))
	}

	user.Name = u.Request.Form.Get("name")
	user.Age, _ = strconv.Atoi(u.Request.Form.Get("age"))
	user.Address = u.Request.Form.Get("address")
	user.Validate(u.Validation)

	if u.Validation.HasErrors() {
		return u.RenderJSON(sendResponse(200, "Insufficient Input Arguments", nil))
	}

	_, errQuery := app.DB.Exec("INSERT INTO users(name, age, address) VALUES (?, ?, ?, ?)",
		user.Name,
		user.Age,
		user.Address,
	)

	if errQuery != nil {
		return u.RenderJSON(sendResponse(400, errQuery.Error(), nil))
	} else {
		return u.RenderJSON(sendResponse(200, "Success", nil))
	}
}

func (u Users) UpdateUser(id int, user models.User) revel.Result {
	err := u.Request.ParseForm()
	if err != nil {
		return u.RenderJSON(sendResponse(400, err.Error(), nil))
	}

	user.ID = id
	user.Name = u.Request.Form.Get("name")
	user.Age, _ = strconv.Atoi(u.Request.Form.Get("age"))
	user.Address = u.Request.Form.Get("address")
	user.Validate(u.Validation)

	if u.Validation.HasErrors() {
		return u.RenderJSON(sendResponse(200, "Insufficient Input Arguments", nil))
	}

	_, errQuery := app.DB.Exec("UPDATE users SET name=?, age=?, address=? WHERE id=?",
		user.Name,
		user.Age,
		user.Address,
		user.ID,
	)

	if errQuery != nil {
		return u.RenderJSON(sendResponse(400, errQuery.Error(), nil))
	} else {
		return u.RenderJSON(sendResponse(200, "Success", nil))
	}
}

func (u Users) DeleteUser(id int) revel.Result {
	_, errQuery := app.DB.Exec("DELETE FROM users WHERE id=?", id)

	if errQuery != nil {
		return u.RenderJSON(sendResponse(400, errQuery.Error(), nil))
	} else {
		return u.RenderJSON(sendResponse(200, "Success", nil))
	}
}

func sendResponse(status int, message string, data []models.User) models.UserResponse {
	var response models.UserResponse
	response.Status = status
	response.Message = message
	response.Data = data
	return response
}
