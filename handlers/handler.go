package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CamiloOrbes/CrudAPI/controllers"
	"github.com/gorilla/mux"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) ActualizarUnEstudiante(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un estudiante, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.controller.ActualizarUnEstudiante(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un estudiante, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) EliminarUnEstudiante(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	err := h.controller.EliminarUnEstudiante(id)
	if err != nil {
		log.Printf("fallo al eliminar un estudiante, con error: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("fallo al eliminar un estudiante con id %s", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) LeerUnEstudiante(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	estudiante, err := h.controller.LeerUnEstudiante(id)
	if err != nil {
		log.Printf("fallo al leer un estudiante, con error: %s", err.Error())
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el estudiante con id %s no se pudo encontrar", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(estudiante)
}

func (h *Handler) LeerEstudiante(writer http.ResponseWriter, req *http.Request) {
	estudiante, err := h.controller.LeerEstudiante(100, 0)
	if err != nil {
		log.Printf("fallo al leer estudiante, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los estudiante", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(estudiante)
}

func (h *Handler) CrearEstudiante(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo estudiante", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	nuevoId, err := h.controller.CrearEstudiante(body)
	if err != nil {
		log.Println("fallo al crear un nuevo estudiante, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo estudiante", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("Id nuevo estudiante: %d", nuevoId)))
}
