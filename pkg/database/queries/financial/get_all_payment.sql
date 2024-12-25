SELECT
    fp.id,
    fp.status,
    fp."type",
    fp."name",
    fp."date",
    fp.installments,
    fp.payment_date,
    fp.fixed,
    fp.value,
    fp.invoice_id
FROM
    financial_payment fp
WHERE
    1 = 1
    AND fp.status = $1
    AND fp.type = $2
    AND fp.name LIKE $3
    AND fp."date" BETWEEN $4
    AND $5
    AND fp.installments = $6
    AND fp.payment_date BETWEEN $7
    AND $8
    AND fp.fixed = $9
    AND fp.active = $10
    AND fp.user_id = $11
LIMIT
    $12 OFFSET $13;