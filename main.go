package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var GlobalDatabase *sql.DB

func main() {
	ConfigRuntime()
	ConfigDatabase()
	ConfigServer()
}

func ConfigRuntime() {
	numeroCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numeroCpu)
	fmt.Printf("Running with %d CPUs\n", numeroCpu)
}

func ConfigDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred on .env file please check")
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	localDatabase, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is a error while connecting to the database", errSql)
		panic(errSql)
	} else {
		GlobalDatabase = localDatabase
		fmt.Println("Successfully connected to database!")
	}
}

func ConfigGracefulStop(router *gin.Engine) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutddown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func ConfigServer() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/financial", financialEndpoint)
	}
	ConfigGracefulStop(router)

}

type Payments struct {
	PaymentsDate string
	UserId       int
	Total        int
	Debit        int
	Credit       int
	Dif          int
	Accumulated  int
}

func financialEndpoint(context *gin.Context) {

	data, err := GlobalDatabase.Query("select * from financial_paymentsummary")
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, "Falhou em buscar pagamentos")
	} else {
		var paymentsArray []Payments

		for data.Next() {
			var payment Payments
			if errPayment := data.Scan(
				&payment.PaymentsDate,
				&payment.UserId,
				&payment.Total,
				&payment.Debit,
				&payment.Credit,
				&payment.Dif,
				&payment.Accumulated,
			); errPayment != nil {
				context.AbortWithStatusJSON(http.StatusInternalServerError, errPayment)
				break
			}
			paymentsArray = append(paymentsArray, payment)
		}
		if err = data.Err(); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, "Falhou na execução da query")
		}

		context.JSON(http.StatusOK, paymentsArray)
	}
}
