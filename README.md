# TodoList

## Tech stack
- [Golang](https://golang.org/)
- [Gorm](https://gorm.io/)
- [JWT](https://jwt.io/)
- [PostgreSQL](https://www.postgresql.org/)

## Config

copy `./cmd/configExample.yaml to ./.local/config.yaml`

## Run TodoList

```
go run cmd/main.go
```

## Project structure

```
│   .gitignore
│   go.mod
│   go.sum
│   README.md
│
├───.idea
│       .gitignore
│       modules.xml
│       Todo_List.iml
│       vcs.xml
│       workspace.xml
│
├───cmd
│       configExample.yaml
│       main.go
│
├───config
│       config.go
│
├───helper
│       config.go
│       error.go
│       jwtClaims.go
│       utility.go
│
├───storage
│       db.go
│       s_todo.go
│       s_user.go
│
└───web
    │   route.go
    │   web.go
    │   w_auth.go
    │   w_todo.go
    │
    └───potal
            p_auth.go
            p_item.go
```