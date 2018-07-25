#!/usr/bin/env bash
go build -ldflags "-X main.BuildTime=`date '+%Y-%m-%d_%I:%M:%S'`"