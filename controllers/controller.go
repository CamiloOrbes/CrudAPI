package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/CamiloOrbes/CrudAPI/models"
	repositorio "github.com/CamiloOrbes/CrudAPI/repository"
)

var (
	updateQuery = "UPDATE estudiante SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM estudiante WHERE id=$1;"
	selectQuery = "SELECT id, nombre, edad, carrera, semestre, materias, activo, hobbie FROM estudiante WHERE id=$1;"
	listQuery   = "SELECT id, nombre, edad, carrera, semestre, materias, activo, hobbie FROM estudiante limit $1 offset $2"
	createQuery = "INSERT INTO estudiante (id, nombre, edad, carrera, semestre, materias, activo, hobbie) VALUES (:nombre, :edad, :carrera, :semestre, :materias, :activo, :hobbie) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Estudiante]
}

func NewController(repo repositorio.Repository[models.Estudiante]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarUnEstudiante(reqBody []byte, id string) error {
	nuevosValoresEstudiante := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresEstudiante)
	if err != nil {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un estudiante, con error: %s", err.Error())
	}

	if len(nuevosValoresEstudiante) == 0 {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un estudiante, con error: %s", err.Error())
	}

	query := construirUpdateQuery(nuevosValoresEstudiante)
	nuevosValoresEstudiante["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresEstudiante)
	if err != nil {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un estudiante, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]any) string {
	columns := []string{}
	for key := range nuevosValores {
		columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
	}
	columnsString := strings.Join(columns, ",")
	return fmt.Sprintf(updateQuery, columnsString)
}

func (c *Controller) EliminarUnEstudiante(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un estudiante, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un estudiante, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnEstudiante(id string) ([]byte, error) {
	estudiante, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un estudiante, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un estudiante, con error: %s", err.Error())
	}

	estudianteJson, err := json.Marshal(estudiante)
	if err != nil {
		log.Printf("fallo al leer un estudiante, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un estudiante, con error: %s", err.Error())
	}
	return estudianteJson, nil
}

func (c *Controller) LeerEstudiante(limit, offset int) ([]byte, error) {
	estudiante, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer estudiante, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer estudiante, con error: %s", err.Error())
	}

	jsonEstudiante, err := json.Marshal(estudiante)
	if err != nil {
		log.Printf("fallo al leer estudiante, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer estudiante, con error: %s", err.Error())
	}
	return jsonEstudiante, nil
}

func (c *Controller) CrearEstudiante(reqBody []byte) (int64, error) {
	nuevoEstudiante := &models.Estudiante{}
	err := json.Unmarshal(reqBody, nuevoEstudiante)
	if err != nil {
		log.Printf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
	}

	valoresColumnasNuevoEstudiante := map[string]any{
		"nombre":   nuevoEstudiante.Nombre,
		"edad":     nuevoEstudiante.Edad,
		"carrera":  nuevoEstudiante.Carrera,
		"semestre": nuevoEstudiante.Semestre,
		"materias": nuevoEstudiante.Materias,
		"activo":   nuevoEstudiante.Activo,
		"hobbie":   nuevoEstudiante.Hobbie,
	}

	nuevoId, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoEstudiante)
	if err != nil {
		log.Printf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
	}
	return nuevoId, nil
}
