package shared

import (
	"bufio"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc"
)

var status Status = Unknow //当前节点的角色[Leader, Follower, Unknow]
var config conf //配置信息
var state State //当前节点的状态
var logs []LogEntry //日志
var thisNodeConfigureOK bool = false

func GetThisNodeConfigureOK() bool {
	return thisNodeConfigureOK
}

type Status int

const (
	Running Status = iota
	Leader
	Follower
	Unknow
)

func Getstatus() Status {
	return status
}

func Setstatus(pstatus Status) {
	status = pstatus
}

type node struct{
	Host string
	Port int
}

func Newnode() *node{
	return &node{
		Host: "127.0.0.1",
		Port: 0,
	}
}

type conf struct{
	Nodenums int
	Nodes []node
}

type Nodestate struct{
	conn *grpc.ClientConn
	configed bool
}

func NewNodestate(confed bool) *Nodestate{
	return &Nodestate{
		configed: confed,
	}
}

type StatesTable map[node]Nodestate

var statesTable StatesTable

func GetstatesTable() *StatesTable {
	return &statesTable
}

var instance *conf
var once sync.Once

type LogEntry struct{
	Term int
	Index int
	Command string
}

type State struct{
	term int
	commitIndex int
	lastApplied int
}

func GetInstance() *conf {
    once.Do(func(){
        instance = &conf{
			Nodenums: 0,
			Nodes: []node{},
		}
    })
    return instance
}

func ReadConfig(filename string) ([]node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var configs []node
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, err
		}
		i, _:= strconv.Atoi(strings.TrimSpace(parts[0]))
		config := node{
			Host:   strings.TrimSpace(parts[0]),
			Port: 	i,
		}
		configs = append(configs, config)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return configs, nil
}
func InitLeader() (error) {
	state.term = 1
	state.commitIndex = 0
	state.lastApplied = 0
	logs = append(logs, LogEntry{1, 1, "init"})
	return nil
}
func InitStateTable() (error) {
	for i := 0; i < config.Nodenums; i++ {
		statesTable[config.Nodes[i]] = Nodestate{nil, false}
	}
	return nil
}
func InitStub(addr node) (*grpc.ClientConn, error){
	if (statesTable[addr].conn != nil) {
		return statesTable[addr].conn, nil
	}
	serverAddr := addr.Host + ":" + strconv.Itoa(addr.Port)
	var opts []grpc.DialOption
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	statesTable[addr] = Nodestate{conn, true}
	return conn, nil
}

func GetHostIpv4() (string) {
	hostname, err := os.Hostname()
	if err != nil {
		return " "
	}
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return " "
	}
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return " "
}