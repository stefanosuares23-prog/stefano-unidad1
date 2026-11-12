# Tarea 03: Análisis de Ámbito y Sombreado en Go

- **Autor:** Stefano de Jesús Suárez Elizalde

## 1. ¿Qué imprime?
El programa imprime por consola:
total: 0
## 2. Explicación del funcionamiento (Ámbito)
Ocurre un problema de sombreado de variables (*variable shadowing*). Dentro del ciclo `for`, al utilizar el operador de declaración corta `:=` en la línea `total := total + i`, se está declarando una **nueva variable local** cuyo **ámbito** (*scope*) se limita estrictamente al bloque interno del bucle. Esta nueva variable oculta a la variable `total` externa que fue declarada al inicio de la función `suma()`. Como resultado, la variable `total` que la función retorna nunca es modificada dentro del ciclo y se mantiene con su valor inicial de `0`.
go

## 3. Corrección cambiando un solo carácter
Para solucionarlo, se debe cambiar el operador de declaración corta `:=` por el operador de asignación `=` en la línea del ciclo:
total = total + i
## 4. Extra: Ejecución de go vet y el analizador shadow
- **¿Lo detecta go vet por defecto?** Al ejecutar el comando estándar `go vet ./...`, es posible que no advierta este sombreado a menos que se invoque explícitamente el analizador secundario o se configure el flag correspondiente para revisión de ámbito.
- **Investigación del analizador shadow:** El analizador `shadow` es una herramienta estática integrada en el ecosistema de Go que examina el código fuente en busca de declaraciones de variables que ocultan (*shadow*) a otras variables existentes en un ámbito exterior. Su objetivo principal es alertar a los desarrolladores sobre errores lógicos sutiles causados por la reutilización involuntaria de nombres de variables en bloques anidados.
