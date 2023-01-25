## Go Expert - Clean Architecture

Clean architecture com Go

## Como utilizar

- Clonar repositório
- (Opcional) Executar os testes
- Inicializar serviços(docker)
    - RabbitMQ
    - Mysql
    - Aplicação Go


```console

git clone https://github.com/virb30/goexpert-cleanarch.git .
cd goexpert-cleanarch
go test ./...
docker-compose up -d

```

### Endpoints REST

- `POST localhost:8000/order` - criar um novo pedido
- `GET localhost:8000/orders` - obter a lista de pedidos

### Playground GraphQL

- `localhost:8080`

### Serviços gRPC

- `CreateOrder` - criar um novo pedido
- `ListOrders` - listar pedidos

### Acesso RabbitMQ

```
endpoint: http://localhost:15672

usuário: guest
senha: guest

```
