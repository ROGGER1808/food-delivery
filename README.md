
### Setup docker
```shell
docker run --name mysql --privileged=true -e \
MYSQL_ROOT_PASSWORD="Z2Vuc29uMTgwOAo=" -e \
MYSQL_USER="food_delivery" -e \
MYSQL_PASSWORD="Z2Vuc29uMjAwMAo=" -e \
MYSQL_DATABASE="food_delivery" -p 3306:3306 bitnami/mysql:5.7

```

[database sharding ](https://medium.com/pinterest-engineering/sharding-pinterest-how-we-scaled-our-mysql-fleet-3f341e96ca6f)