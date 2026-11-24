## 03-tarea3 - Conversor de almacenamiento en Go

- **Problema 1:** Al ejecutar `convertir.exe` directamente en PowerShell no se reconocia el comando por seguridad de rutas relativas, y al intentar el primer commit Git bloqueo la operacion por falta de `user.name` y `user.email`.
- **Solucion:** Ejecutar el binario anteponiendo `.\` (`.\convertir.exe 1024 B KB`), configurar la identidad global de Git con el correo de GitHub y excluir `convertir.exe` del control de versiones.
