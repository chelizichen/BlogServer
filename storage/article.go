package storage

import (
	c "Simp/servers/BlogServer/configuration"
	"fmt"
)

func GetArticleList(offset int, size int, keyword string) map[string]interface{} {
	var dataList []c.Article
	var total int64
	c.GORM.Model(dataList).Where("title like ?", "%"+keyword+"%").Count(&total).Limit(size).Offset(offset).Find(&dataList)
	resp := make(map[string]interface{})
	resp["list"] = dataList
	resp["total"] = total
	return resp
}

func SaveArticle(article c.Article) int64 {
	if article.ID != 0 {
		r := c.GORM.Debug().
			Model(&article).
			Select("content", "title").
			Where("id = ?", article.ID).
			Updates(&c.Article{
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

func GetArticle(id int) *c.Article {
	var article *c.Article
	c.GORM.First(&article, id)
	fmt.Println("article", article)
	return article
}

func DelArticleById(id int) {
	c.GORM.Model(&c.Article{}).Update("type", -1).Where("id = ?", id)
}
