# 04-taller1 - FizzBuzz con switch sin condición

## Descripción
Programa en Go que evalúa los números del 1 al 30 mediante un switch sin condición para imprimir FizzBuzz (múltiplos de 15), Fizz (múltiplos de 3), Buzz (múltiplos de 5) o el número correspondiente, mostrando al final los contadores.

## Cómo ejecutar
```bash
go run .
```

## Resultado

### Salida normal
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz
Fizz=8 Buzz=4 FizzBuzz=2

### Experimento (case i%3 == 0 antes que case i%15 == 0)
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 Fizz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 Fizz
Fizz=10 Buzz=4 FizzBuzz=0

## Preguntas
- **a)** Al poner case i%3 == 0 antes que case i%15 == 0, los múltiplos de 15 (15 y 30) también cumplen i%3 == 0 y entran primero en Fizz. Por eso FizzBuzz queda en 0 y Fizz sube a 10 (Buzz sigue en 4).
- **b)** Cuando i vale 15 se cumplen las tres condiciones (múltiplo de 15, 3 y 5), pero Go ejecuta solo el primer case verdadero (i%15 == 0). No hace falta break porque en Go los case no tienen fallthrough automático.
