package bosh_test

import (
	. "github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bosh", func() {
	Describe("ConvertAgentSettings", func() {

		var config = cfg.Config{
			Agent: cfg.Agent{
				Mbus: "some-mbus-url",
				NTP: []string{"some-ntp"},
				Blobstore: cfg.Blobstore{
					Provider: "some-blobstore-provider",
					Options: map[string]interface{}{
						"kind": "some-blobstore-provider-option",
					},
				},
			},
		}

		var input = []interface{}{
			"some-agent-guid",
			"sc-5bd9284c-bb5d-4a4a-6564-1b5ecada5417",
			map[string]interface{}{},
			map[string]interface{}{
				"default": map[string]interface{}{
					"type": "manual",
					"ip": "10.244.0.133",
					"netmask": "255.255.240.0",
					"cloud_properties" : map[string]interface{}{
						"name": "random",
					},
					"default" : []string{"dns", "gateway" },
					"dns": []string{"8.8.8.8"},
					"gateway": "10.244.0.1",
				},
			},
			[]string{},
			map[string]interface{}{
				"bosh" : map[string]interface{}{
					"blobstores": []interface{}{
						map[string]interface{}{
							"options" : map[string]interface{}{
								"endpoint":    "http://192.168.50.6:25250",
								"password":    "some-blobstore-password",
								"tls" : map[string]interface{}{
									"cert":    map[string]interface{}{
										"ca": "some-bosh-cert-1",
									},
								},
								"user": "agent",
							},
							"provider": "dav",
						},
					},
					"mbus" : map[string]interface{}{
						"cert" : map[string]interface{}{
							"ca":    "some-mbus-ca",
							"certificate":    "some-mbus-cert",
							"private_key":    "some-mbus-private-key",
						},
					},
					"password": "some-password",
					"group": "some-group",
					"groups": []interface{}{
						"bosh-lite",
						"some-group-1",
						"some-group-2",
					},
				},
			},
		}

		It("converts the input to json agent settings", func() {
			settings := ConvertAgentSettings(input, config)
			Expect(settings).To(MatchJSON(`
{
  "agent_id": "some-agent-guid",
  "vm": {
    "id": "",
    "name": ""
  },
  "mbus": "some-mbus-url",
  "ntp": ["some-ntp"],
  "blobstore": {
    "provider": "some-blobstore-provider",
    "options": {
      "kind": "some-blobstore-provider-option"
    }
  },
  "networks": {
    "default": {
      "type": "manual",
      "ip": "10.244.0.133",
      "netmask": "255.255.240.0",
      "gateway": "10.244.0.1",
      "dns": [ "8.8.8.8" ],
      "default": [
        "dns",
        "gateway"
      ],
      "mac": "",
      "preconfigured": true,
      "cloud_properties": {
        "name": "random"
      }
    }
  },
  "disks": {
    "system": "",
    "ephemeral": null,
    "persistent": null
  },
  "env": {
    "bosh": {
      "group": "some-group",
      "groups": [
        "bosh-lite",
        "some-group-1",
        "some-group-2"
      ],
      "mbus": {
        "cert": {
          "ca": "some-mbus-ca",
          "certificate": "some-mbus-cert",
          "private_key": "some-mbus-private-key"
        }
      },
      "password": "some-password"
    }
  }
}
			`))
		})
	})
})
