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
	@cd ./rpc_gen && cwgo client --type RPC --service order --module rpc_gen  -I ../idl  --idl ../idl/order.thrift
	@cd ./app/admin && cwgo server --type RPC --module order --service order --pass "-use rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.thrift



.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module frontend --idl ../../idl/frontend/member.thrift

.PHONY: gen-ent
gen-ent:
	go generate ./pkg/db/ent