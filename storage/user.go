package storage

import (
	c "Simp/servers/BlogServer/configuration"
	"Simp/servers/BlogServer/obj/dao"
	"Simp/servers/BlogServer/obj/dto"
	"Simp/servers/BlogServer/utils"
	"fmt"
)

func ChangeUserLevel(lvl dto.ChangeLevelDto) int64 {
	c.GORM.Debug().
		Model(&dao.User{}).
		Select("level", "id").
		Where("id = ?", lvl.Id).
		Updates(&dao.User{
			Level: lvl.Level,
		})
	return int64(lvl.Id)
}

func GetUserList(pagenation utils.Pagination) (resp []dao.User, tt int64, err error) {
	var userList []dao.User
	var total int64
	err = c.GORM.Debug().
		Model(userList).
		Select("*").
		Where("name like ?", "%"+pagenation.Keyword+"%").
		Count(&total).
		Offset(pagenation.Offset).
		Limit(pagenation.Size).
		Find(&userList).
		Error
	fmt.Println("userList", userList)
	return userList, total, err
}
