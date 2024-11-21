#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/schedule"
exec "$CURDIR/bin/schedule"