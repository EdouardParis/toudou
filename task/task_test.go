package task

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTask(t *testing.T) {
	is := assert.New(t)
	db, _ := gorm.Open("sqlite3", "test.db")
	db.LogMode(true)

	db.CreateTable(new(Task))

	tasks := &Tasks{Db: *db}

	router := gin.Default()
	router.GET("/tasks", tasks.GetAll)
	router.POST("/tasks", tasks.Create)
	router.GET("/task/:id", tasks.Get)
	router.PATCH("/tasks/:id", tasks.Update)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Code, 404)

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/task/1", nil)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Code, 404)

	resp = httptest.NewRecorder()
	reader := strings.NewReader("{ \"name\": \"first task\", \"description\":\"find some coffee\"}")
	req, _ = http.NewRequest("POST", "/tasks", reader)
	router.ServeHTTP(resp, req)

	db.DropTable(&Task{})

	defer db.Close()
}
