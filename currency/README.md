## Create Proto
Run following line
```shell
#protos:
# YOUTUBE EXAMPLE https://www.youtube.com/watch?v=pMgty_RYIOc&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_&index=13&ab_channel=NicJackson
protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency
# OFFICIAL https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/currency
```
The difference could be by the version difference
