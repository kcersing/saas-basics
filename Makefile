# start the environment of FreeCar

.PHONY: start
start:
	docker-compose up -d --build 

# stop the environment of FreeCar

.PHONY: stop
stop:
	docker-compose down

.PHONY: gen-rpc
gen-rpc:
	cwgo client --type RPC --server_name system --module rpc_gen  -I ../idl  --idl ../idl/system/menu.thrift
	cwgo client --type RPC --server_name system --module rpc_gen  -I ../idl  --idl ../idl/system/auth.thrift

.PHONY: gen-server
gen-server:
	@cd ./rpc_gen && cwgo server --type RPC --module system --server_name system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/menu.thrift
	@cd ./rpc_gen && cwgo server --type RPC --module system --server_name system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/auth.thrift

.PHONY: gen-http
gen-http:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service facade --module facade --idl ../../idl/facade/member.thrift

.PHONY: gen-ent
gen-ent:
	@cd app/system/biz/dal/mysql && go generate ./ent
