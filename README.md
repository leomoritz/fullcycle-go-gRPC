# Como começar a usar gRPC com go
* https://grpc.io/docs/languages/go/quickstart/

# Comando para gerar os fontes grpc e protocol buffers 
* protoc --go_out=. --go-grpc_out=. proto/course_category.proto
OBS: Utiliza o pb e os plugins instalados com base nas orientações do link anterior

# Comando para download das dependências
go mod tidy

# Comando para executar o server
go run cmd/grpcServer/main.go

# Comando para executar o client evans
* https://github.com/ktr0731/evans
* Comando: evans -r repl
OBS: Para usar esse comando foi necessário implementar reflection.Register(grpcServer) na classe main.go

# Comando para salvar as dependências no PATH do go
export PATH="$PATH:$(go env GOPATH)/bin" 

# Observações finais
Em alguns casos, pode ser necessário utilizar o sudo.


