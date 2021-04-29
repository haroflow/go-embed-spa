# go-embed-spa
Example of single binary Golang + embedded VueJS SPA.

Uses `go:embed` and a simple list app showing how to call the backend.

Proof of concept, for educational purposes.

# Requirements

- npm
- go 1.16+ (with go:embed support: https://golang.org/pkg/embed/)

# Run in development mode

Clone the repo:
```
git clone https://github.com/haroflow/go-embed-spa
```

Start the frontend with hot-reload on port 8080:
```
cd go-embed-spa/spa
npm install # First run only
npm run build # First run only, to create spa/dist which is required by go:embed
npm run serve
```
VueJS is configured to proxy requests to 127.0.0.1:8081, to avoid CORS.

On another terminal, Start the backend on port 8081:
```
cd go-embed-spa
go run . -dev
```

It will open the browser on port 8080 (with vue hot-reload).
You should be able to view a list of items and add random items using go as the backend.

Change something inside go-embed-spa/spa/src/App.vue, save the file, the browser should reload automatically.

When changing backend code, re-run `go run . -dev`.

# Build and run 

Tested on Windows: PowerShell seems to have a problem with this, use cmd to build/run the backend.

Tested on WSL: works fine, but the browser can't access localhost:8081, needed to use the IP of the VM.

Build the frontend:
```
cd go-embed-spa/spa
npm install # First run only
npm run build
```

Build the backend, embedding spa/dist folder:
```
cd go-embed-spa
go build .
```

Run:
```
./go-embed-spa
```

It should open the browser on port 8081 (backend serving embedded SPA).
You should be able to view a list of items and add random items using go as the backend.

This binary should contain the whole project, try copying it to another folder/system and running.

# Notes

If you change the default port of the backend (8081), change the address of the proxy inside spa/vue.config.js
