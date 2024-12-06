package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"stuoj-ai/internal/conf"
	"time"
)

// 初始化数据库
func InitDatabase() error {
	var err error
	config := conf.Conf.Datebase

	gormConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	}

	dsn := config.User + ":" + config.Pwd + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("Connecting to MySQL:", dsn)
	Db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Println("Failed to connect database!")
		return err
	}

	SqlDb, err = Db.DB()
	if err != nil {
		log.Println("Failed to get sql.SqlDb!")
		return err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDb.SetMaxIdleConns(config.MaxIdle)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDb.SetMaxOpenConns(config.MaxConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	SqlDb.SetConnMaxLifetime(time.Hour)

	/*		err = SqlDb.Ping()
			if err != nil {
				log.Println("Error pinging the database!")
				return err
			}
	*/
	log.Println("Database init success!")
	return nil
}
