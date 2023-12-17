Inspirado en laravel 
Para probara si el server esta corriendo segun la configuracion de ip y proxy creada en
# config/app.ini
    Aqui se encuentra la configuracion correspondiente al servidor Web y a la App
# db.ini 
    Aqui ó, correspondiente a base datos, driver path del archivo env con variable
var para base de datos ecs_govel, ej:
---------------------Config archivo .env para base de datos driver=mysql
```
 DB_HOST=localhost
 DB_PORT=${MYSQL_FORWARD_PORT}
 DB_USER=${MYSQL_USER}
 DB_NAME=${MYSQL_DATABASE}
 DB_PASSWD=${MYSQL_PASSWORD}
```
----------------------------------------------------
