# start the environment of FreeCar

.PHONY: start
start:
	docker-compose up -d --build 

# stop the environment of FreeCar

.PHONY: stop
stop:
	docker-compose down

# run the user
.PHONY: run
admin:
	go run .

.PHONY: gen-gen
gen-user:
	hz update -idl idl_gen/idl/admin.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/auth.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/contest.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/dictionary.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/menu.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/pub.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/user.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/token.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/venue.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/member.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/contract.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/entry.thrift -model_dir idl_gen/model/  --unset_omitempty
.PHONY: gen-ent
gen-ent:
	go generate ./pkg/db/ent