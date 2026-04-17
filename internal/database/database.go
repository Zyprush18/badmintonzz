package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func Connect_DB() (*sqlx.DB, error) {
	config_db := config.GetEnv()
	db, err := sqlx.Connect(config_db.Database_Type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", config_db.Username, config_db.Password, config_db.Host_DB, config_db.PORT_DB, config_db.Database_Name))
	if err != nil {
		return nil, err
	}


	if err:= db.Ping();err != nil {
		return nil, err
	}

	schema := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			no_hp VARCHAR(20) UNIQUE NOT NULL,
			role ENUM('user', 'admin') NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT uk_users_email_phone UNIQUE (email, no_hp)
		);

		CREATE TABLE IF NOT EXISTS services (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS payments (
			id INT AUTO_INCREMENT PRIMARY KEY,
			amount DECIMAL(10, 2) NOT NULL,
			payment_method VARCHAR(255) NOT NULL,
			payment_status ENUM('pending', 'completed', 'failed', 'refunded', 'expired') NOT NULL,
			payment_url TEXT NOT NULL,
			transaction_id TEXT DEFAULT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS bussiness_hour (
			id int AUTO_INCREMENT PRIMARY KEY,
			day VARCHAR(50) NOT NULL,
			open_time TIME,
			close_time TIME,
			is_open BOOL NOT NULL DEFAULT FALSE,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);


		CREATE TABLE IF NOT EXISTS bookings (
			id INT AUTO_INCREMENT PRIMARY KEY,
			date DATE NOT NULL,
			start_time TIME NOT NULL,
			end_time TIME NOT NULL,
			type_payment VARCHAR(50) NOT NULL,
			status_booking ENUM('pending', 'confirmed', 'cancelled') NOT NULL,
			description TEXT DEFAULT NULL,
			user_id INT NOT NULL,
			payments_id INT NOT NULL,
			service_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (payments_id) REFERENCES payments(id),
			FOREIGN KEY (service_id) REFERENCES services(id)
		);
`

	result := db.MustExec(schema)
	if result == nil {
		return nil, errors.New("Failed Migrate Tables")	
	}

	

	log.Println("Success Migrate Database")

	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(30)

	return db, nil
}
