package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobkaufmann/hallpass/pkg/datastore"
)

var store = datastore.NewDatastore(nil)

// Handler builds and returns a router for the hallpass API
func Handler() *gin.Engine {
	m := gin.Default()

	// User routes
	m.GET("/api/users/:id", serveUser)
	m.GET("/api/users", serveUsers)
	m.POST("/api/users", serveCreateUser)
	m.DELETE("/api/users/:id", serveDeleteUser)
	m.PUT("/api/users/:id", serveUpdateUser)

	// Pass routes
	m.GET("/api/passes/:id", servePass)
	m.GET("/api/passes", servePasses)
	m.POST("/api/passes", serveCreatePass)
	m.PUT("/api/passes/:id", serveUpdatePass)

	return m
}
