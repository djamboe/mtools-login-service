package models

type UserModel struct {
	Id        int    `json:"id"`
	_Id       string `json:"_id"`
	Username  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	Level     int    `json:"level"`
	Parent    int    `json:"parent"`
	Status    int    `json:"status"`
}
