package storage

import (
	c "Sgrid/server/SubServer/BlogServer/configuration"
	"Sgrid/server/SubServer/BlogServer/obj/dao"
	"Sgrid/server/SubServer/BlogServer/obj/vo"
	"Sgrid/server/SubServer/BlogServer/utils"
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
	var columnResp []*vo.ColumnDetail
	var columnList []dao.Column
	var total int64
	c.GORM.Model(columnList).
		Where("title like ? and type != ? ", "%"+pagination.Keyword+"%", -1).
		Count(&total).Limit(pagination.Size).
		Offset(pagination.Offset).
		Find(&columnList)

	ids := make([]uint, 0)
	for _, v := range columnList {
		Val := v
		columnResp = append(columnResp, &vo.ColumnDetail{
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

func GetColumnDetail(columnID int) (column *dao.Column) {
	var articles []dao.Article

	// 通过 Column ID 查询对应的 Article
	c.GORM.Preload("Articles").Where("id = ?", columnID).First(&column)

	// 遍历 Column 下的所有 Article，并提取其 ID
	for _, article := range column.Articles {
		articles = append(articles, dao.Article{ID: article.ID, Title: article.Title})
	}
	column.Articles = articles
	return column
}

func SaveColumn(column dao.Column) int64 {
	r := c.GORM.Create(&column)
	return (r.RowsAffected)
}

func SaveArticleInColumn(id, cid int) int64 {
	v := &dao.Article{
		ColumnId: &cid,
	}
	r := c.GORM.Debug().
		Model(&v).
		Select("column_id").
		Where("id = ?", id).
		Updates(&dao.Article{
			ColumnId: v.ColumnId,
		})
	return r.RowsAffected
}

func DeleteColumnById(id int) int64 {
	t := -1
	d := c.GORM.Model(&dao.Column{}).Select("type").Where("id = ?", id).Updates(&dao.Column{
		Type: &t,
	})
	return d.RowsAffected
}
