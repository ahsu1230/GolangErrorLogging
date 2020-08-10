package services

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ahsu1230/golangwebservertutorial/src/entities"
	"github.com/ahsu1230/golangwebservertutorial/src/logger"
)

//
// Success pings
//

// Only status, no body
func GetPing1(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Only string "success" as body
func GetPing2(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

// Object as response body
func GetPing3(c *gin.Context) {
	hero := entities.Hero { "Steve", "Cap" }
	logger.Info("Hero found", logger.Fields{
		"hero": hero,
	})
	c.JSON(http.StatusOK, &hero)
}

func GetPing4(c *gin.Context) {
	requestUuid := c.MustGet("requestUuid").(uuid.UUID)
	ctx := context.WithValue(context.Background(), "requestUuid", requestUuid)
	rowId, err := CreateHero1(ctx, 42)
	if err != nil {
		err = errors.Wrap(err, "Repo Failure")
		logger.Error("Error creating hero1", err, logger.Fields{})
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{ "rowId": rowId})
}

//
// Failure pongs
//

// Response Body: "Error -> String", and prints in webserver handler
func GetPong1(c *gin.Context) {
	err := errors.New("Error -> String")
	c.Error(err)
	c.String(http.StatusInternalServerError, err.Error())
}

// No response body, but logs error (no need for handler log)
func GetPong2(c *gin.Context) {
	err := errors.New("Error -> Abort")
	c.AbortWithError(http.StatusInternalServerError, err)
}

// Ideal, prints out response body, and handler log
func GetPong3(c *gin.Context) {
	err1 := entities.ErrSQL
	err2 := errors.Wrap(err1, "wrap1 failure")
	err3 := errors.Wrap(err2, "wrap2 failure")

	c.Error(err3)
	c.Abort()
	// c.AbortWithError(http.StatusOK, err3)
	// let handler decide response status code, message, etc.
}

// Uses a repo function to pass things around
func GetPong4(c *gin.Context) {
	requestUuid := c.MustGet("requestUuid").(uuid.UUID)
	ctx := context.WithValue(context.Background(), "requestUuid", requestUuid)
	rowId, err := CreateHero2(ctx, 42)
	if err != nil {
		err = errors.Wrap(err, "Repo Failure")
		logger.Error("Error creating hero1", err, logger.Fields{})
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{ "rowId": rowId})
}