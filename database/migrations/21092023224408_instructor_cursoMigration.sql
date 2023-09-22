CREATE TABLE IF NOT EXISTS instructor_cursos (
    instructor_curso_id INT NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    instructor_ci VARCHAR(11),
    curso_id INT,
    FOREIGN KEY (instructor_ci) REFERENCES instructores(ci),
    FOREIGN KEY (curso_id) REFERENCES cursos(id)
) ENGINE = InnoDB;