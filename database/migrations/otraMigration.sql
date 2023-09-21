
CREATE TABLE IF NOT EXISTS otra (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    name VARCHAR(255),
    dir_compleja VARCHAR(255)
) ENGINE = InnoDB;



