package request

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
	"regexp"
	"time"
)

var (
	db               *gorm.DB
	Rule_name, _     = regexp.Compile("[a-zA-Z0-9]{4,20}")
	Rule_password, _ = regexp.Compile("[a-zA-Z0-9]{6,30}")
	Rule_email, _    = regexp.Compile("[a-zA-Z0-9][\\w\\\\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\\w\\\\.-]*[a-zA-Z0-9]\\.[a-zA-Z][a-zA-Z\\\\.]*[a-zA-Z]")
)

const (
	dsn = "sqlserver://sa:czp233@localhost:1433?database=gorm"
)

type Account struct {
	gorm.Model
	Name     string `gorm:"varchar(20);not null;primarykey" json:"name" binding:"required"`
	Password string `gorm:"size:255;not null" json:"password" binding:"required"`
	Email    string `gorm:"varchar(60);not null" json:"email" binding:"required"`
}

func (a *Account) TableName() string {
	return "account"
}

type Message struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;primarykey comment:姓名 " json:"name" binding:"required"`
	Motto    string `gorm:"type:varchar(100);not null comment:座右铭 default:'无' " json:"motto" default:"无"`
	Interest string `gorm:"type:varchar(100);not null comment:兴趣 default:'无'" json:"interest" default:"无"`
	Group    string `gorm:"type:varchar(100);not null comment:社团 default:'无'" json:"group" default:"无"`
	Awards   string `gorm:"type:varchar(300);not null comment:获奖记录 default:'无'" json:"awards" default:"无"`
	Xueli    string `gorm:"type:varchar(50);not null comment:求学经历 default:'无'" json:"xueli" default:"无"`
}

func (a *Message) TableName() string {
	return "message"
}
func DBInit() {
	var err error
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		os.Exit(15)
	}
	sqlDB, _ := db.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//  设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//  设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Message{})
	if !db.Migrator().HasTable(&Account{}) {
		os.Exit(17)
	}
	if !db.Migrator().HasTable(&Message{}) {
		os.Exit(18)
	}

}
