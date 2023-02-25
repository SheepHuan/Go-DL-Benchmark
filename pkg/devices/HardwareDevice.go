package devices

import (
	"go-dl-benchmark/pkg/terminal"
)

const (
	Android = "android"
	Ios     = "ios"
	Mac     = "mac"
	Linux   = "linux"
	Windows = "windows"
	NoneOS  = ""

	X86_64 = "x86_64"
	Arm64  = "arm64"

	Cpu = "cpu"
	Gpu = "gpu"
	Npu = "npu"
)

type HardwareDescription struct {
	DeviceName     string `json:"DeviceName"`
	OSType         string `json:"OSType"`
	Architecture   string `json:"Architecture"`
	ComputableChip string `json:"ComputableChip"`
	// 一级设备标识
	Ip   string
	Port int

	// 二级设备标识
	andoridIpAddress string
	androidSerial    string

	// USB 设备标识
	usbIpAddress string
}

type HardwareDevice struct {
	_terminal   *terminal.RemoteTerminal
	_isAlive    bool
	Description HardwareDescription `json:"Description"`
}

func (t *HardwareDevice) IsAlive() bool {
	if t._isAlive != true {
		return false
	}
	return true
}

func (t *HardwareDevice) Connect() error {
	// 检查当前设备类型是否属于一级设备!
	if t.Description.OSType == Windows {
		t._terminal = &terminal.RemoteTerminal{}
		err := t._terminal.Init(t.Description.Ip, t.Description.Port)
		if err != nil {
			t._isAlive = false
			return err
		}
		t._isAlive = true
	} else {
		// 如果不是一级设备应当先和一级设备沟通
	}
	return nil
}

func (t *HardwareDevice) Disconnect() error {
	if t.Description.OSType != Windows {
		err := t._terminal.Close()
		if err != nil {
			t._isAlive = true
			return err
		}
		t._isAlive = false
	} else {
		// 如果不是一级设备应当先和一级设备沟通
	}
	return nil
}

func (t *HardwareDevice) ExecuteCommand(cmd string) (string, string, error) {
	if t.Description.OSType == Windows {
		stdout, stderr, err := t._terminal.Execute(cmd)
		return stdout, stderr, err
	} else {
		// 如果不是一级设备应当先和一级设备沟通
		return "", "", nil
	}

}
