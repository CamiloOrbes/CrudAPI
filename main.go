package main

import (
	"log"
	"net/http"

	"github.com/CamiloOrbes/CrudAPI/controllers"
	"github.com/CamiloOrbes/CrudAPI/handlers"
	"github.com/CamiloOrbes/CrudAPI/models"
	repositorio "github.com/CamiloOrbes/CrudAPI/repository"
	gorillaHandlers "github.com/gorilla/handlers" // Importa el paquete gorilla/handlers para CORS
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

/*
función para conectarse a la instancia de PostgreSQL, en general sirve para cualquier base de datos SQL.
Necesita la URL del host donde está instalada la base de datos y el tipo de base datos (driver)
*/
func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	/* creando un objeto de conexión a PostgreSQL */
	db, err := ConectarDB("postgres://cfmjhlfj:bmbhMqaU6Ux9unepmvxgRIJ6n1iLq_TL@snuffleupagus.db.elephantsql.com/cfmjhlfj", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/* creando una instancia del tipo Repository del paquete repository
	se debe especificar el tipo de struct que va a manejar la base de datos
	para este ejemplo es Estudiante y se le pasa como parámetro el objeto de
	conexión a PostgreSQL */
	repo, err := repositorio.NewRepository[models.Estudiante](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := handlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter()

	/* rutas a los endpoints de la API */
	router.Handle("/estudiante", http.HandlerFunc(handler.LeerEstudiante)).Methods(http.MethodGet)
	router.Handle("/estudiante", http.HandlerFunc(handler.CrearEstudiante)).Methods(http.MethodPost)
	router.Handle("/estudiante/{id}", http.HandlerFunc(handler.LeerUnEstudiante)).Methods(http.MethodGet)
	router.Handle("/estudiante/{id}", http.HandlerFunc(handler.ActualizarUnEstudiante)).Methods(http.MethodPatch)
	router.Handle("/estudiante/{id}", http.HandlerFunc(handler.EliminarUnEstudiante)).Methods(http.MethodDelete)

	headers := gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := gorillaHandlers.AllowedOrigins([]string{"*"}) // Ajusta esto según tus necesidades de seguridad

	// Agrega los manejadores CORS a tu enrutador
	handlerWithCORS := gorillaHandlers.CORS(headers, methods, origins)(router)

	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router */
	http.ListenAndServe(":8080", handlerWithCORS)
}
