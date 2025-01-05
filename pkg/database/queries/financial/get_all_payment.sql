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
-- WHERE
--     1 = 1
--     AND fp.status = $4
--     AND fp.type = $5
--     AND fp.name LIKE $6
--     AND fp."date" BETWEEN $7
--     AND $8
--     AND fp.installments = $9
--     AND fp.payment_date BETWEEN $10
--     AND $11
--     AND fp.fixed = $12
--     AND fp.active = $13
--     AND fp.user_id = $3
-- LIMIT
--     $1 OFFSET $2;