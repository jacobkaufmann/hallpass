package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobkaufmann/hallpass/pkg/datastore"
)

var store = datastore.NewDatastore(nil)

// Handler builds and returns a router for the hallpass API
func Handler() *gin.Engine {
	m := gin.Default()

	return m
}
