package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/Zyprush18/badmintonzz/internal/infrastucture/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func Connect_DB() (*sqlx.DB, error) {
	config_db := config.GetEnv()
	db, err := sqlx.Connect(config_db.Database_Type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", config_db.Username, config_db.Password, config_db.Host_DB, config_db.PORT_DB, config_db.Database_Name))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err:= db.Ping();err != nil {
		return nil, err
	}

	// schema
	schema := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255)  NOT NULL,
			password VARCHAR(255) NOT NULL,
			no_hp VARCHAR(20) NOt NULL,
			create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS services (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS schedules (
			id INT AUTO_INCREMENT PRIMARY KEY,
			service_id INT NOT NULL,
			date DATE NOT NULL,
			time TIME NOT NULL,
			duration INT NOT NULL,
			FOREIGN KEY (service_id) REFERENCES services(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP	
		);


		CREATE TABLE IF NOT EXISTS bookings (
			id INT AUTO_INCREMENT PRIMARY KEY,
			type_payment VARCHAR(255) NOT NULL,
			status ENUM('pending', 'confirmed', 'cancelled') NOT NULL,
			description TEXT DEFAULT NULL,
			user_id INT NOT NULL,
			schedule_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (schedule_id) REFERENCES schedules(id)
		);

		CREATE TABLE IF NOT EXISTS payments (
			id INT AUTO_INCREMENT PRIMARY KEY,
			booking_id INT NOT NULL,
			amount DECIMAL(10, 2) NOT NULL,
			payment_method VARCHAR(255) NOT NULL,
			payment_status ENUM('pending', 'completed', 'failed') NOT NULL,
			payment_url TEXT DEFAULT NULL,
			transaction_id VARCHAR(255) DEFAULT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (booking_id) REFERENCES bookings(id)
		);


		CREATE UNIQUE INDEX idx_users_email ON users(email);
		CREATE UNIQUE INDEX idx_services_name ON services(name);
		CREATE UNIQUE INDEX idx_schedules_date_time ON schedules(date, time);
		CREATE UNIQUE INDEX idx_bookings_user_schedule ON bookings(user_id, schedule_id);
		CREATE UNIQUE INDEX idx_payments_booking ON payments(booking_id);

`


	// migrate table to database
	result := db.MustExec(schema)
	if result == nil {
		return nil, errors.New("Failed Migrate Tables")	
	}

	log.Println("Success Migrate Database")

	return db, nil
}
