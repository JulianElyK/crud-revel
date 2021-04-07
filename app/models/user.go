package models

import "github.com/revel/revel"

type User struct {
	ID      int    `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	Age     int    `form:"age" json:"age"`
	Address string `form:"address" json:"address"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

func (user User) Validate(v *revel.Validation) {
	v.Check(
		user.Name,
		revel.Required{},
		revel.MaxSize{50},
	)

	v.Check(
		user.Address,
		revel.MaxSize{50},
	)
}
