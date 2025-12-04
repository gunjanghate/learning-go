# Variables & Constants

This folder explores declaring variables and constants in Go, including type inference and best practices.

## What’s Inside

- `main.go`: Examples of `var` declarations, short declaration `:=`, typed vs inferred variables, and `const` usage.
- Focus: How and when to use variables vs constants, and writing clear, idiomatic code.

## Key Concepts

- `var` vs `:=`: Use `:=` for local, inferred types; `var` for package-level or explicit type
- Constants: `const` for values known at compile time; untyped constants can adapt to context
- Type inference: Go infers types on `:=`; explicit types improve clarity when needed

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\variables&constants"
go run main.go
```

## Notes

- Prefer immutable values with `const` when appropriate.
- Keep examples small and focused to highlight one concept at a time.

## Examples 

नीचे बहुत ही बुनियादी उदाहरण दिए गए हैं ताकि सिंटैक्स समझ में आए। विस्तृत कोड पहले से `main.go` में मौजूद है।

### Variables 

```go
package main

import "fmt"

func main() {
	var name string = "Gunjan"
	var age int = 25

	city := "Pune"

	fmt.Println(name, age, city)
}
```

### Constants 

```go
package main

import "fmt"

const(
    AppName = "Learning Go"
    Pi = 3.14
    )


func main() {
	fmt.Println(AppName, Pi)
}
```
