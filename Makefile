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
	@cd ./rpc_gen && cwgo client --type RPC --server_name system --module rpc_gen  -I ../idl  --idl ../idl/system/menu.thrift
	@cd ./rpc_gen && cwgo client --type RPC --server_name system --module rpc_gen  -I ../idl  --idl ../idl/system/auth.thrift
	@cd ./rpc_gen && cwgo client --type RPC --server_name system --module rpc_gen  -I ../idl  --idl ../idl/system/dictionary.thrift
	@cd ./rpc_gen && cwgo client --type RPC --server_name system --module rpc_gen  -I ../idl  --idl ../idl/system/sys.thrift

	@cd ./rpc_gen && cwgo client --type RPC --service user --module rpc_gen  -I ../idl  --idl ../idl/user/user.thrift
	@cd ./app/user/ && cwgo client --type RPC --service user --module rpc_gen  -I ../idl  --idl ../idl/user/token.thrift

	@cd ./app/schedule/ && cwgo client --type RPC --service schedule --module rpc_gen  -I ../idl  --idl ../idl/schedule/schedule.thrift

	@cd ./app/company/ && cwgo client --type RPC --service company --module rpc_gen  -I ../idl  --idl ../idl/company/contract.thrift
	@cd ./app/company/ && cwgo client --type RPC --service company --module rpc_gen  -I ../idl  --idl ../idl/company/entry.thrift
	@cd ./app/company/ && cwgo client --type RPC --service company --module rpc_gen  -I ../idl  --idl ../idl/company/venue.thrift


.PHONY: gen-server
gen-server:
	@cd ./app/system/ && cwgo server --type RPC --module system --server_name system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/menu.thrift
	@cd ./app/system/ && cwgo server --type RPC --module system --server_name system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/auth.thrift
	@cd ./app/system/ && cwgo server --type RPC --module system --server_name system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/dictionary.thrift
	@cd ./app/system/ && cwgo server --type RPC --module system --server_name system --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/system/sys.thrift



	@cd ./app/user/ && cwgo server --type RPC --module user --server_name user --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user/user.thrift
	@cd ./app/user/ && cwgo server --type RPC --module user --server_name user --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user/token.thrift

	@cd ./app/schedule/ && cwgo server --type RPC --module schedule --server_name schedule --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/schedule/schedule.thrift

	@cd ./app/company/ && cwgo server --type RPC --module company --server_name company --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/company/contract.thrift
	@cd ./app/company/ && cwgo server --type RPC --module company --server_name company --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/company/entry.thrift
	@cd ./app/company/ && cwgo server --type RPC --module company --server_name company --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/company/venue.thrift


.PHONY: gen-http
gen-http:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service facade --module facade --idl ../../idl/facade/member.thrift

.PHONY: gen-ent
gen-ent:
	@cd app/company/biz/dal/mysql && go generate ./ent --feature sql/schemaconfig
	@cd app/system/biz/dal/mysql && go generate ./ent --feature sql/schemaconfig
	@cd app/user/biz/dal/mysql && go generate ./ent --feature sql/schemaconfig
	@cd app/schedule/biz/dal/mysql && go generate ./ent --feature sql/schemaconfig



