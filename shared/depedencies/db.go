package depedencies

import (
	"fmt"
	"time"
	"tracking-server/shared/config"
	"tracking-server/shared/dto"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(env *config.EnvConfig, log *logrus.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Errorf("failed to connect to database, with error: %s", err.Error())
	}

	setConnectionConfiguration(db)

	log.Printf("connected to databse with configuration: %s", dsn)

	migrateSchema(db, log)

	seedDatabase(db)

	return db
}

func migrateSchema(db *gorm.DB, log *logrus.Logger) {
	err := db.AutoMigrate(
		&dto.Bus{},
		&dto.News{},
		&dto.Terminal{},
	)

	if err != nil {
		log.Errorf("error migrateing schema, err: %s", err.Error())
	}

	log.Infoln("database migrated")
}

func seedDatabase(db *gorm.DB) {
	terminal := dto.Terminal{}
	db.Create(terminal.Seeder())
}

func setConnectionConfiguration(db *gorm.DB) {
	postgresDb, _ := db.DB()
	postgresDb.SetMaxIdleConns(10)
	postgresDb.SetMaxOpenConns(100)
	postgresDb.SetConnMaxLifetime(time.Hour)
}
