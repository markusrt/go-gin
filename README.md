# Go example app using Gin web framework



### Building

```bash
go build -o app
./app
```

### Accessing the app

Either via browser on <http://localhost:8080> or via `curl`:

```bash
curl -X GET -H "Accept: application/json" http://localhost:8080/articles
curl -X GET -H "Accept: application/json" http://localhost:8080/articles/1
```

## Testing

```bash
go test -v
```

## Deploy to Tanzu Platform

```bash
tanzu apps workload apply go-gin --git-repo https://github.com/markusrt/go-gin --git-branch main --type web --build-env "BP_KEEP_FILES=templates/*"
```