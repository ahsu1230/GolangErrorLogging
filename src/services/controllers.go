package services

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
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
	logger.LogInfo("Hero found", logger.LogFields{
		"hero": hero,
	})
	c.JSON(http.StatusOK, &hero)
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
	// err1 := errors.New("error inserting entry")
	err2 := errors.Wrap(err1, "repo failure")
	err3 := errors.Wrap(err2, "ctrl failure")
	err4 := errors.Wrap(fmt.Errorf("%w (%w)", entities.ErrCtrl, err3), "trace")
	// err2 := fmt.Errorf("%w (%w)", ErrSQL, err1)
	// err3 := fmt.Errorf("%w (%w)", ErrRepo, err2)
	// err4 := fmt.Errorf("%w (%w)", ErrCtrl, err3)

	if (errors.Is(err1, entities.ErrSQL)) {
		logger.LogInfo("err1 matched SQL", logger.LogFields{
			"error": err1,
		})
	}

	if (errors.Is(err2, entities.ErrSQL)) {
		logger.LogInfo("err2 matched SQL", logger.LogFields{})
	}

	if (errors.Is(err3, entities.ErrSQL)) {
		logger.LogInfo("err3 matched SQL", logger.LogFields{})
	}

	if (errors.Is(err3, entities.ErrRepo)) {
		logger.LogInfo("err3 matched Repo", logger.LogFields{})
	}

	if (errors.Is(err3, entities.ErrCtrl)) {
		logger.LogInfo("err3 matched Ctrl", logger.LogFields{})
	}

	if (errors.Is(err4, entities.ErrCtrl)) {
		logger.LogInfo("err4 matched Ctrl", logger.LogFields{})
	}

	c.Error(err4)
	return
	// c.AbortWithError(http.StatusInternalServerError, err4)
}