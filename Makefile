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
#本项目使用 `hz` 生成代码. `hz` 详细使用说明参考 [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/).
#- hz install.
#```bash
#go install github.com/cloudwego/hertz/cmd/hz@latest
#```
#- hz new: 新建一个 Hertz 项目
#```bash
#hz new -I api -idl_gen ./../idl_gen/idl/token.thrift -model_dir ./../app/biz/model/ --unset_omitempty
#```
#- hz update: 当你的IDL文件更新，使用该指令进行项目代码更新
#- api.proto 与 base.proto是不需要更新与生成的，因为它们是由导入它们的proto文件生成的
#```bash
gen-user:
	hz update -idl idl_gen/idl/admin.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/auth.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/contest.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/contract.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/dictionary.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/entry.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/member.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/menu.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/pub.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/token.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/venue.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/user.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/payment.thrift -model_dir idl_gen/model/  --unset_omitempty

	swag init

.PHONY: gen-ent
gen-ent:
	go generate ./biz/dal/db/ent
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/ent/schema

#cd /xingfufit/go/saas-basics/
#lsof -i :9039
#export PATH=$PATH:/usr/local/go/bin
#alias air='$(go env GOPATH)/bin/air'
#git pull
#go run .


