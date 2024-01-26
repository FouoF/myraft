package main

import (
	"context"
	"fmt"
	"myraft/asyclog"
	"myraft/shared"
	"myraft/timer"
	"os"

	consistency "myraft/consistency"
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

func clinetConfigureInit() error{
	for ;;{
		if shared.GetThisNodeConfigureOK() {
			break
		}
	}
	return nil
}
func main() {
	log := asynclog.NewAsyncLogger(1024, 10)
	inittimer := mytimer.NewTimerWrapper(1000, func() {
		log.Log("init time out, process abort!")
		os.Exit(1)
	})
	if err := clinetConfigureInit(); err != nil {
		log.Log(fmt.Sprintf("ConfigureInit error: %v", err))
		os.Exit(1)
	}
	inittimer.Stop()
	for ;; {
		log.Log("start loop")
	}
}