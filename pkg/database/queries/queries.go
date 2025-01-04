package queries

import (
	_ "embed"
)

//go:embed financial/create_payment.sql
var CreatePayment string

//go:embed financial/update_payment.sql
var UpdatePayment string

//go:embed financial/get_payment_summary.sql
var GetPaymentSummary string

//go:embed financial/get_all_payment.sql
var GetAllPayments string

//go:embed financial/get_payment.sql
var GetPayment string
