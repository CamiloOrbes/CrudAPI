document.addEventListener('DOMContentLoaded', function () {
    const btnLeerEstudiantes = document.getElementById('btnLeerEstudiantes');
    const btnCrearEstudiante = document.getElementById('btnCrearEstudiante');
    const btnGuardarEstudiante = document.getElementById('btnGuardarEstudiante');
    const btnActualizarEstudiante = document.getElementById('btnActualizarEstudiante');
    const btnEliminarEstudiante = document.getElementById('btnEliminarEstudiante');
    const resultadoDiv = document.getElementById('resultado');
    const crearEstudianteForm = document.getElementById('crearEstudianteForm');
    const inputIdActualizar = document.getElementById('inputIdActualizar');
    const inputIdEliminar = document.getElementById('inputIdEliminar');

    btnLeerEstudiantes.addEventListener('click', async function () {
        try {
            const response = await fetch('http://localhost:8080/estudiante');
            const estudiantes = await response.json();

            // Mostrar los estudiantes en el resultadoDiv
            mostrarResultado(estudiantes);
        } catch (error) {
            manejarError('Error al leer estudiantes:', error);
        }
    });

    btnCrearEstudiante.addEventListener('click', function () {
        crearEstudianteForm.style.display = 'block';
    });

    btnGuardarEstudiante.addEventListener('click', async function () {
        const nuevoEstudiante = obtenerDatosFormulario();
        try {
            const nuevoEstudianteCreado = await realizarSolicitud('POST', 'http://localhost:8080/estudiante', nuevoEstudiante);
            mostrarResultado([nuevoEstudianteCreado]);
            limpiarFormulario();
        } catch (error) {
            manejarError('Error al crear estudiante:', error);
        }
    });

    btnActualizarEstudiante.addEventListener('click', async function () {
        const idActualizar = inputIdActualizar.value;
        if (idActualizar.trim() === '') {
            alert('Ingrese un ID para actualizar.');
            return;
        }

        try {
            const estudiante = await obtenerEstudiantePorId(idActualizar);
            mostrarResultado([estudiante]);
        } catch (error) {
            manejarError('Error al obtener estudiante para actualizar:', error);
        }
    });

    btnEliminarEstudiante.addEventListener('click', async function () {
        const idEliminar = inputIdEliminar.value;
        if (idEliminar.trim() === '') {
            alert('Ingrese un ID para eliminar.');
            return;
        }

        try {
            await realizarSolicitud('DELETE', `http://localhost:8080/estudiante/${idEliminar}`);
            mostrarResultado([{ mensaje: `Estudiante con ID ${idEliminar} eliminado.` }]);
        } catch (error) {
            manejarError('Error al eliminar estudiante:', error);
        }
    });

    function mostrarResultado(datos) {
        resultadoDiv.innerHTML = '<h2>Resultado</h2>';
        datos.forEach(dato => {
            resultadoDiv.innerHTML += `<p>${JSON.stringify(dato)}</p>`;
        });
    }

    function manejarError(mensaje, error) {
        console.error(mensaje, error);
        resultadoDiv.innerHTML = `<p>${mensaje}</p>`;
    }

    async function realizarSolicitud(metodo, url, cuerpo) {
        const opciones = {
            method: metodo,
            headers: {
                'Content-Type': 'application/json'
            },
            body: cuerpo ? JSON.stringify(cuerpo) : undefined
        };

        const response = await fetch(url, opciones);
        return await response.json();
    }

    function obtenerDatosFormulario() {
        return {
            nombre: document.getElementById('nombre').value,
            edad: parseInt(document.getElementById('edad').value),
            carrera: document.getElementById('carrera').value,
            semestre: parseInt(document.getElementById('semestre').value),
            materias: parseInt(document.getElementById('materias').value),
            activo: document.getElementById('activo').checked,
            hobbie: document.getElementById('hobbie').value
        };
    }

    async function obtenerEstudiantePorId(id) {
        const url = `http://localhost:8080/estudiante/${id}`;
        const response = await fetch(url);
        return await response.json();
    }

    function limpiarFormulario() {
        crearEstudianteForm.style.display = 'none';
        document.getElementById('nombre').value = '';
        document.getElementById('edad').value = '';
        document.getElementById('carrera').value = '';
        document.getElementById('semestre').value = '';
        document.getElementById('materias').value = '';
        document.getElementById('activo').checked = true;
        document.getElementById('hobbie').value = '';
    }
});

