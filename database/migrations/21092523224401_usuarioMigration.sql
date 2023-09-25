CREATE TABLE IF NOT EXISTS usuarios (
    id INT NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    nombre varchar(35),
    password varchar(35)
) ENGINE = InnoDB;