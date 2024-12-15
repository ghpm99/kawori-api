SELECT
    fp.debit
FROM
    financial_paymentsummary fp
WHERE
    1 = 1
    AND fp.payments_date BETWEEN ?
    AND ?
    AND fp.user_id = ?
LIMIT
    ? OFFSET ?