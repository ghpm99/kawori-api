package payment_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"kawori/api/internal/api/v1/financial/payment"
	"log"
	"time"
)

func CreatePaymentTableFixture(db *sql.DB) {

	createTableSQL := `
		CREATE TABLE financial_payment (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			"type" INTEGER NOT NULL,
			"name" TEXT NOT NULL,
			"date" DATE NOT NULL,
			installments INTEGER NOT NULL,
			payment_date DATE NOT NULL,
			fixed BOOLEAN NOT NULL,
			active BOOLEAN NOT NULL,
			value NUMERIC NOT NULL,
			status INTEGER NOT NULL,
			invoice_id INTEGER NULL,
			user_id INTEGER NOT NULL
		);
	`
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatalf("Falha ao criar tabela: %v", err)
	}
}

func CreatePaymentsDataFixture(repository *payment.Repository, db *sql.DB) {
	for _, value := range createPaymentsData() {
		ctx := context.Background()
		transaction, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatalf("Falha ao popular tabela: %v", err)
		}

		payment, err := repository.CreatePayment(transaction, value)
		if err != nil {
			log.Fatalf("Falha ao popular tabela: %v", err)
		}
		fmt.Println("Payment criado")
		if json, _ := json.Marshal(payment); json != nil {
			fmt.Println(string(json))

		}

		transaction.Commit()
	}
}

func createPaymentsData() []payment.Payment {
	dataRef, _ := time.Parse("2006-01-02", "2024-01-02")

	var paymentArray = []payment.Payment{
		{
			Status:       1,
			Type:         1,
			Name:         "teste",
			Date:         dataRef,
			Installments: 1,
			PaymentDate:  dataRef,
			Fixed:        false,
			Active:       true,
			Value:        100.0,
			InvoiceId:    1,
			UserId:       1,
		}, {
			Status:       1,
			Type:         1,
			Name:         "teste 2",
			Date:         dataRef,
			Installments: 1,
			PaymentDate:  dataRef,
			Fixed:        false,
			Active:       true,
			Value:        150.0,
			InvoiceId:    1,
			UserId:       1,
		},
	}

	return paymentArray
}
