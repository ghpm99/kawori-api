UPDATE
    financial_payment
SET
    "type" = $1,
    "name" = $2,
    "date" = $3,
    installments = $4,
    payment_date = $5,
    fixed = $6,
    active = $7,
    value = $8,
    status = $9,
    invoice_id = $10
WHERE
    1 = 1
    AND id = $11
    AND user_id = $12