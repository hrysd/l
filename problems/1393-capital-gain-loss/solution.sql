select
    stock_name,
    SUM(
        CASE
        WHEN operation = "Buy" THEN -price
        ELSE price
        end
    ) as capital_gain_loss
from
    Stocks
group by
    stock_name
