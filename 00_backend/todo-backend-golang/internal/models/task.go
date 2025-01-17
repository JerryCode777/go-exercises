package models

//task representa una tarea en el sistema

type Task struct {
	ID        int    `json:id`        //Identificador unico de tares
	Title     string `json:title`     //Titulo de la tarea
	Completed bool   `json:completed` //estado de la tarea = completada o no
}
