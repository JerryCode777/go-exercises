### objetivo
El objetivo de este ejercicio es crear un proyecto backend usando Golang, La api sera un servicio de
Rest que administre una lista de tareas to-do list. Este servicio permitira crear, leer, actualizar y 
eliminar tareas(CRUD). Las tareas se almacenan en memoria - no base de datos - para simplificar el 
ejemplo, y las pruebas se hacen usando Thunder client.

### EndPoints a desarrollar:
-GET /tasks : Obtiene la lista de tareas
-POST /tasks : Crea una nueva tarea
-GET /tasks/{id} : Obtiene una tarea especifica por su ID
-PUT /tasks/{id} : Actualiza una tarea existente por su ID
-DELETE /tasks/{id} : Elimina una tarea por su ID

### Estructura de datos de la tarea:
JSON
{
    "id": 1,
    "title": "Estudiar Golang",
    "completed": false,
}

### Requerimientos tecnicos
- Leer y utilizar el paquete estandar net/http de Go
- Manejar rutas utilizando gorilla/mux (o el router estandar de Go si prefieres simplicidad)
- Validar los datos enviados por el cliente - Thunder Client
- Devolver respuestas JSON apropiadas con los codigos de estado HTTP correspondientes

### Estructura del proyecto
El proyecto debe seguir la siguiente estructura basica que permita escalabilidad en proyectos futuros

todo-app/
cmd/
    cmd/main.go
internal/
    internal/handlers/task.go
    internal/models/task.go
    internal/storage/memory.go
pkg/
    router/router.go

go.mod
    




