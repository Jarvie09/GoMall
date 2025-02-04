#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/demo/demo_thrift"
exec "$CURDIR/bin/demo/demo_thrift"