select
    Users.name as name,
    IFNULL(SUM(Rides.distance), 0) as travelled_distance
from Users
    left outer join Rides on Rides.user_id = Users.id
group by
    Users.id
order by
  travelled_distance desc,
  name asc
