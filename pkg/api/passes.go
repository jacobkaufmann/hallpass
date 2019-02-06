package api

import (
	"net/http"
	"strconv"

	"github.com/jacobkaufmann/hallpass"

	"github.com/gin-gonic/gin"
)

func servePass(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	pass, err := store.Passes.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = writeJSON(c.Writer, pass); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func servePasses(c *gin.Context) {
	var opt *hallpass.PassListOptions

	if err := c.ShouldBindQuery(&opt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passes, err := store.Passes.List(opt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if passes == nil {
		passes = []*hallpass.Pass{}
	}

	if err = writeJSON(c.Writer, passes); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func serveCreatePass(c *gin.Context) {
	var pass *hallpass.Pass
	if err := c.ShouldBindJSON(&pass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := store.Passes.Create(pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "pass created successfully"})
}

func serveUpdatePass(c *gin.Context) {
	var pass *hallpass.Pass
	if err := c.ShouldBindJSON(&pass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := store.Passes.Update(pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "pass updated successfully"})
}
