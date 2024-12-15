package queries

import (
	_ "embed"
)

//go:embed financial/get_payment_summary.sql
var GetPaymentSummary string
