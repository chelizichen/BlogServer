package dao

import "time"

type User struct {
	ID         uint
	Name       string
	Password   string
	CreateTime time.Time
	Level      int
}

// 专栏
type Column struct {
	Id         uint       `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	Desc       string     `json:"desc,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime *time.Time `json:"update_time,omitempty" gorm:"autoCreateTime"`
	UserId     *uint      `json:"user_id,omitempty" gorm:"default:0"`
	Status     *int       `json:"status,omitempty" gorm:"default:0"`
	Weight     *int       `json:"weight,omitempty" gorm:"default:0"`
	Type       *int       `json:"type,omitempty" gorm:"default:0"`
	Articles   []Article  `json:"articles,omitempty" gorm:"foreignKey:column_id"`
}

// 文章
type Article struct {
	ID         uint       `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	Content    string     `json:"content,omitempty"`
	Type       *int       `json:"type,omitempty" gorm:"default:0"`
	ColumnId   *int       `json:"column_id,omitempty" gorm:"column:column_id;default:0"`
}
