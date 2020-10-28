# Akri: Nessie protocol

Golang implementation of Microsoft's Akri's [Extensibility]() example "Nessie"

This implementation returns an arbitrary []byte rather than taking a URL pointing to a Nessie image.


## Protoc

 ```bash
 protoc \
 --proto_path=./samples/brokers/nessie \
 --go_out=plugins=grpc,module=github.com/DazWilkin/akri-nessie/protos:./golang/nessie/protos \
 ./samples/brokers/nessie/nessie.proto
 ```

 > **NOTE** For temporary convenience, `nessie.proto` exists outside of this repo

## Run

### Server

```bash
go run ./cmd/server --grpc_endpoint=:8083
```

### Client

```bash
go run ./cmd/client --grpc_endpoint=:8083
```
