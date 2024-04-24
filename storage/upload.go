package storage

import (
	c "Sgrid/server/SubServer/BlogServer/configuration"
	"Sgrid/server/SubServer/BlogServer/obj/dao"
	"Sgrid/server/SubServer/BlogServer/utils"
)

func SaveUploadInfo(info dao.UploadInfo) int {
	c.GORM.Model(&dao.UploadInfo{}).Create(&dao.UploadInfo{
		Chunksize:   info.Chunksize,
		ChunkLength: info.ChunkLength,
		Hash:        info.Hash,
		ServerName:  info.ServerName,
		SpendTime:   info.SpendTime,
	})
	return 1
}

func GetUploadInfoList(pagenation utils.Pagination) (resp []dao.UploadInfo, tt int64, err error) {
	var dataList []dao.UploadInfo
	var total int64
	err = c.GORM.Model(dataList).
		Where("server_name like ?", "%"+pagenation.Keyword+"%").
		Count(&total).
		Limit(pagenation.Size).
		Offset(pagenation.Offset).
		Find(&dataList).Error

	return dataList, total, err
}
