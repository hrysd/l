select Name as Customers from Customers left outer join Orders on Orders.customerId = Customers.id where Orders.customerId is NULL
