package g_physical_devices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	"github.com/sheephuan/go-dl-benchmark/pkg/utils"
	"io"
	"net/http"
)

var BASE_URL = "http://127.0.0.1:8888/v1"

type PhysicalDevicesForm struct {
	RegisterDevicesList []string `json:"registerDevicesList"`
	DeleteDevicesList   []string `json:"deleteDevicesList"`
}

func RegisterDevicesSelf(descriptions []*protos.PhysicalDeviceDescription) error {
	apiUrl := BASE_URL + "/PhysicalDeviceComm/registerPhysicalDevices"

	var stringList []string
	stringList = make([]string, len(descriptions))
	for i, item := range descriptions {
		d, _ := proto.Marshal(item)
		s := utils.Pb2Base64(d)
		stringList[i] = s
	}
	registerDevicesForm := PhysicalDevicesForm{RegisterDevicesList: stringList, DeleteDevicesList: []string{}}
	b, err := json.Marshal(registerDevicesForm)
	if err == nil {
		payload := bytes.NewBuffer(b)
		resp, err := http.Post(apiUrl, "application/json;charset=utf-8", payload)
		if err == nil {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				fmt.Println(string(body))
			}

		}
	}

	return nil
}
