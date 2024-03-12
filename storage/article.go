package storage

import (
	c "Simp/servers/BlogServer/configuration"
	"Simp/servers/BlogServer/obj/dao"
	"fmt"
)

func GetArticleList(offset int, size int, keyword string) map[string]interface{} {
	var dataList []dao.Article
	var total int64
	c.GORM.Model(dataList).Where("title like ?", "%"+keyword+"%").Count(&total).Limit(20).Offset(offset).Find(&dataList)
	resp := make(map[string]interface{})
	resp["list"] = dataList
	resp["total"] = total
	return resp
}

func SaveArticle(article dao.Article) int64 {
	if article.ID != 0 {
		r := c.GORM.Debug().
			Model(&article).
			Select("content", "title").
			Where("id = ?", article.ID).
			Updates(&dao.Article{
				Content:  article.Content,
				Title:    article.Title,
				ColumnId: article.ColumnId,
			})
		return r.RowsAffected
	}
	r := c.GORM.Create(&article)
	i := r.RowsAffected
	return i
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
