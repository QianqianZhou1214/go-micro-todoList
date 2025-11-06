package dao

import (
	"fmt"
	"go-micro-todoList/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func InitDB() {
	host := config.DbHost
	port := config.DbPort
	user := config.DbUser
	password := config.DbPassword
	database := config.DbName
	charset := config.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)
	fmt.Println(dsn)
	err := Database(dsn)
	if err != nil {
		fmt.Println("database init fail", err)
	}
}

func Database(connString string) error {
	var ormLogger logger.Interface
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString, // dsn data source name
		DefaultStringSize:         256,        // default length of string type
		DisableDatetimePrecision:  true,       // disable datatime precision mysql 5.6
		DontSupportRenameIndex:    true,       // rename index, must delet index
		DontSupportRenameColumn:   true,       // no rename of columns
		SkipInitializeWithVersion: false,      // auto config according to versions
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // user --> users (x)
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	_db = db // change to local db
	return nil
}
