// pkg/router/router.go
package router

import (
	"net/http"
	"todo-backend-golang/internal/handlers"
	"todo-backend-golang/internal/storage"

	"github.com/gorilla/mux"
)

// SetupRouter configura y devuelve el router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Placeholder para verificar que la API funciona
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API funcionando correctamente"))
	}).Methods("GET")

	// Crear almacenamiento en memoria
	memStorage := storage.NewMemoryStorage()

	// Crear manejador de tareas
	taskHandler := handlers.NewTaskHandler(memStorage)

	// Registrar rutas de tareas
	r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTaskByID).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")

	return r
}
