select
    *
from
    financial_payment fp
where
    1 = 1
    and fp.status = 1
    and fp.type = 1
    and fp.name like '%%'
    and fp."date" between ''
    and ''
    and fp.installments = 1
    and fp.payment_date between ''
    and ''
    and fp.fixed = false
    and fp.active = true
    and fp.user_id = 1;