package main

import (
	"context"
	"flag"
	"fmt"
	"myraft/asyclog"
	"myraft/shared"
	"myraft/timer"
	"net"
	"os"

	consistency "myraft/consistency"

	"google.golang.org/grpc"
)
var (
	port = flag.Int("port", 11451, "The port to listen on for HTTP requests.")
	role = flag.String("role", "follower", "The role of this server")
	configurePath = flag.String("path", "", "The configure file to read")
)

type consistencyServer struct{
	consistency.UnimplementedConsistencyServer
	Nodenums int32
	nodes []*consistency.Node
}

func (s* consistencyServer) sendConfigure(ctx context.Context, request* consistency.ConfigureRequest) (response* consistency.ConfigureResponse, err error){
	s.nodes = request.Nodes
	s.Nodenums = request.Nodenum
	return &consistency.ConfigureResponse{Success: true}, nil
}

func serverConfigureInit(confname string) error{
	var err error
	if shared.Getstatus() == shared.Leader {
		if err := shared.InitLeader(); err != nil {
			return err
		}
		config := shared.GetInstance()
		if config.Nodes, err = shared.ReadConfig(confname); err != nil {
			return err
		}
		if err := shared.InitStateTable(); err != nil {
			return err
		}
		var configureRequest consistency.ConfigureRequest
		configureRequest.Nodenum = int32(config.Nodenums)
		for i := 0; i < config.Nodenums; i++ {
			configureRequest.Nodes[i].Ip = config.Nodes[i].Host
			configureRequest.Nodes[i].Port = int32(config.Nodes[i].Port)
		}
		for i := 0; i < config.Nodenums; i++ {
			var conn* grpc.ClientConn
			if conn, err = shared.InitStub(config.Nodes[i]); err != nil {
				continue
			}
			client := consistency.NewConsistencyClient(conn)
			defer conn.Close()
			client.Sendconfigure(context.Background(), &consistency.ConfigureRequest{})
			stateTable := *shared.GetstatesTable()
			stateTable[config.Nodes[i]] = *shared.NewNodestate(true)
		}
	} else {
		lis, _ := net.Listen("tcp", ":11451")
		s := grpc.NewServer()
		consistency.RegisterConsistencyServer(s, &consistencyServer{})
		s.Serve(lis)
		defer s.Stop()
	}
	return nil
}
func main() {
	flag.Parse()
	if *role == "Leader" {
		shared.Setstatus(shared.Leader)
	} else {
		shared.Setstatus(shared.Follower)
	}
	thisIp = shared.GetHostIpv4()
	log := asynclog.NewAsyncLogger(1024, 10)
	inittimer := mytimer.NewTimerWrapper(1000, func() {
		log.Log("init time out, process abort!")
		os.Exit(1)
	})
	if err := serverConfigureInit(*configurePath); err != nil {
		log.Log(fmt.Sprintf("ConfigureInit error: %v", err))
		os.Exit(1)
	}
	inittimer.Stop()
	for ;; {
		log.Log("server start loop")
		if (shared.Getstatus() == shared.Leader){
			beattimer := mytimer.NewTickerWrapper()
		}
	}
}