#!/bin/sh

#ent 代码生成
# /biz/del/db/ent/schema/
go generate ./pkg/db/ent --feature sql/versioned-migration

go generate ./pkg/db/ent --feature sql/versioned-migration ./schema