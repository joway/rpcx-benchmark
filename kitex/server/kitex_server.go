package main

import (
	"flag"
	"github.com/rpcxio/rpcx-benchmark/kitex/pb/hello"
	"net"
	"runtime"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/server"
	pb "github.com/rpcxio/rpcx-benchmark/proto"
	"github.com/smallnest/rpcx/log"
	"golang.org/x/net/context"
)

var (
	host  = flag.String("s", "127.0.0.1:8972", "listened ip and port")
	delay = flag.Duration("delay", 0, "delay to mock business processing")
)

type Hello struct{}

func (t *Hello) Say(ctx context.Context, args *pb.BenchmarkMessage) (reply *pb.BenchmarkMessage, err error) {
	s := "OK"
	var i int32 = 100
	args.Field1 = s
	args.Field2 = i
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return args, nil
}

func main() {
	flag.Parse()

	ipAddr, portNum, _ := net.SplitHostPort(*host)
	ip := net.ParseIP(ipAddr)
	port, _ := strconv.Atoi(portNum)
	svr := hello.NewServer(new(Hello), server.WithServiceAddr(&net.TCPAddr{IP: ip, Port: port}))
	if err := svr.Run(); err != nil {
		log.Fatalf("server stopped with error:", err)
	} else {
		log.Info("server stopped")
	}
}
