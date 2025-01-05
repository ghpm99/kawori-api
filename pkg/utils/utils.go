package utils

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ParseInt(value string, context *gin.Context) int {
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	return valueInt
}

func ParseDate(value string, context *gin.Context) time.Time {
	valueTime, err := time.Parse("2006-01-02", value)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}
	return valueTime
}

func PrintQuery(query string, args []interface{}) {
	log.Printf("Executing Query: %s\nArgs: %v\n", query, args)
}
