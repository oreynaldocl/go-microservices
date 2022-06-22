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
