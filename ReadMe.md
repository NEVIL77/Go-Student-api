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


- How to start the project : 
go run cmd/students-api/main.go -config config/local.yaml

- gress stutdown : when something happen and server is shutdown then it also termoinate the ongoing request as well . 
but in the production we dont want this , we want to server the ongoing request . we cant stop the request in the middle .

for that we will use signal with os channel

done := make(chan os.Signal, 1)       -> making channel which store the os signal (ctrl +c) . 
signal.notify(done, os.Interrupt, syscall.SIGTERM)   -> signal notified channel that something for os signal is triggered and it store the info in the channel and what kind of thing is there then it store like os.Interrupt and syscall.SIGTERM .


gresStutDonw

	slog.Info("Server started", slog.String("adress", cfg.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to start server")
		}

	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}
	slog.Info("server gracefully stopped")


client validation package : go get github.com/go-playground/validator/v10