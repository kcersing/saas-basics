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
	@cd ./app/admin && go hz update -idl idl_gen/admin/schedule.thrift -model_dir idl_gen/model/  --unset_omitempty

.PHONY: gen-ent
gen-ent:
	go generate ./pkg/db/ent