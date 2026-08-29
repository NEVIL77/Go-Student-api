# Students API

## How to init the GO module
```bash
go mod init <module-name>
```

---

## Project Folder Structure
`folder = cmd -> students-api`

---

## Run the Server
```bash
go run ./cmd/students-api/main.go
```

---

## IDE Shortcuts

- **Mac:** `cmd + control + p` — open the command palette
- **Windows:** `ctrl + shift + p` — open the command palette

---

## Database

sqlLite is inbuild file database

---

## Packages & Internal

- internal folder carry the project packages

- we have install this package : `go get -u github.com/ilyakaznacheev/cleanenv` : dont know why ?
- use to add the struct tag

---

## `[]byte`

`[]byte` means a slice (list) of bytes.

- `byte` → one byte of data (an integer from 0 to 255).
- `[]` → means multiple values.
- `[]byte` → multiple bytes stored together.

Go converts `"Hello"` into bytes:

```
H → 72
e → 101
l → 108
l → 108
o → 111
```

So conceptually:
```
[]byte("Hello")
[72, 101, 108, 108, 111]
```

### Why is `[]byte` used so much?

Computers ultimately handle data as bytes, so Go uses `[]byte` for things like:

- Reading files
- HTTP request/response bodies
- JSON data
- Network communication
- Images/files