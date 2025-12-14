# Packages

This folder demonstrates how to organize Go code into reusable packages and how to import both your own packages and third-party ones.

## What's Inside

- `main.go`: Uses
  - Local packages from this module: `auth` and `user` (under `github.com/gunjanghate/learning-go/...`)
  - A third-party package: `github.com/fatih/color` for colored terminal output.
- Focus: Learn how to structure code into packages, import them, and call exported functions.

## Key Concepts

### Why Packages?

- Packages help you follow **DRY** (Don't Repeat Yourself) by grouping related functions and types.
- They improve code organization, reuse, and testing.
- Each folder with a `go.mod` (module root) or Go files becomes a package (with a `package` name at the top of the file).

### Importing Local Packages

```go
import (
    "github.com/gunjanghate/learning-go/auth"
    "github.com/gunjanghate/learning-go/user"
    "github.com/fatih/color"
)
```

- `auth` and `user` are packages within your module.
- Their **module path** is based on the `module` line in your `go.mod` file.

### Using Exported Functions and Types

```go
auth.LoginWithCreds("gg", "gg1234")

session := auth.GetSession()
println("Session ID:", session)

u := user.NewUser("gunjan", "gg1234")
println("Created user:", u.Username)

uss := user.AllUsers()
println("Total users:", len(uss))

for _, usr := range uss {
    color.Green("User: %s, Password: %s", usr.Username, usr.Password)
}
```

- Functions and types that start with an **uppercase letter** are **exported** (public) from a package (`LoginWithCreds`, `GetSession`, `NewUser`, `AllUsers`, `Username`).
- Lowercase identifiers are package-private.

### Third-Party Packages

- `github.com/fatih/color` is used to print colored output to the terminal:

  ```go
  color.Green("User: %s, Password: %s", usr.Username, usr.Password)
  ```

- Such dependencies are managed via `go mod` (`go mod init`, `go get`, `go mod tidy`).

## Run The Example

From PowerShell (assuming `go.mod` is set up at the module root and dependencies are installed):

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\files&packages\packages"
go run main.go
```

You should see login-related output, a printed session ID, created user details, and a list of users printed in green.

## Notes

- Keep package responsibilities small and focused (e.g., `auth` for authentication, `user` for user management).
- Name packages by what they **do**, not by project or technology (`auth`, `user`, `store`, not `util1`, `helpers2`).
- Use Go modules to manage dependencies and import paths cleanly.
- When exposing functionality from a package, choose clear, expressive exported names.
