package g_physical_devices

import (
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	"github.com/sheephuan/go-dl-benchmark/pkg/terminal"
)

type PhysicalDeviceClient struct {
	_terminal         *terminal.RemoteTerminal
	_isAlive          bool
	DeviceDescription *protos.PhysicalDeviceDescription
}

func (s *PhysicalDeviceClient) IsAlive() bool {
	return s._isAlive
}

func (s *PhysicalDeviceClient) Connect() error {
	if s.DeviceDescription.OSType == protos.DeviceOSType_windows {
		s._terminal = &terminal.RemoteTerminal{}
		pcAddr := s.DeviceDescription.GetPcAddr()
		if err := s._terminal.Init(pcAddr.DeviceIp, int(pcAddr.DevicePort)); err != nil {
			s._isAlive = false
			return err
		} else {
			s._isAlive = true
			return nil
		}
	} else if s.DeviceDescription.OSType == protos.DeviceOSType_linux {

	} else if s.DeviceDescription.OSType == protos.DeviceOSType_android {

	}
	return nil
}

func (s *PhysicalDeviceClient) Disconnect() error {
	if s.DeviceDescription.OSType == protos.DeviceOSType_windows {
		err := s._terminal.Close()
		if err != nil {
			s._isAlive = true
			return err
		}
		s._isAlive = false
	} else if s.DeviceDescription.OSType == protos.DeviceOSType_linux {

	} else if s.DeviceDescription.OSType == protos.DeviceOSType_android {

	}
	return nil
}

func (s *PhysicalDeviceClient) ExecuteCommand(cmd string) (string, string, error) {
	if s.DeviceDescription.OSType == protos.DeviceOSType_windows {
		stdout, stderr, err := s._terminal.Execute(cmd)
		return stdout, stderr, err
	} else if s.DeviceDescription.OSType == protos.DeviceOSType_linux {

	} else if s.DeviceDescription.OSType == protos.DeviceOSType_android {

	}
	return "", "", nil
}
