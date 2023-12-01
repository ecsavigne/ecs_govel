CREATE TABLE IF NOT EXISTS estructura_contenidos (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    contenido_id INT,
    estructura_id INT,
    cantidad_clase_plane INT,
    cantidad_tematica_programa INT,
    FOREIGN KEY (contenido_id) REFERENCES contenidos(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (estructura_id) REFERENCES estructuras(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE = InnoDB;