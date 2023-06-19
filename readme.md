#### Go setup
    
```bash
go mod init github.com/username/projectname
go mod init github.com/drsimplegraffiti/gojwt
```

#### Install dependencies

```bash
go get github.com/dgrijalva/jwt-go
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/joho/godotenv
go get -u github.com/golang-jwt/jwt/v4
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon@latest
export PATH="$PATH:$GOPATH/bin"

```

#### Run with CompileDaemon

```bash
compiledaemon --command="./gojwt"
```

The `./gojwt` is the name of your project module

