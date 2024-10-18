#!/bin/sh

#if [ "$1" = "init" ]; then
#    hz new -mod saas-basics
#    hz update -api api/init.proto
#fi

#hz 代码生成
# /api  saas-basics/app/api 下
#hz new -idl_gen ./../idl_gen/http/user.thrift





#本项目使用 `hz` 生成代码. `hz` 详细使用说明参考 [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/).
#- hz install.
#```bash
#go install github.com/cloudwego/hertz/cmd/hz@latest
#```
#- hz new: 新建一个 Hertz 项目
#```bash
#hz new -I api -idl_gen ./../idl_gen/admin/token.thrift -model_dir ./../app/biz/model/ --unset_omitempty
#```
#- hz update: 当你的IDL文件更新，使用该指令进行项目代码更新
#- api.proto 与 base.proto是不需要更新与生成的，因为它们是由导入它们的proto文件生成的
#```bash

#hz new -I api -idl_gen ./../idl_gen/admin/token.thrift -model_dir ./../app/biz/model/ --unset_omitempty

hz update -idl idl_gen/base/errno.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/user.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/admin.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/dictionary.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/logs.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/menu.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/role.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/token.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/sys.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/product.thrift -model_dir idl_gen/model/  --unset_omitempty
hz update -idl idl_gen/admin/schedule.thrift -model_dir idl_gen/model/  --unset_omitempty
