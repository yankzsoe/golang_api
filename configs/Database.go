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
		// url := os.Getenv("externalurl")
		dbhost := os.Getenv("dbhost")
		dbname := os.Getenv("dbname")
		dbpassword := os.Getenv("dbpassword")
		dbusername := os.Getenv("dbusername")
		dbport := os.Getenv("dbport")
		dsn := "host=" + dbhost + " user=" + dbusername + " password=" + dbpassword + " dbname=" + dbname + " port=" + dbport

		log.Println(dsn)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: customLogger,
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
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
