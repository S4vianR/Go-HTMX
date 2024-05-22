package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Alumno struct {
	Nombre string
	NumeroControl string
	Carrera string
}

func main() {
	port := ":8080"
	host := "http://localhost"+port
	handlerRoot := func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("index.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        alumnos := map[string][]Alumno{
            "Alumnos": {
                {Nombre: "Juan", NumeroControl: "123", Carrera: "Ingeniería"},
                {Nombre: "Ana", NumeroControl: "456", Carrera: "Matemáticas"},
                {Nombre: "Pedro", NumeroControl: "789", Carrera: "Física"},
            },
        }

        err = tmpl.Execute(w, alumnos)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }

    handlerAlumnos := func(w http.ResponseWriter, r *http.Request) {
        nombre := r.PostFormValue("nombre")
        numeroControl := r.PostFormValue("numeroControl")
        carrera := r.PostFormValue("carrera")
        htmlStrg := fmt.Sprintf("<li> %s - %s - %s</li>", nombre, numeroControl, carrera)
        tmpl, err := template.New("alumnos").Parse(htmlStrg)
        tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

	http.HandleFunc("/", handlerRoot)
    http.HandleFunc("/alumnos", handlerAlumnos)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Server running on", host)
    log.Fatal(http.ListenAndServe(port, nil))
}