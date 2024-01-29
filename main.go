package main

import (
	"context"
	"math/rand"
	"flag"
	"fmt"
	"myraft/asyclog"
	"myraft/shared"
	"myraft/timer"
	"net"
	"os"
	"sync"
	"time"

	"myraft/consistency"
	"myraft/election"

	"google.golang.org/grpc"
)
var (
	port = flag.Int("port", 11451, "The port to listen on for HTTP requests.")
	role = flag.String("role", "follower", "The role of this server")
	configurePath = flag.String("path", "", "The configure file to read")
	stopping = false
	statechange chan bool
)

type electionServer struct{
	election.UnimplementedElectionServer
	timeout *mytimer.TimerWrapper
}

type consistencyServer struct{
	consistency.UnimplementedConsistencyServer
	Nodenums int32
	nodes []*consistency.Node
	mutex sync.Mutex
}

func (s* consistencyServer) Sendconfigure(ctx context.Context, request* consistency.ConfigureRequest) (response* consistency.ConfigureResponse, err error){
	s.mutex.Lock()
	s.nodes = request.Nodes
	s.Nodenums = request.Nodenum
	s.mutex.Unlock()
	return &consistency.ConfigureResponse{Success: true}, nil
}

func (s* electionServer) Sendheartbeat(ctx context.Context, request* election.HeartbeatRequest) (response* election.HeartbeatResponse, err error){
	request.Term = int32(shared.GetTerm())
	s.timeout.Reset(150 + time.Duration(rand.Intn(100)) * time.Millisecond)
	return &election.HeartbeatResponse{Success: true}, nil
}

func (s *electionServer) Vote(ctx context.Context, request *election.VoteRequest) (response *election.VoteResponse, err error){
	if request.IsPre {
		return &election.VoteResponse{VoteGranted: true}, nil
	}
	if (request.Term > int32(shared.GetTerm())){
		return &election.VoteResponse{VoteGranted: false, Term: int32(shared.GetTerm())}, nil
	}
	return &election.VoteResponse{VoteGranted: true}, nil
}

func newConsistencyServer() *consistencyServer{
	return &consistencyServer{}
}

func startHeartbeat() *mytimer.TickerWrapper{
	Table := shared.GetstatesTable()
	timer := mytimer.NewTickerWrapper(100 * time.Millisecond, func() {
		for _, v := range *Table{
			client := election.NewElectionClient(v.Conn)
			client.Sendheartbeat(context.Background(), &election.HeartbeatRequest{})
		}
	})
	return timer
}

func starttimeout(randnum int) *mytimer.TimerWrapper{
	Table := shared.GetstatesTable()
	var count = 1
	timer := mytimer.NewTimerWrapper((150 + time.Duration(randnum))* time.Millisecond, func() {
		for _, v := range *Table{
			client := election.NewElectionClient(v.Conn)
			response, _ := client.Vote(context.Background(), &election.VoteRequest{IsPre: true})
			if response.GetVoteGranted(){
				count ++
			}
		}
		if count > shared.GetInstance().Nodenums / 2 {
			shared.Setstatus(shared.Candidate)
			statechange <- true
		}
	})
	return timer
}

func startvote(){
	Table := shared.GetstatesTable()
	var count = 1
		for _, v := range *Table{
			client := election.NewElectionClient(v.Conn)
			response, _ := client.Vote(context.Background(), &election.VoteRequest{IsPre: false, Term: int32(shared.GetTerm())})
			if response.GetVoteGranted(){
				count ++
			}
		if count > shared.GetInstance().Nodenums / 2 {
			shared.Setstatus(shared.Leader)
			statechange <- true
		}else {
			shared.Setstatus(shared.Follower)
			statechange <- true
		}
	}
}

func startLogConsisitence(){
	
}

func ConfigureInit(confname string) error{
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
			if config.Nodes[i].Host == shared.GetHostIpv4() {
				continue
			}
			configureRequest.Nodes[i].Ip = config.Nodes[i].Host
			configureRequest.Nodes[i].Port = int32(config.Nodes[i].Port)
		}
		for i := 0; i < config.Nodenums; i++ {
			var conn* grpc.ClientConn
			if conn, err = shared.InitStub(config.Nodes[i]); err != nil {
				continue
			}
			client := consistency.NewConsistencyClient(conn)
			client.Sendconfigure(context.Background(), &configureRequest)
			stateTable := *shared.GetstatesTable()
			stateTable[config.Nodes[i]] = shared.Nodestate{Conn: nil, Configed: true}
		}
	} else {
		lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
		s := grpc.NewServer()
		consistencyserver := newConsistencyServer()
		consistency.RegisterConsistencyServer(s, consistencyserver)
		s.Serve(lis)
		for ;; {
			consistencyserver.mutex.Lock()
			if (consistencyserver.Nodenums != 0) {
				break
			}
			consistencyserver.mutex.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
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
//	thisIp := shared.GetHostIpv4()
	log := asynclog.NewAsyncLogger(1024, 10)
	inittimer := mytimer.NewTimerWrapper(1000 * time.Millisecond, func() {
		log.Log("init time out, process abort!")
		os.Exit(1)
	})
	if err := ConfigureInit(*configurePath); err != nil {
		log.Log(fmt.Sprintf("ConfigureInit error: %v", err))
		os.Exit(1)
	}
	inittimer.Stop()
	for ;; {
		if (stopping) {
			break
		}
		if (shared.Getstatus() == shared.Leader){
			heartbeattimer := startHeartbeat()
			startLogConsisitence()
			defer heartbeattimer.Stop()
		}else if (shared.Getstatus() == shared.Follower){
			lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
			s := grpc.NewServer()
			randnum := rand.Intn(100)
			electionserver := &electionServer{timeout:starttimeout(randnum)}
			election.RegisterElectionServer(s, electionserver)
			s.Serve(lis)
			for ;; {
				if (<- statechange) {
				s.Stop()
				break
			}
			}
		}else if (shared.Getstatus() == shared.Candidate){
			startvote()
		}
	}
}