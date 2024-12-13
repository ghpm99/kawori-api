package app

import "github.com/gin-gonic/gin"

type Application struct {
	Router *gin.Engine
}

func InitializeApp() (*Application, error) {

}
