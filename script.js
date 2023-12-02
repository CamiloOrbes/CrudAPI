function request(method, url, callback, data = null) {
    const xhr = new XMLHttpRequest();
    xhr.open(method, url, true);
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200 || xhr.status === 201) {
                callback(null, JSON.parse(xhr.responseText));
            } else {
                callback(xhr.statusText, null);
            }
        }
    };

    xhr.send(data ? JSON.stringify(data) : null);
}

function listarEstudiantes() {
    request("GET", "http://localhost:8080/estudiante", function (err, data) {
        if (err) {
            document.getElementById("resultado").innerText = "Error al listar estudiantes";
        } else {
            document.getElementById("resultado").innerText = JSON.stringify(data, null, 2);
        }
    });
}

function crearEstudiante() {
    const nuevoEstudiante = {
        nombre: prompt("Ingrese el nombre del nuevo estudiante:"),
        edad: parseInt(prompt("Ingrese la edad del nuevo estudiante:")),
        carrera: prompt("Ingrese la carrera del nuevo estudiante:"),
        semestre: parseInt(prompt("Ingrese el semestre del nuevo estudiante:")),
        materias: parseInt(prompt("Ingrese la cantidad de materias del nuevo estudiante:")),
        activo: confirm("El estudiante está activo?"),
        hobbie: prompt("Ingrese el hobbie del nuevo estudiante:"),
    };

    request("POST", "http://localhost:8080/estudiante", function (err, data) {
        if (err) {
            document.getElementById("resultado").innerText = "Error al crear estudiante";
        } else {
            document.getElementById("resultado").innerText = "Estudiante creado con ID: " + data;
        }
    }, nuevoEstudiante);
}

function editarEstudiante() {
    const estudianteId = prompt("Ingrese el ID del estudiante a editar:");
    request("GET", `http://localhost:8080/estudiante/${estudianteId}`, function (err, estudianteActual) {
        if (err) {
            document.getElementById("resultado").innerText = "Error al obtener la información del estudiante";
        } else {
            const nuevosValores = {
                nombre: prompt(`Nombre actual: ${estudianteActual.nombre}\nIngrese el nuevo nombre del estudiante:`) || estudianteActual.nombre,
                edad: parseInt(prompt(`Edad actual: ${estudianteActual.edad}\nIngrese la nueva edad del estudiante:`)) || estudianteActual.edad,
                carrera: prompt(`Carrera actual: ${estudianteActual.carrera}\nIngrese la nueva carrera del estudiante:`) || estudianteActual.carrera,
                semestre: parseInt(prompt(`Semestre actual: ${estudianteActual.semestre}\nIngrese el nuevo semestre del estudiante:`)) >>> 0 || estudianteActual.semestre,
                materias: parseInt(prompt(`Materias actuales: ${estudianteActual.materias}\nIngrese la nueva cantidad de materias del estudiante:`)) >>> 0 || estudianteActual.materias,
                activo: confirm(`Activo actual: ${estudianteActual.activo}\nEl estudiante está activo?`) || estudianteActual.activo,
                hobbie: prompt(`Hobbie actual: ${estudianteActual.hobbie}\nIngrese el nuevo hobbie del estudiante:`) || estudianteActual.hobbie,
            };


            request("PATCH", `http://localhost:8080/estudiante/${estudianteId}`, function (err) {
                if (err) {
                    document.getElementById("resultado").innerText = "Error al editar estudiante";
                } else {
                    document.getElementById("resultado").innerText = "Estudiante editado con éxito";
                    listarEstudiantes();  
                }
            }, nuevosValores);
        }
    });
}

function eliminarEstudiante() {
    const estudianteId = prompt("Ingrese el ID del estudiante a eliminar:");

    request("DELETE", `http://localhost:8080/estudiante/${estudianteId}`, function (err) {
        if (err) {
            document.getElementById("resultado").innerText = "Error al eliminar estudiante";
        } else {
            document.getElementById("resultado").innerText = "Estudiante eliminado con éxito";
        }
    });
}

