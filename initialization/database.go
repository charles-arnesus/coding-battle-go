package initialization

import (
	"fmt"
	"time"

	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func dbInitialization() *gorm.DB {
	username := "myuser"
	password := "mypassword"
	dbName := "mydb"
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta", username, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError:         false,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}

	err = migrateDB(db)
	if err != nil {
		fmt.Printf("Failed to migrate table: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour * 6)

	return db
}

func migrateDB(db *gorm.DB) (err error) {
	// jangan lupa hapus droptable
	// untuk testing
	db.Migrator().DropTable(&flight_model.Aircraft{})
	db.Migrator().DropTable(&flight_model.Destination{})
	db.Migrator().DropTable(&flight_model.FlightRoute{})
	db.Migrator().DropTable(&flight_model.FlightRouteSeat{})
	db.Migrator().DropTable(&user_model.User{})
	db.Migrator().DropTable(&booking_model.BookingSystem{})

	err = db.AutoMigrate(
		&flight_model.Aircraft{},
		&flight_model.Destination{},
		&flight_model.FlightRoute{},
		&flight_model.FlightRouteSeat{},
		&user_model.User{},
		&booking_model.BookingSystem{},
	)

	db.Create(&user_model.User{Username: "admin", Name: "admin", Role: "admin"})
	db.Create(&booking_model.BookingSystem{IsActive: false})

	return
}
