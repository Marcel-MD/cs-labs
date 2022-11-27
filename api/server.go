package api

import (
	"github.com/Marcel-MD/cs-labs/api/dto"
	"github.com/gin-gonic/gin"
)

func ListenAndServe() error {
	s := NewUserService()
	e := gin.Default()
	g := e.Group("/api")

	g.GET("/users", func(c *gin.Context) {
		users, err := s.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, users)
	})

	g.GET("/users/:email", func(c *gin.Context) {
		email := c.Param("email")

		user, err := s.Get(email)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, user)
	})

	g.POST("/users/register", func(c *gin.Context) {
		var dto dto.CreateUser

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := s.Register(dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, nil)
	})

	g.POST("/users/login", func(c *gin.Context) {
		var dto dto.CreateUser

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		jwt, err := s.Login(dto)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"token": jwt})
	})

	return e.Run(":8080")
}
