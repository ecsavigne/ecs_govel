CREATE TABLE IF NOT EXISTS estructuras (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    tipo_estructura VARCHAR(70)
) ENGINE = InnoDB;