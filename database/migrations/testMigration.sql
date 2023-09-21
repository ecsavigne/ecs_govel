
CREATE TABLE IF NOT EXISTS Test (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    name_a VARCHAR(255)
    otra_id int, 
    FOREIGN KEY (otra_id) REFERENCES otra(id)
) ENGINE = InnoDB;
