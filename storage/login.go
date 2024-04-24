package storage

import (
	c "Sgrid/server/SubServer/BlogServer/configuration"
	"Sgrid/server/SubServer/BlogServer/obj/dao"
	"Sgrid/server/SubServer/BlogServer/obj/dto"
	"Sgrid/server/SubServer/BlogServer/utils"

	"fmt"
)

func LoginUser(user dto.UserDto) (resp *dao.User, err error) {
	r := c.GORM.Where(&dao.User{
		Name:     user.Name,
		Password: user.Password,
	}).First(&resp).Error
	fmt.Println("e", r)
	return resp, r
}

func LoginByCache(user dto.UserDto) (resp *dao.User, err error) {
	r := c.GORM.Where(&dao.User{
		Name: user.Name,
	}).First(&resp).Error
	fmt.Println("e", r)
	fmt.Println("resp", resp)
	if utils.AuthenticateUser(user.Name, user.Password, *resp) {
		return resp, r
	}
	return nil, r
}

func SaveUser(user dto.UserDto) int {
	var usr *dao.User
	c.GORM.Where("name = ?", user.Name).Find(&usr)
	if usr.Name == "" && usr.ID == 0 {
		c.GORM.Model(&dao.User{}).Create(&dao.User{
			Name:     user.Name,
			Password: user.Password,
			Level:    1,
		})
		return 1
	}
	return 0
}
