package request

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"main/binary"
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
	dsn = `sqlserver://%s:%s@localhost:%s?database=%s`
)

type Account struct {
	Name     string `gorm:"type:varchar(20);not null;primaryKey" json:"name" binding:"required"`
	Password string `gorm:"size:255;not null" json:"password" binding:"required"`
	Email    string `gorm:"type:varchar(60);not null;unique" json:"email" binding:"required"`
}

func (a *Account) TableName() string {
	return "account"
}

type Message struct {
	Name     string `gorm:"type:varchar(20);not null;primaryKey ;comment:姓名 " json:"name" binding:"required"`
	Motto    string `gorm:"type:varchar(100); comment:座右铭 default:'无' " json:"motto" default:"无"`
	Interest string `gorm:"type:varchar(100);comment:兴趣 default:'无'" json:"interest" default:"无"`
	Group    string `gorm:"type:varchar(100); comment:社团 default:'无'" json:"group" default:"无"`
	Awards   string `gorm:"type:varchar(300); comment:获奖记录 default:'无'" json:"awards" default:"无"`
	Xueli    string `gorm:"type:varchar(50); comment:求学经历 default:'无'" json:"xueli" default:"无"`
	Url      string `gorm:"type:varchar(40);comment:头像地址" json:"url" default:"无"`
}

func (a *Message) TableName() string {
	return "message"
}

type Feature struct {
	Name        string `gorm:"primaryKey;not null; comment: 风景名称" json:"name" binding:"required"`
	Description string `gorm:"not null; comment:介绍" json:"description" binding:"required"`
	URL         string `gorm:"not null; comment: 资源地址" json:"url" binding:"required"`
}

func (a *Feature) TableName() string {
	return "feature"
}

type Comment struct {
	gorm.Model
	Name  string `gorm:"not null; comment: 评论者"json:"name" binding:"required"`
	Inner string `gorm:"not null; comment:评论内容" json:"inner" binding:"required"`
}

func (a *Comment) TableName() string {
	return "comment"
}

type Send struct {
	gorm.Model
	Inner string `gorm:"inner"json:"inner"binding:"inner"`
	Name  string `gorm:"name" json:"name"binding:"name"`
}

func (a *Send) TableName() string { return "send" }
func DBInit() {
	var err error
	var DSN = fmt.Sprintf(dsn, binary.Setting.Sqlusername, binary.Setting.Sqlpassword, binary.Setting.Sqlport, binary.Setting.Sqlbase)
	db, err = gorm.Open(sqlserver.Open(DSN), &gorm.Config{})
	if err != nil {
		binary.DebugLog.Fatalln(err)
		os.Exit(15)
	}
	sqlDB, _ := db.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//  设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//  设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟
	db.AutoMigrate(&Account{}, &Comment{}, &Feature{}, &Message{}, &Send{})
	if !db.Migrator().HasTable(&Account{}) {
		os.Exit(17)
	}
	if !db.Migrator().HasTable(&Message{}) {
		os.Exit(18)
	}
	if !db.Migrator().HasTable(&Comment{}) {
		os.Exit(19)
	}
	if !db.Migrator().HasTable(&Feature{}) {
		os.Exit(20)
	}
}
