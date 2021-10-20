# :green_heart: _platzi_: Curso básico de programación con Go
<div>
    <img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" alt="Go" width="70%" />
	<img src="https://go.dev/images/gophers/blue.svg" alt="Go" width="25%" />
</div>
<br />

**Go** es un lenguaje de programación concurrente y compilado inspirado en la sintaxis de C, que intenta ser dinámico como python y con el rendimiento de C o C++.

`go build`: Generar un binario ejecutable para nuestra aplicación Go, lo que le permitirá distribuir e implementar el programa donde lo desee.

**Ejemplo:**
```go
go build src/main.go
```
------
`go run`: Ejecuta nuestra aplicación Go y no genera un archivo binario ejecutable.

**Ejemplo:**
```go
go run src/main.go
```

## Constantes, variables y zero values
### Constantes
Las constantes mantienen siempre un valor y este nunca puede ser cambiado. Si se intenta hacerlo, el compilador mostrará un error.

```go
const pi float64 = 3.14
const pi2 = 3.1415

fmt.Println("pi:" pi)
fmt.Println("pi2:" pi2)
```

### Variables
Una variable es una referencia a un valor, lo que permite utilizarle para construir instrucciones lógicas dentro de un programa.
```go
base := 12 // se conoce como declaración corta de variable
var altura int = 14 // se conoce como declaración larga de variable
var area int // Variable llamada `area` del tipo de datos `int` sin iniciarla. Esto significa que declararemos un espacio para ubicar un valor, pero no le daremos un valor inicial.
```
**Importante:** Todas la variables que se declaren en Go se deben de usar.

### Zero values
A diferencia de otros lenguajes en donde una variable sin valor asignado es `null` o `undefined`, las variables tienen por default un valor asignado en Go. Este es conocido como el valor cero (Zero values).
```go
var nombre string
var edad int
var peso float64
var estudiante bool
fmt.Println("Nombre:", nombre)
fmt.Println("Edad:", edad)
fmt.Println("Peso:", peso)
fmt.Println("Estudiante:", estudiante)
```