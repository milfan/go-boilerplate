## Boilerplate Arch
This is the architecture in this boilerplate

```
├── .env.example
├── .gitignore
├── Makefile
├── api
    └── rest
    │   ├── routes
    │       └── web
    │       │   └── v1
├── cmd
    ├── cli
    └── rest
├── configs
    ├── config
    ├── constants
    ├── middleware
    └── postgres
├── go.mod
├── go.sum
├── internal
    ├── api
    │   ├── controllers
    │   │   └── web
    │   ├── entities
    │   ├── errors
    │   ├── helpers
    │   ├── models
    │   ├── presenters
    │   │   └── requests
    │   ├── repositories
    │   └── usecases
    │   │   └── web
    └── cli
    │   ├── commands
    │   ├── repositories
    │   └── usecases
├── migrations
└── pkg
    ├── constants
    ├── errors
    ├── log
    └── response
```