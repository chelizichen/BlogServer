package configuration

import (
	"Simp/src/http"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	ID         uint
	Name       string
	Password   string
	CreateTime time.Time
}

type Article struct {
	ID         uint       `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	Content    string     `json:"content,omitempty"`
	Type       *int       `json:"type,omitempty" gorm:"default:0"`
}

var GORM *gorm.DB

func InitStorage(ctx http.SimpHttpServerCtx) {
	db, err := gorm.Open(mysql.Open(ctx.StoragePath), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "blog_",
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("Error To init gorm", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error To init db.DB", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.Debug().AutoMigrate(&User{})
	db.Debug().AutoMigrate(&Article{})
	GORM = db
}
