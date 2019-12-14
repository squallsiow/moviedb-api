#### moviedb-api

### Installation

```bash
export GO111MODULE=on

go run main.go

go get github.com/oxequa/realize

 realize start
```

### Tools

Tools used for this development :

1. [Golang](https://golang.org/dl/)
2. [VS Code](https://code.visualstudio.com/download)
3. [POSTMAN](https://www.getpostman.com/)

### Available API
## Retrieve and store Movie information into local server
# Header required for ADMIN :
X-Authenticati : CflFPa89BzSiVdamikDavDBpKtC9A2zk

API [PUT]       : http://localhost:8080/admin/movie 
```json
{
        "ID" : 550
}
```

## Get movie information from local
# movieID as parameter
API [GET]       : http://localhost:8080/movie/:movieID

## Show all available movies in datastore
API GET       : http://localhost:8080/showall

