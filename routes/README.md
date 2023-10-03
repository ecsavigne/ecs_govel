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
# eje: Esta ruta carga archivo html u otro que quiera ser servido
    loadRoute.HandleFunc("/(Ruta)/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./page/index.html") // Sirve el archivo HTML 
    })

# ej: Creacion de una ruta de servicio
    loadRoute.HandleFunc("/(Ruta)/", func(w http.ResponseWriter, r *http.Request) {
        // Cabecera permitidas por la ruta cords, todas las que se quieran
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Content-Type", "application/text")
        
        // vars := mux.Vars(r) --> Variables que son pasada por la cabecera de la url, o son tratada como variables en la url de la ruta ej: 
            1. http://server.com/test?a=4&b=5
            2. "/(Ruta)/{var}/{varN}"

        //Carga variables enviada via form-data
        r.FormValue("") 
        r.FormFile("Multipar")
        // y mucho mas que ver ayuda http.Request


        //Logica del codigo

        // devolucucion de la respuesta via json y via texto
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
        //w.Write([]byte("<h1>Registro de curso</h1>"))
    })
