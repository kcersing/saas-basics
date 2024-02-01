#!/bin/sh

#if [ "$1" = "init" ]; then
#    hz new -mod saas-basics
#    hz update -api api/init.proto
#fi

#hz 代码生成
# /api  saas-basics/app/api 下
hz new -idl ./../idl/http/user.thrift


hz update -idl ./../idl/base/errno.thrift
hz update -idl ./../idl/admin/user.thrift
hz update -idl ./../idl/admin/admin.thrift
hz update -idl ./../idl/admin/dictionary.thrift
hz update -idl ./../idl/admin/logs.thrift
hz update -idl ./../idl/admin/menu.thrift
hz update -idl ./../idl/admin/role.thrift
hz update -idl ./../idl/admin/token.thrift

