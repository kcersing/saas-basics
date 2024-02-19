#!/bin/sh

#if [ "$1" = "init" ]; then
#    hz new -mod saas-basics
#    hz update -api api/init.proto
#fi

#hz 代码生成
# /api  saas-basics/app/api 下
#hz new -idl ./../idl/http/user.thrift





#本项目使用 `hz` 生成代码. `hz` 详细使用说明参考 [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/).
#- hz install.
#```bash
#go install github.com/cloudwego/hertz/cmd/hz@latest
#```
#- hz new: 新建一个 Hertz 项目
#```bash
#hz new -I api -idl api/admin/admin.proto -model_dir api/model -module formulago --unset_omitempty
#```
#- hz update: 当你的IDL文件更新，使用该指令进行项目代码更新
#- api.proto 与 base.proto是不需要更新与生成的，因为它们是由导入它们的proto文件生成的
#```bash
hz update -I api -idl api/admin/admin.proto -model_dir api/model --unset_omitempty


hz update -idl ./../idl/base/errno.thrift
hz update -idl ./../idl/admin/user.thrift
hz update -idl ./../idl/admin/admin.thrift
hz update -idl ./../idl/admin/dictionary.thrift
hz update -idl ./../idl/admin/logs.thrift
hz update -idl ./../idl/admin/menu.thrift
hz update -idl ./../idl/admin/role.thrift
hz update -idl ./../idl/admin/token.thrift


hz update -I api -idl ./../idl/admin/token.thrift -model_dir ./../app/biz/model/ --unset_omitempty
