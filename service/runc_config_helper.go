package service

import "encoding/json"

func attachBindMount(contents []byte, source string, destination string) []byte {
	var spec map[string]interface{}
	json.Unmarshal(contents, &spec)

	if mounts, ok := spec["mounts"].([]interface{}); ok {
		spec["mounts"] = append(mounts, map[string]interface{}{
			"source":      source,
			"destination": destination,
			"type":        "bind",
			"options": []string{
				"rw",
				"rbind",
				"rprivate",
			},
		})
	}

	data, _ := json.Marshal(spec)
	return data
}

func detachBindMount(contents []byte, source string) []byte {
	var spec map[string]interface{}
	json.Unmarshal(contents, &spec)

	var filteredMounts []interface{}
	if mounts, ok := spec["mounts"].([]map[string]interface{}); ok {
		for _, mount := range mounts {
			if src, ok := mount["source"]; ok {
				if src != source {
					filteredMounts = append(filteredMounts, mount)
				}
			}
		}
	}

	spec["mounts"] = filteredMounts

	data, _ := json.Marshal(spec)
	return data
}