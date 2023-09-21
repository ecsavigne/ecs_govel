# Ejemplo de contenido arquivo de Migrations
 Los archivos dentro de la carpeta migrations seran .sql
 que van hacer ejcutados por medio de *gorm.DB.Exec()
 los mismos deberan ser ejecutados si en la base de datos no existe la table 
 a la que refiere la migration o sip se modifica algun campo.Los nombres de la migration
 se forman como el nombre del archivo .sql que representa la migration sin Migration.sql 
 con la primera letra en mayuscula

# Ejemplo de contenido del archivo .sql
 
Crear la migration o tabla nombre_otronombre(file= nombre_otronombreMigration.sql)
--------------------------------------------------
	CREATE TABLE IF NOT EXISTS <nombre_otronombre> (
		id INT PRIMARY KEY,
		nombre VARCHAR(255),
		edad INT
	);

Agregar un campo a la tabla si es necesario nombre_otronombre(file= nombre_otronombreMigration.sql)
--------------------------------------------------
	ALTER TABLE nombre_otronombre
		ADD direccion VARsCHAR(255);

Modificar un campo de la tabla si es necesario nombre_otronombre(file= nombre_otronombreMigration.sql)
--------------------------------------------------
ALTER TABLE nombre_otronombre
    MODIFY edad INT NOT NULL;

# Ejemplo de Migraciones con Primary y Foreign key
CREATE TABLE IF NOT EXISTS Test (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    name_a VARCHAR(255)
    otra_id int, 
    FOREIGN KEY (otra_id) REFERENCES otra(id)[otratabla(nombre_primarykey)]
) ENGINE = InnoDB;