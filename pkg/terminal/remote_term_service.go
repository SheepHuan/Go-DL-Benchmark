package terminal

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

type RemoteTerminalService struct {
	shellIdsMap map[uint64]*Terminal
}

func (s *RemoteTerminalService) mustEmbedUnimplementedRemoteTerminalServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *RemoteTerminalService) NewTerminal(ctx context.Context, request *RemoteTerminalRequest) (*RemoteTerminalResponse, error) {
	handle := uint64(time.Now().Unix())
	terminal, err := NewPowerShell()
	resp := new(RemoteTerminalResponse)
	if err != nil {
		resp.Err = err.Error()
		return resp, err
	}
	//log.Info(fmt.Sprintf("We new a terminal with %d", handle))
	s.shellIdsMap[handle] = terminal
	resp.ShellId = handle
	return resp, nil
}

func (s *RemoteTerminalService) CloseTerminal(ctx context.Context, request *RemoteTerminalRequest) (*RemoteTerminalResponse, error) {
	resp := new(RemoteTerminalResponse)
	handle := request.ShellId
	if terminal, ok := s.shellIdsMap[handle]; ok {
		terminal.Exit()
	} else {
		resp.Err = fmt.Sprint("shell:%d is not exsiting!", handle)
	}
	return resp, nil
}

func (s *RemoteTerminalService) ExecCommand(ctx context.Context, request *RemoteTerminalRequest) (*RemoteTerminalResponse, error) {
	resp := new(RemoteTerminalResponse)
	handle := request.ShellId
	//fmt.Println(fmt.Sprintf("shell:%d should exec %s", handle, request.Command))
	if terminal, ok := s.shellIdsMap[handle]; ok {
		command := request.Command
		sout, serr, err := terminal.Execute(command)
		resp.Stdout = sout
		resp.Stderr = serr
		if err != nil {
			resp.Err = err.Error()
		}
	} else {
		resp.Err = fmt.Sprint("shell:%d is not exsiting!", handle)
	}
	return resp, nil
}

func LaunchRemoteTerminalService(ip string, port int) {
	l := newNetListener(ip, port)
	// 此处设置最佳发送文件大小512M
	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 512),
		grpc.MaxSendMsgSize(1024 * 1024 * 512),
	}
	s := grpc.NewServer(options...)
	rs := RemoteTerminalService{shellIdsMap: make(map[uint64]*Terminal)}
	RegisterRemoteTerminalServiceServer(s, &rs)
	log.Info("Start Remote Terminal Service!")
	err := s.Serve(l)
	if err != nil {
		log.Error(err)
		return
	}

}

func newNetListener(ip string, port int) net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Fatal(err)
	}
	return l
}
