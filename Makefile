# start the environment of FreeCar

.PHONY: start
start:
	docker-compose up -d --build 

# stop the environment of FreeCar

.PHONY: stop
stop:
	docker-compose down

# run the user
.PHONY: admin
admin:
	go run ./app/admin

# run the api
.PHONY: api
api:
	go run ./app/api

.PHONY: gen-user
gen-user:
	@cd ./rpc_gen && cwgo client --type RPC --service order --module rpc_gen  -I ../idl  --idl ../idl/order/order.thrift
	@cd ./app/admin && cwgo server --type RPC --module order --service order --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order/order.thrift

	@cd ./rpc_gen && cwgo client --type RPC --service member --module rpc_gen  -I ../idl  --idl ../idl/member/member.thrift
	@cd ./app/admin && cwgo server --type RPC --module member --service member --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/member/member.thrift


	@cd ./rpc_gen && cwgo client --type RPC --service system --module rpc_gen  -I ../idl  --idl ../idl/system/dictionary.thrift
	@cd ./app/system && cwgo server --type RPC --module system --service system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/dictionary.thrift

	@cd ./rpc_gen && cwgo client --type RPC --service system --module rpc_gen  -I ../idl  --idl ../idl/system/logs.thrift
	@cd ./app/system && cwgo server --type RPC --module system --service system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/logs.thrift
	@cd ./rpc_gen && cwgo client --type RPC --service system --module rpc_gen  -I ../idl  --idl ../idl/system/menu.thrift
	@cd ./app/system && cwgo server --type RPC --module system --service system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/menu.thrift
	@cd ./rpc_gen && cwgo client --type RPC --service system --module rpc_gen  -I ../idl  --idl ../idl/system/pub.thrift
	@cd ./app/system && cwgo server --type RPC --module system --service system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/pub.thrift
	@cd ./rpc_gen && cwgo client --type RPC --service system --module rpc_gen  -I ../idl  --idl ../idl/system/role.thrift
	@cd ./app/system && cwgo server --type RPC --module system --service system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/role.thrift
	@cd ./rpc_gen && cwgo client --type RPC --service system --module rpc_gen  -I ../idl  --idl ../idl/system/sys.thrift
	@cd ./app/system && cwgo server --type RPC --module system --service system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/sys.thrift
	@cd ./rpc_gen && cwgo client --type RPC --service system --module rpc_gen  -I ../idl  --idl ../idl/system/token.thrift
	@cd ./app/system && cwgo server --type RPC --module system --service system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/token.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service facade --module facade --idl ../../idl/facade/member.thrift

.PHONY: gen-ent
gen-ent:
	go generate ./pkg/db/ent