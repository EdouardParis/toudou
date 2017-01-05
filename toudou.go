package main

import (
	"fmt"
	"github.com/edouardparis/toudou/task"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"))

	db, err := gorm.Open("postgres", dsn)

	fmt.Println(dsn)

	if err != nil {
		panic(err)
	}

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	db.LogMode(true)

	tasks := &task.Tasks{Db: *db}

	router := gin.Default()
	router.GET("/tasks", tasks.GetAll)
	router.POST("/tasks", tasks.Create)
	router.GET("/task/:id", tasks.Get)
	router.PATCH("/tasks/:id", tasks.Update)
	router.Run() // listen and server on 0.0.0.0:8080
	defer db.Close()
}
