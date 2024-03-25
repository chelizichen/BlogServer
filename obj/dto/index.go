package dto

type UserDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ChangeLevelDto struct {
	Id    uint `json:"id"`
	Level int  `json:"level"`
}
