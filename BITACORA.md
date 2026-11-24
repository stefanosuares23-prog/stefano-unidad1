## 03-tarea3 - Conversor de almacenamiento en Go

- **Problema 1:** Al ejecutar `convertir.exe` directamente en PowerShell no se reconocia el comando por seguridad de rutas relativas, y al intentar el primer commit Git bloqueo la operacion por falta de `user.name` y `user.email`.
- **Solucion:** Ejecutar el binario anteponiendo `.\` (`.\convertir.exe 1024 B KB`), configurar la identidad global de Git con el correo de GitHub y excluir `convertir.exe` del control de versiones.

## 04-taller1
- En un switch sin condición en Go se evalúan los case en orden de arriba hacia abajo y se ejecuta el primero verdadero sin requerir break; moví analizador.go fuera de 04-taller1 para evitar main redeclared in this block.

## 03-tarea4
- Probada versión monolito con los 8 casos base (códigos 0, 2, 3 y 1).

## 03-tarea4
- Probada versión monolito con los 8 casos base (códigos 0, 2, 3 y 1).

- Implementado paquete internal/almacenamiento y verificada salida de go doc ./internal/almacenamiento.
