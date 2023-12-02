document.addEventListener('DOMContentLoaded', function () {
    const btnLeerEstudiantes = document.getElementById('btnLeerEstudiantes');
    const resultadoDiv = document.getElementById('resultado');

    btnLeerEstudiantes.addEventListener('click', async function () {
        try {
            const response = await fetch('http://localhost:8080/estudiante');
            const estudiantes = await response.json();

            // Mostrar los estudiantes en el resultadoDiv
            resultadoDiv.innerHTML = '<h2>Estudiantes</h2>';
            estudiantes.forEach(estudiante => {
                resultadoDiv.innerHTML += `<p>${JSON.stringify(estudiante)}</p>`;
            });
        } catch (error) {
            console.error('Error al leer estudiantes:', error);
            resultadoDiv.innerHTML = '<p>Error al leer estudiantes</p>';
        }
    });
});