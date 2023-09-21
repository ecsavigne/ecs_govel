# Nombres de Modelo
 si la migracion=nombre_otronombre y archivo de migration(file= nombre_otronombreMigration.sql).
 LA modelo debe de Llamarse NombreOtronombre y se ecribiria la estructura golang
 # nombre estructura golang segun noombre migrate y file.sql
    ---------------NombreOtronombreModelo.go----------------
    type NombreOtronombre struct{
        gorm.Model
        name string
        DirCompleja string
    }

.sql equivalente para la modelo:
    CREATE TABLE IF NOT EXISTS nombre_otronombre (
        --gorm.Model hasta delet_at timestamp------------
        id INT AUTO_INCREMENT PRIMARY KEY,
        created_at DATETIME,
        updated_at DATETIME,
        deleted_at DATETIME,
        --------------Fin *gorm.Model---------------
        name VARCHAR(255),
        dir_compleja VARCHAR(255)
    ) ENGINE = InnoDB;

# Relaciones entre modelos
# has Many( A one to Many B)
 ej: A tiene una relacion 1 to Mucho B
 type A struct {
    gorm.Model
    Bs []B
 }  
 
 type B struct {
    gorm.Model
    AID int
 }  
# Belong (B belong to A)
ej: B Pertenece a A ( 1 B pertenece a un unico A)
type A struct {
    gorm.Model
 }  
 
 type B struct {
    gorm.Model
    AID int
    B B
 }