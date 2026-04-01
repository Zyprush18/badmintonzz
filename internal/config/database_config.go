package config

import "os"


type Database_Env struct {
	Host_DB 		string
	PORT_DB			string
	Database_Type string
	Username      string
	Password      string
	Database_Name string
}

func GetEnv() *Database_Env {
	return &Database_Env{
		Host_DB: 		os.Getenv("HOST_DB"),
		PORT_DB:		os.Getenv("PORT_DB"),
		Database_Type: os.Getenv("DATABASE_TYPE"),
		Username:      os.Getenv("USERNAME_DB"),
		Password:      os.Getenv("PASSWORD_DB"),
		Database_Name: os.Getenv("DATABASE_NAME"),
	}
}