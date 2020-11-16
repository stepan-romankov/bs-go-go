#!/usr/bin/env bash

protoc -I=. --go_out=./ --go-grpc_out=require_unimplemented_servers=false:. ./service.proto