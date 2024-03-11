package configuration

import (
	"Simp/src/http"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

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
	db.Debug().AutoMigrate(&Column{})
	GORM = db
}
