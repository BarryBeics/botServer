run with docker compose up
take down with docker compose down --volumes

to generate new graph code

go get github.com/99designs/gqlgen/codegen/config@v0.17.40
go get github.com/99designs/gqlgen/internal/imports@v0.17.40
go get github.com/99designs/gqlgen@v0.17.40
go run github.com/99designs/gqlgen

new version
docker build -t backend:0.62 .

remember to update the .env file before doing docker compose up
make sure you dont copy over the bson values in models_gen.go