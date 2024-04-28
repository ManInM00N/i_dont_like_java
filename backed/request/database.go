package request

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db *gorm.DB
)

const (
	dsn = "sqlserver://sa:czp233@localhost:7236?database=gorm"
)

type Account struct {
	Name     string `gorm:"type:varchar(20);not null;primaryKey" json:"name" binding:"required"`
	Password string `gorm:"type:varchar();not null" json:"password" binding:"required"`
	Email    string `gorm:"type:varchar(60);not null" json:"email" binding:"required"`
}

func (a *Account) TableName() string {
	return "account"
}

type Message struct {
	Name             string `gorm:"type:varchar(20);not null;primaryKey" json:"name" binding:"required"`
	SelfIntroduction string `gorm:"type:varchar(400);not null" json:"selfIntroduction" `
}

func (a *Message) TableName() string {
	return "message"
}
func DBInit() {
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB, _ := db.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//  设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//  设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟
	db.AutoMigrate(Account{})
}
