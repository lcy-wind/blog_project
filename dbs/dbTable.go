package dbs

import (
	loggerutils "blog_project/loggerUtils"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ERR error

func InitDB() {
	dsn := "p2p:p2pA!123@tcp(192.168.66.149:3306)/zjbxinsurance?charset=utf8mb4&parseTime=True&loc=Local"
	DB, ERR = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if ERR != nil {
		loggerutils.Logger.Error("数据库连接失败", zap.Error(ERR))
	} else {
		loggerutils.Logger.Info("数据库连接成功")
	}
	// 自动迁移模式，根据结构体自动创建表
	DB.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func init() {
	InitDB()
}

type User struct {
	Id        uint           `gorm:"primarykey;"`
	Username  string         `gorm:"type:varchar(255);not null;unique"`
	Password  string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除标记"`
}

type Post struct {
	Id        uint           `gorm:"primarykey"`
	UserId    uint           `gorm:"index"`
	Title     string         `gorm:"type:varchar(255);not null"`
	Content   string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除标记"`
}

type Comment struct {
	Id        uint           `gorm:"primarykey"`
	PostId    uint           `gorm:"index"`
	UserId    uint           `gorm:"index"`
	Content   string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除标记"`
}
