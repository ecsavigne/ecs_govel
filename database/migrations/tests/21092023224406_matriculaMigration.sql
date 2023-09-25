CREATE TABLE IF NOT EXISTS matriculas (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    curso_id INT,
    estudiante_ci VARCHAR(11),
    fecha_ingreso DATETIME,
    FOREIGN KEY (estudiante_ci) REFERENCES estudiantes(ci)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (curso_id) REFERENCES cursos(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE = InnoDB;