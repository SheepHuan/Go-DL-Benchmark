package terminal

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type RemoteTerminal struct {
	_handle  uint64
	_dstIP   string
	_dstPort int
}

func (t *RemoteTerminal) Init(serverIP string, serverPort int) error {
	var err error
	t._handle, err = NewRemotePowershell(serverIP, serverPort)
	if err != nil {
		t._handle = 0
		log.Error(err)
		return err
	}
	log.Info(fmt.Sprintf("Connect to terminal service[%s:%d].", serverIP, serverPort))
	t._dstIP = serverIP
	t._dstPort = serverPort
	return nil
}

func (t *RemoteTerminal) Execute(cmd string) (string, string, error) {
	req := newRemoteTerminalRequest("", "", cmd, t._handle, 0)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", t._dstIP, t._dstPort), grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return "", "", err
	}
	defer conn.Close()
	c := NewRemoteTerminalServiceClient(conn)
	rep, err := c.ExecCommand(context.Background(), &req)
	if err != nil {
		log.Error(err.Error())
		return "", "", err
	}
	err = conn.Close()
	if err != nil {
		log.Error(err.Error())
		return "", "", err
	}
	return rep.Stdout, rep.Stderr, nil
}

func (t *RemoteTerminal) Close(cmd string) error {
	req := newRemoteTerminalRequest("", "", cmd, t._handle, 0)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", t._dstIP, t._dstPort), grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer conn.Close()
	c := NewRemoteTerminalServiceClient(conn)
	_, err = c.CloseTerminal(context.Background(), &req)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = conn.Close()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func NewRemotePowershell(serverIP string, serverPort int) (uint64, error) {
	req := newRemoteTerminalRequest("", "", "", 0, 0)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serverIP, serverPort), grpc.WithInsecure())
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	c := NewRemoteTerminalServiceClient(conn)
	res, err := c.NewTerminal(context.Background(), &req)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	err = conn.Close()
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return res.ShellId, nil
}

func newRemoteTerminalRequest(dstIP string, srcIP string, command string, shellId uint64, deviceId uint64) RemoteTerminalRequest {
	request := RemoteTerminalRequest{
		DstIP:    dstIP,
		SrcIP:    srcIP,
		ShellId:  shellId,
		Command:  command,
		DeviceId: deviceId,
	}
	return request
}
