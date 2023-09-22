CREATE TABLE IF NOT EXISTS personas (
    ci VARCHAR(11) NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    nombre VARCHAR(35),
    apellidos VARCHAR(100),
    dir VARCHAR(255),
    mail VARCHAR(70)
) ENGINE = InnoDB;