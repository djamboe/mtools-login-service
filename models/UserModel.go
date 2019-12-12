package models

type UserModel struct {
	Id        int64  `json:"id"`
	DbId      string `json:"_id"`
	Username  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	Level     int32  `json:"level"`
	Parent    int32  `json:"parent"`
	Status    int32  `json:"status"`
	MemberId  int64  `json:"memberId"`
}
