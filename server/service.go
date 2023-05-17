package monitor

import (
	"fmt"
	"log"
	"net"

	monitor "../api-monitor/monitor/v1"
	"google.golang.org/grpc"
)

type Service struct {
}

func Start() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	monitor.RegisterMonitoringServer(server, &Service{})
	fmt.Println("server started on port", ":50051")
	return server.Serve(lis)

}
