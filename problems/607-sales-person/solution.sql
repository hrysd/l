select
    SalesPerson.name
from
    SalesPerson
     left outer join (
        select
            sales_id
        from
            Orders
            inner join Company on Company.com_id = Orders.com_id
        where
            Company.name = 'RED'
    ) o on o.sales_id = SalesPerson.sales_id
where
    o.sales_id is null
