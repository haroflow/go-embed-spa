# go-embed-spa
Example of single binary Golang + embedded VueJS SPA.

For learning purposes.

Uses go:embed and a simple list app showing how to call the backend.

# Requirements

- npm
- go 1.16+ (with go:embed support: https://golang.org/pkg/embed/)

# Run in development mode

Start the backend on port 8081:
```
cd go-embed-spa
go run .
```

Start the frontend with hot-reload on port 8080:
```
cd go-emded-spa/spa
npm install # First run only
npm run serve
```

It should open the browser on port 8081 (with vue hot-reload).
You should be able to view a list of items and add random items using go as the backend.

Change something inside go-embed-spa/src/App.vue, save the file, the broser should reload automatically.
When changing backend code, re-run `go run .`.

# Build and run 

Tested on Windows only.

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

# Notes

If you change the default port of the backend (8081), change the address of the proxy inside spa/vue.config.js
