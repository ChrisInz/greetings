package greeting

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juanito"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juanito")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}

//AL PROGRAMAR PRUEBAS ES IMPORTANTE AGREGAR EL ARCHIVO CON _TEST Y QUE LAS FUNCIONES COMIECEN CON TEST, YA QUE DE OTRA FORMA N0 FUNCIONARÁN
//Ademas es importante agregar el parametro *testing.T para que se formateen a pruebas
//Al enviar log.fatalF es el mensaje que se guardara como motivo de rechazo de las pruebas y es como sabe que una prueba fallo
//revisar video 49 del curso de GO si hay dudas
