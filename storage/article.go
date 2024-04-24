package storage

import (
	c "Sgrid/server/SubServer/BlogServer/configuration"
	"Sgrid/server/SubServer/BlogServer/obj/dao"
	"fmt"
)

func GetArticleList(offset int, size int, keyword string) map[string]interface{} {
	var dataList []dao.Article
	var total int64
	c.GORM.Model(dataList).Where("title like ?", "%"+keyword+"%").Count(&total).Limit(size).Offset(offset).Find(&dataList)
	resp := make(map[string]interface{})
	resp["list"] = dataList
	resp["total"] = total
	return resp
}

func SaveArticle(article dao.Article) int64 {
	if article.ID != 0 {
		c.GORM.Debug().
			Model(&article).
			Select("content", "title").
			Where("id = ?", article.ID).
			Updates(&dao.Article{
				Content:  article.Content,
				Title:    article.Title,
				ColumnId: article.ColumnId,
			})
		return int64(article.ID)
	}
	c.GORM.Create(&article)
	return int64(article.ID)
}

func GetArticle(id int) *dao.Article {
	var article *dao.Article
	c.GORM.First(&article, id)
	fmt.Println("article", article)
	return article
}

func DelArticleById(id int) {
	c.GORM.Model(&dao.Article{}).Update("type", -1).Where("id = ?", id)
}
