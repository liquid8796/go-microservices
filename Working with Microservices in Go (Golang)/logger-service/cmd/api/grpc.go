package main

import "log-service/logs"

type LogServer struct {
	logs.UnimplementedLogServiceServer
}
