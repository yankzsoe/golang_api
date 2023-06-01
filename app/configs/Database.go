package configs

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var onceDb sync.Once

var customLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		LogLevel: logger.Info, // Log level
	},
)

func InitDB() {
	onceDb.Do(func() {
		// Create connection to postgres
		// dbhost := os.Getenv("dblocalhost")
		dbhost := os.Getenv("externalurl")

		db, err := gorm.Open(postgres.Open(dbhost), &gorm.Config{
			Logger: customLogger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

		if err != nil {
			log.Fatal("Error when try connecting to database: ", err.Error())
			panic(err)
		}

		// Active debug mode
		// db.Debug()

		// if err != nil {
		// 	log.Fatal("Error when open connection to database: ", err.Error())
		// 	panic(err)
		// }

		DB = db
	})
}

func GetDB() *gorm.DB {
	return DB
}
