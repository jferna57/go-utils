# Go-Utils

Este es un simple paquete en Go que proporciona una función para saludar.

## Instalación

Para usar este modulo, necesitas tener Go instalado en tu sistema. A continuación, puedes instalar el paquete con el siguiente comando:

```bash
go get "github.com/jferna57/go-utils"
```

## Uso
Para usar el paquete en tu código, importa el paquete greetings y llama a la función Hello:


```go
package main

import (
	"fmt"
	"github.com/jferna57/go-utils"
)

func main() {
	message := greetings.Hello("John")
	fmt.Println(message) // Salida: "Hello, John!"
}
```

## Contribuir

Si encuentras algún problema o tienes alguna sugerencia para mejorar este proyecto, no dudes en abrir un [issue](https://github.com/jferna57/
go-utils/issues) o enviar un [pull request](https://github.com/jferna57/go-utils/pulls).
