package storage

import (
	c "Simp/servers/BlogServer/configuration"
	"Simp/servers/BlogServer/obj/dao"
	"Simp/servers/BlogServer/obj/dto"
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
