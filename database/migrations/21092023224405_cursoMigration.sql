CREATE TABLE IF NOT EXISTS cursos (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    fecha_ingreso DATETIME,
    duracion_hora INT,
    si_certificado BOOLEAN
) ENGINE = InnoDB;