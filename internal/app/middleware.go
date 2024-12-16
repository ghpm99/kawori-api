package app

import (
	"encoding/json"
	"kawori/api/internal/config"
	"kawori/api/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authUrl = config.Get("AUTH_ENDPOINT", "http://localhost:8500/auth/user")

func authMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Empty authorization."})
			context.Abort()
			return
		}

		client := http.Client{}
		request, err := http.NewRequest("GET", authUrl, nil)
		if err != nil {
			log.Println(err)
		}
		request.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {token},
		}
		response, err := client.Do(request)
		if err != nil {
			log.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			context.AbortWithStatus(http.StatusUnauthorized)
			context.Abort()
			return
		}
		var user utils.User

		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&user)
		if err != nil {
			log.Panicln(err)
		}

		context.Set("user", user)

		context.Next()

	}
}
