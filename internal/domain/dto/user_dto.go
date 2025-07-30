package dto

import "devtracker/internal/domain/model"


type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func FromModelUser(user *model.User) User {
	return User{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
}

