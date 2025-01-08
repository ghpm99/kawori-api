package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
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
	args_json, _ := json.Marshal(args)
	log.Printf("Executing Query: %s\nArgs: %v\n", query, string(args_json))
}

func GenerateQueryFilter[T any](filter T) string {

	fmt.Println("GenerateQueryFilter")
	t := reflect.TypeOf(filter)
	v := reflect.ValueOf(filter)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		fmt.Printf("Campo: %s, Valor: %v, tipo: %s\n", field.Name, value, reflect.TypeOf(value))

	}

	return ""

}
