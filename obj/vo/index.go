package vo

import (
	"Simp/servers/BlogServer/obj/dao"
)

type ColumnDetail struct {
	Column     dao.Column `json:"column"`
	ArticleLen int        `json:"article_len"`
}
