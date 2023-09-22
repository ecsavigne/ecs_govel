CREATE TABLE IF NOT EXISTS contenidos (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    tema VARCHAR(255)
) ENGINE = InnoDB;