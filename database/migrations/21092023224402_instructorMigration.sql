CREATE TABLE IF NOT EXISTS instructores (
    ci VARCHAR(11) NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (ci) REFERENCES personas(ci)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE = InnoDB;