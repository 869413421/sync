package user

import "sync/pkg/model"

type User struct {
	model.BaseModel
	Name     string
	Email    string
	Password string
}
