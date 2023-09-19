# ---------------------------------------------------
Aqui ó, se crean archivos .go que seran asociados a los grupos de rutas
Se recomiendan que los archivos tengan nombres representativos.
Ej:
# fronted.go  --> Rutas que tienen que ver con un front
Ej: 
# service.go  --> Rutas que tienen que ver con un Servicio

Ej: 
# db.go  --> Rutas que tienen que ver con una bd
# ---------------------------------------------------
    Auqi ó, cada archivo tendra una funcion con el nombre de archivo(en minuscula para que
    sea privada) que Recivira por parametro 
# *mux.Router 
    que sera el Objeto encargado de Manejar los Handler
# loadRoute
    Es el objeto pasado por parametro a la funcion con nombre igual al del fichero 
# eje: 
el cuerpo seria algo como esto:
    loadRoute.HandleFunc("/(Ruta)/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./page/index.html") // Sirve el archivo HTML 
    })