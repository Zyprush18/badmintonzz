package main

import (
	"fmt"
	"log"

	"github.com/Zyprush18/badmintonzz/internal/database"
	"github.com/Zyprush18/badmintonzz/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err:= godotenv.Load();err != nil {
		log.Fatal("Failed Load Env")
	}

	db, err := database.Connect_DB()
	if err != nil {
		log.Fatal("Failed Connect Database")
	}
	
	defer db.Close()

	router := gin.Default()

	v1 := router.Group("/v1")
	users := v1.Group("/users")

	routes.RegisterRouteUser(users, db)
	
	fmt.Println("🚀 server running on Port: 8080")

	router.Run()
}
