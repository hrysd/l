select
    customer_id,
    SUM(
        CASE
        WHEN Transactions.transaction_id IS NULL THEN 1
        ELSE 0
        END
    ) as count_no_trans
from
    Visits
    left outer join Transactions on Transactions.visit_id = Visits.visit_id
group by
    Visits.customer_id
having
    count_no_trans > 0
