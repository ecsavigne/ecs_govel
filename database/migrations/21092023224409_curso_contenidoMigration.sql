CREATE TABLE IF NOT EXISTS curso_contenidos (
    curso_contenido_id INT NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    contenido_ci INT,
    curso_id INT,
    FOREIGN KEY (contenido_ci) REFERENCES contenidos(id),
    FOREIGN KEY (curso_id) REFERENCES cursos(id)
) ENGINE = InnoDB;