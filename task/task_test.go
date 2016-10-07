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

	db.DB()
	db.SingularTable(true)
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
	is.Equal(resp.Body.String(), "{\"error\":\"no tasks in the table\"}\n")
	is.Equal(resp.Code, 404)

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/task/1", nil)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Body.String(), "{\"error\":\"no task found with id: 1\"}\n")
	is.Equal(resp.Code, 404)

	resp = httptest.NewRecorder()
	reader := strings.NewReader(`{ "name": "first task", "description": "find some coffee", "progression": 50 }`)
	req, _ = http.NewRequest("POST", "/tasks", reader)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Code, 201)

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/task/1", nil)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Code, 200)

	resp = httptest.NewRecorder()
	reader = strings.NewReader(`{ "name": "first task", "description": "find some coffee", "progression": 101}`)
	req, _ = http.NewRequest("POST", "/tasks", reader)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Code, 400)

	resp = httptest.NewRecorder()
	reader = strings.NewReader(`{ "name": "", "description": "find some coffee", "progression": 100}`)
	req, _ = http.NewRequest("POST", "/tasks", reader)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Body.String(), "{\"error\":\"Task must have a name\"}\n")
	is.Equal(resp.Code, 400)

	resp = httptest.NewRecorder()
	reader = strings.NewReader(`{ "name": "updated task", "description": "find some coffee", "progression": 50 }`)
	req, _ = http.NewRequest("PATCH", "/tasks/1", reader)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Code, 201)

	resp = httptest.NewRecorder()
	reader = strings.NewReader(`{ "name": "updated task", "description": "find some coffee", "progression": -1 }`)
	req, _ = http.NewRequest("PATCH", "/tasks/1", reader)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Body.String(), "{\"error\":\"task progression must be between 0 and 100\"}\n")
	is.Equal(resp.Code, 400)

	resp = httptest.NewRecorder()
	reader = strings.NewReader(`{ "name": "updated task", "description": "find some coffee", "progression": 101 }`)
	req, _ = http.NewRequest("PATCH", "/tasks/1", reader)
	router.ServeHTTP(resp, req)
	is.Equal(resp.Body.String(), "{\"error\":\"task progression must be between 0 and 100\"}\n")
	is.Equal(resp.Code, 400)

	db.DropTable(&Task{})

	defer db.Close()
}
