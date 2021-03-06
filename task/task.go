package task

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Progression int       `json:"progression"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks struct {
	Db gorm.DB
}

func (t *Tasks) GetAll(c *gin.Context) {
	tasks := []Task{}
	res := t.Db.Order("created_at").Find(&tasks)

	if len(tasks) == 0 {
		c.JSON(404, gin.H{"error": "no tasks in the table"})
	} else {
		c.JSON(200, gin.H{"data": res.Value})
	}
}

func (t *Tasks) Get(c *gin.Context) {
	task := &Task{}
	id := c.Param("id")
	res := t.Db.First(&task, id)

	if res.RecordNotFound() {
		c.JSON(404, gin.H{"error": ("no task found with id: " + id)})
	} else {
		c.JSON(200, gin.H{"data": res.Value})
	}
}

func (t *Tasks) Create(c *gin.Context) {
	task := &Task{}

	if c.BindJSON(task) == nil {
		if 0 > task.Progression || 100 < task.Progression {
			c.JSON(400, gin.H{"error": "task progression must be between 0 and 100"})
		} else {
			task.CreatedAt = time.Now()

			res := t.Db.Create(&task)

			if res.Error != nil {
				c.JSON(500, gin.H{"error": "Unable to create the task"})
			} else {
				c.JSON(201, gin.H{"data": res.Value})
			}
		}
	} else {
		c.JSON(400, gin.H{"error": "Task must have a name"})
	}
}

func (t *Tasks) Update(c *gin.Context) {
	updatedTask := &Task{}
	id := c.Param("id")

	if c.BindJSON(updatedTask) == nil {
		if 0 > updatedTask.Progression || 100 < updatedTask.Progression {
			c.JSON(400, gin.H{"error": "task progression must be between 0 and 100"})
		} else {
			res := t.Db.First(&Task{}, id).Updates(Task{
				Name:        updatedTask.Name,
				Description: updatedTask.Description,
				Progression: updatedTask.Progression,
				UpdatedAt:   time.Now(),
			})

			if res.Error != nil {
				c.JSON(500, gin.H{"error": "Unable to update the task"})
			} else {
				c.JSON(201, gin.H{"data": res.Value})
			}
		}
	}
}
