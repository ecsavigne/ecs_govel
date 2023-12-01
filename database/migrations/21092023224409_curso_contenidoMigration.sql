CREATE TABLE IF NOT EXISTS curso_contenidos (
    id INT NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    contenido_id INT,
    curso_id INT,
    FOREIGN KEY (contenido_id) REFERENCES contenidos(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (curso_id) REFERENCES cursos(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE = InnoDB;