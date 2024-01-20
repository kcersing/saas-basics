#!/bin/sh

#if [ "$1" = "init" ]; then
#    hz new -mod saas-basics
#    hz update -api api/init.proto
#fi

#hz 代码生成
# /api  saas-basics/app/api 下
hz new -idl ./../../idl/http/user.thrift -mod app/api
#kitex 代码生成
#  saas-basics/app/user 下
kitex -module saas -service user -gen-path ../../kitex_gen  ./../../idl/rpc/user.thrift

kitex -module saas -service order -gen-path ../../kitex_gen -record ../../idl/rpc/order.thrift

kitex -module saas -service item -gen-path ../../kitex_gen -record ../../idl/rpc/item.thrift