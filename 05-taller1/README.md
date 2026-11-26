# 05-taller1: Estadísticas de calificaciones con función variádica

## Descripción
Este programa define la función variádica estadisticas(nums ...float64) que recibe cualquier cantidad de calificaciones y retorna cuatro valores nombrados: menor, mayor, promedio y err (error). Cuando len(nums) == 0, asigna el error centinela ErrSinDatos y ejecuta return desnudo.

## Cómo ejecutar
Desde la carpeta 05-taller1, ejecutar:
go run .

El programa imprime tres líneas:
1. Llamada con valores sueltos (8.5, 6, 9.75, 4): muestra el menor, el mayor, el promedio exacto y <nil>.
2. Llamada con slice desempaquetado (notas...): muestra el menor, el mayor y el promedio formateados a dos decimales (%.2f).
3. Llamada sin argumentos (estadisticas()): comprueba el error con errors.Is(err, ErrSinDatos) e imprime error: no se recibieron datos.

## Pruebas realizadas
go run .
4 9.75 7.0625 <nil>
min=7.00 max=10.00 prom=8.17
error: no se recibieron datos

## Preguntas

### a) ¿Por qué los resultados se llaman menor y mayor y no min y max?
Porque desde Go 1.21 existen las funciones integradas min y max. En func estadisticas(nums ...float64) (menor, mayor, promedio float64, err error), usar menor y mayor evita ocultar (shadowing) esas funciones integradas dentro del cuerpo de estadisticas. Si los resultados se llamaran min y max, al ejecutar min(menor, n) y max(mayor, n) dentro del bucle for Go intentaría llamar a variables float64 como funciones y fallaría al compilar.

### b) ¿Qué hacen los tres puntos en estadisticas(notas...)? ¿Qué mensaje da Go si los quitas y por qué?
Los tres puntos (...) en estadisticas(notas...) desempaquetan los elementos de notas ([]float64) para pasarlos como argumentos float64 sueltos a nums ...float64. Si quitas los tres puntos y escribes estadisticas(notas), Go devuelve un error de tipo porque estás pasando un slice ([]float64) donde se esperan valores float64 individuales.

- [05-taller1](./05-taller1): Función variádica en Go para calcular nota menor, mayor, promedio y validar error centinela con errors.Is.
