select
  name,
  sum(Transactions.amount) as balance
from
  Users
  left outer join Transactions on Transactions.account = Users.account
group by
  Transactions.account
having
  sum(Transactions.amount) > 10000
