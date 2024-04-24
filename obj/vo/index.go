package vo

import (
	"Sgrid/server/SubServer/BlogServer/obj/dao"
	"time"
)

type ColumnDetail struct {
	Column     dao.Column `json:"column"`
	ArticleLen int        `json:"article_len"`
}

type UserVo struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Password   string    `json:"password,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	Level      int       `json:"level,omitempty"`
	Token      string    `json:"token,omitempty"`
}
