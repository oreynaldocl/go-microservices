## Create Proto
Run following line
```shell
#protos:
# YOUTUBE EXAMPLE https://www.youtube.com/watch?v=pMgty_RYIOc&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_&index=13&ab_channel=NicJackson
protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency
# OFFICIAL https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/currency.proto
# adding require_unimplemented_servers=false https://github.com/grpc/grpc-go/issues/3794#issuecomment-725860916
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false protos/currency.proto
```
The difference could be by the version difference

### Error (missing mustEmbedUnimplementedCurrencyServer method)
Not working yet 
```shell
#--go-grpc_out=require_unimplemented_servers=false
#   protoc --go_out=. **--go-grpc_opt=require_unimplemented_servers=false** --go-grpc_out=. proto/*.proto 
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/currency.proto

protoc -I ./proto \
   --go_out ./proto --go_opt paths=source_relative \
   --go-grpc_out ./proto --go-grpc_opt paths=source_relative require_unimplemented_servers=false \
   ./proto/currency.proto
```
# Test GRPC
Not possible to use curl, but we can make UT
## Install grpcurl
Install and run following changes
```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
# Add following line in main.go
#	reflection.Register(gs)
```
Note: Run commands in BASH, json format is different for powershell
After that it is possible to run following lines:
```shell
 ~  grpcurl --plaintext localhost:9092 list
Failed to list services: server does not support the reflection API

 ~  grpcurl --plaintext localhost:9092 list Currency
Failed to list methods for service "Currency": server does not support the reflection API

 ~  grpcurl --plaintext localhost:9092 list Currency
Currency.GetRate

 ~  grpcurl --plaintext localhost:9092 list
Currency
grpc.reflection.v1alpha.ServerReflection

 ~  grpcurl --plaintext localhost:9092 list Currency
Currency.GetRate

 ~  grpcurl --plaintext localhost:9092 describe Currency.GetRate
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );

 ~  grpcurl --plaintext localhost:9092 describe .RateRequest
RateRequest is a message:
message RateRequest {
  string Base = 1;
  string Destination = 2;
}

 ~  grpcurl --plaintext localhost:9092 describe .RateResponse
RateResponse is a message:
message RateResponse {
  float Rate = 1;
}

 ~  grpcurl --plaintext -d '{"base":"GBP","destination":"US"}' localhost:9092 Currency.GetRate
{
  "Rate": 0.5
}
```
