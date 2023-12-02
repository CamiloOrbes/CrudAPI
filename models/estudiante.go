package models

/*
es mejor conservar un estándar entre las etiquetas de json y db para no tener problemas al parsear
de json a db en el método ActualizarUnEstudiante
*/
type Estudiante struct {
	Id       int    `db:"id" json:"id"`
	Nombre   string `db:"nombre" json:"nombre"`
	Edad     uint   `db:"edad" json:"edad"`
	Carrera  string `db:"carrera" json:"carrera"`
	Semestre int    `db:"id" json:"semestre"`
	Materias int    `db:"id" json:"materias"`
	Activo   bool   `db:"esta_cerca" json:"activo"`
	Hobbie   string `db:"hobbie" json:"hobbie"`
}
