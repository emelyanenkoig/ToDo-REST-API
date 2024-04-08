# INSTALL GIN FRAMEWORK (для работы с http)
go get -u github.com/gin-gonic/gin  

# INSTALL VIPER FRAMEWORK (для работы с конфигами)
go get -u github.com/spf13/viper

# Docker postgres image
docker pull postgres
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres 

# MIGRATIONS 
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

# SETTING ENVIROMENT
go get github.com/joho/godotenv


# LOGGING LIBRARY
go get -u github.com/sirupsen/logrus

# JWT LIBRARY
go get -u github.com/dgrijalva/jwt-go