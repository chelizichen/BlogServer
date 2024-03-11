package storage

import (
	c "Simp/servers/BlogServer/configuration"
	"Simp/servers/BlogServer/utils"
	"fmt"
)

type Result struct {
	Id    int
	Total int
}

func GetColumnList(pagination utils.Pagination) map[string]interface{} {
	// var columnList []struct {
	// 	Column        c.Column
	// 	NumOfArticles int
	// }
	var columnResp []*c.ColumnDetail
	var columnList []c.Column
	var total int64
	c.GORM.Model(columnList).
		Where("title like ?", "%"+pagination.Keyword+"%").
		Count(&total).Limit(pagination.Size).
		Offset(pagination.Offset).
		Find(&columnList)

	ids := make([]uint, 0)
	for _, v := range columnList {
		Val := v
		columnResp = append(columnResp, &c.ColumnDetail{
			Column:     Val,
			ArticleLen: 0,
		})
		ids = append(ids, (Val.Id))
	}

	var resu []Result
	c.GORM.Raw(`SELECT
	c.id,
	COUNT( a.column_id ) AS total 
FROM
	blog_column c
	LEFT JOIN blog_article a ON a.column_id = c.id 
	WHERE c.id in (?)
GROUP BY
	c.id`, ids).Scan(&resu)
	for _, v := range columnResp {
		for _, r := range resu {
			if uint(r.Id) == (v.Column.Id) {
				R := r
				v.ArticleLen = R.Total
			}
		}
	}
	fmt.Println("columnResp", columnResp)
	resp := make(map[string]interface{})
	resp["list"] = columnResp
	resp["total"] = total
	return resp
}

func GetColumnDetail(columnID int) (column *c.Column) {
	var articles []c.Article

	// 通过 Column ID 查询对应的 Article
	c.GORM.Preload("Articles").Where("id = ?", columnID).First(&column)

	// 遍历 Column 下的所有 Article，并提取其 ID
	for _, article := range column.Articles {
		articles = append(articles, c.Article{ID: article.ID, Title: article.Title})
	}
	column.Articles = articles
	return column
}

func SaveColumn(column c.Column) int64 {
	r := c.GORM.Create(&column)
	return (r.RowsAffected)
}

func SaveArticleInColumn(id, cid int) int64 {
	v := &c.Article{
		ColumnId: &cid,
	}
	r := c.GORM.Debug().
		Model(&v).
		Select("column_id").
		Where("id = ?", id).
		Updates(&c.Article{
			ColumnId: v.ColumnId,
		})
	return r.RowsAffected
}
