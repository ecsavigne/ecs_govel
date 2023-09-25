CREATE TABLE IF NOT EXISTS estudiantes (
    ci VARCHAR(11) NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    si_juridico BOOLEAN,
    FOREIGN KEY (ci) REFERENCES personas(ci)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE = InnoDB;