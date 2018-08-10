package service

import (
	"encoding/json"
	"errors"
	"github.com/aemengo/bosh-runc-cpi/bosh"
)

func extractNetValues(agentSettings []byte) (string, string, string, error) {
	var values struct {
		Networks map[string]map[string]interface{} `json:"networks"`
	}

	json.Unmarshal(agentSettings, &values)

	for _, v := range values.Networks {
		ip, _ := v["ip"].(string)
		mask, _ := v["netmask"].(string)
		gateway, _ := v["gateway"].(string)
		return ip, mask, gateway, nil
	}

	return "", "", "", errors.New("unable to extract network values from provided agent settings")
}

func attachVMID(agentContents []byte, vmID string) []byte {
	var agentSettings map[string]interface{}
	json.Unmarshal(agentContents, &agentSettings)

	agentSettings["vm"] = bosh.VM{
		ID:   vmID,
		Name: vmID,
	}

	data, _ := json.Marshal(agentSettings)
	return data
}

func attachPersistentDisk(contents []byte, diskID string, guestDiskPath string) []byte {
	var agentSettings map[string]interface{}
	json.Unmarshal(contents, &agentSettings)

	agentSettings["disks"] = map[string]interface{}{
		"system":    "",
		"ephemeral": nil,
		"persistent": map[string]interface{}{
			diskID: map[string]string{
				"path": guestDiskPath,
			},
		},
	}

	data, _ := json.Marshal(agentSettings)
	return data
}

func detachPersistentDisk(contents []byte) []byte {
	var agentSettings map[string]interface{}
	json.Unmarshal(contents, &agentSettings)

	agentSettings["disks"] = map[string]interface{}{
		"system":     "",
		"ephemeral":  nil,
		"persistent": nil,
	}

	data, _ := json.Marshal(agentSettings)
	return data
}
