package config

import "os"


func GetEnv() (string, string, string, string)  {
	return os.Getenv("DATABASE_TYPE"), os.Getenv("USERNAME_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("DATABASE_NAME")
}