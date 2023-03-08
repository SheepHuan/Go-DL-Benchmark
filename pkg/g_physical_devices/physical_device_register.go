package g_physical_devices

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	"io"
	"net/http"
	"strings"
)

var BASE_URL = "http://127.0.0.1:8888/v1"

type PhysicalDevicesForm struct {
	registerDevicesList []string `json:"registerDevicesList"`
	deleteDevicesList   []string `json:"deleteDevicesList"`
}

func RegisterDevicesSelf(descriptions []*protos.PhysicalDeviceDescription) error {
	apiUrl := BASE_URL + "/PhysicalDeviceComm/registerPhysicalDevices"
	marshaler := jsonpb.Marshaler{}
	var stringList []string
	stringList = make([]string, len(descriptions))
	for i, item := range descriptions {
		s, _ := marshaler.MarshalToString(item)
		stringList[i] = s
	}
	registerDevicesForm := PhysicalDevicesForm{registerDevicesList: stringList}
	b, err := json.Marshal(registerDevicesForm)
	if err == nil {
		payload := strings.NewReader(string(b))
		resp, err := http.Post(apiUrl, "application/json", payload)
		if err == nil {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				fmt.Println(string(body))
			}

		}
	}

	return nil
}
