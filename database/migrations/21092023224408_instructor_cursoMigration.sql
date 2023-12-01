CREATE TABLE IF NOT EXISTS instructor_cursos (
    id INT NOT NULL PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    instructor_ci VARCHAR(11),
    curso_id INT,
    FOREIGN KEY (instructor_ci) REFERENCES instructores(ci)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (curso_id) REFERENCES cursos(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE = InnoDB;