package bosh

import (
	"encoding/json"
	cfg "github.com/aemengo/bosh-runc-cpi/config"
)

type AgentSettings struct {
	ID        string                 `json:"agent_id"`
	VM        VM                     `json:"vm"`
	Mbus      string                 `json:"mbus"`
	NTP       []string               `json:"ntp"`
	Blobstore Blobstore              `json:"blobstore"`
	Networks  map[string]interface{} `json:"networks"`
	Disks     map[string]interface{} `json:"disks"`
	Env       Env                    `json:"env"`
}

type VM struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Blobstore struct {
	Provider string                 `json:"provider"`
	Options  map[string]interface{} `json:"options"`
}

type Env struct {
	Bosh Bosh `json:"bosh"`
}

type Bosh struct {
	Password string                 `json:"password"`
	Group    string                 `json:"group"`
	Groups   []interface{}          `json:"groups"`
	Mbus     map[string]interface{} `json:"mbus"`
}

func ConvertAgentSettings(env []interface{}, config cfg.Config) []byte {
	var settings AgentSettings

	settings.ID = str(fetch(env, 0))
	settings.Mbus = config.Agent.Mbus
	settings.NTP = config.Agent.NTP
	settings.Blobstore = Blobstore{
		Provider: config.Agent.Blobstore.Provider,
		Options:  config.Agent.Blobstore.Options,
	}

	networks := mapping(fetch(env, 3))
	for _, network := range networks {
		network.(map[string]interface{})["mac"] = ""
		network.(map[string]interface{})["preconfigured"] = true
	}

	settings.Networks = networks
	settings.Disks = map[string]interface{}{
		"system":     "",
		"ephemeral":  nil,
		"persistent": nil,
	}
	settings.Env = Env{
		Bosh: Bosh{
			Password: str(fetch(env, 5, "bosh", "password")),
			Group:    str(fetch(env, 5, "bosh", "group")),
			Groups:   array(fetch(env, 5, "bosh", "groups")),
			Mbus:     mapping(fetch(env, 5, "bosh", "mbus")),
		},
	}

	contents, _ := json.Marshal(settings)
	return contents
}

func fetch(env []interface{}, args ...interface{}) interface{} {
	var val interface{}

	for _, arg := range args {
		sarg, ok := arg.(string)
		if ok {
			if val != nil {
				mapped, _ := val.(map[string]interface{})
				val = mapped[sarg]
			}

			continue
		}

		iarg, ok := arg.(int)
		if ok {
			if val == nil {
				val = env[iarg]
			} else {
				arred, _ := val.([]interface{})
				val = arred[iarg]
			}

			continue
		}
	}

	return val
}

func mapping(input interface{}) map[string]interface{} {
	val, _ := input.(map[string]interface{})
	return val
}

func array(input interface{}) []interface{} {
	val, _ := input.([]interface{})
	return val
}

func str(input interface{}) string {
	val, _ := input.(string)
	return val
}
