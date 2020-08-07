package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ahsu1230/golangwebservertutorial/src/entities"
	"github.com/ahsu1230/golangwebservertutorial/src/logger"
	"github.com/ahsu1230/golangwebservertutorial/src/services"
)

func main() {
	logger.SetupLogger()

    router := gin.Default()
    router.Use(JSONAppErrorReporter())
	router.GET("/ping1", services.GetPing1)
	router.GET("/ping2", services.GetPing2)
	router.GET("/ping3", services.GetPing3)
	router.GET("/ping4", services.GetPing4)

	router.GET("/pong1", services.GetPong1)
	router.GET("/pong2", services.GetPong2)
	router.GET("/pong3", services.GetPong3)
    router.Run(":3000")
}

func JSONAppErrorReporter() gin.HandlerFunc {
    return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
    return func(c *gin.Context) {
		requestUuid := uuid.New()
		logger.Info("Received Request", logger.Fields{
			"requestUuid": requestUuid,
		})

		c.Set("requestUuid", requestUuid)

		c.Next()
		
		// detectedErrors := c.Errors.ByType(errType)
		detectedErrors := c.Errors
		if (len(detectedErrors) > 0) {
			logger.Info("Handling err...", logger.Fields{})
			wrappedErr := detectedErrors[0]
			var message string
			if (errors.Is(wrappedErr, entities.ErrSQL)) {
				message = "Error saving your data"
			} else if (errors.Is(wrappedErr, entities.ErrRepo)) {
				message = "Error retrieving your data from repo"
			} else if (errors.Is(wrappedErr, entities.ErrCtrl)) {
				message = "Error parsing your data"
			} else {
				message = "Unknown error"
			}
			
			appError := entities.AppError {
				http.StatusInternalServerError,
				message,
				wrappedErr.Error(),
			}
			logger.Error(message, wrappedErr, logger.Fields{
				"code": appError.Code,
				"requestUuid": requestUuid,
			})
			c.AbortWithStatusJSON(appError.Code, &appError)
			return
		}

		logger.Info("Succesfully completed request!", logger.Fields{
			"requestUuid": requestUuid,
		})
		return
    }
}