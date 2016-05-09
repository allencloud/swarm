package swarm

import (
	"strings"

	"github.com/docker/engine-api/types"
)

// convertKVStringsToMap converts ["key=value"] to {"key":"value"}
func convertKVStringsToMap(values []string) map[string]string {
	result := make(map[string]string, len(values))
	for _, value := range values {
		kv := strings.SplitN(value, "=", 2)
		if len(kv) == 1 {
			result[kv[0]] = ""
		} else {
			result[kv[0]] = kv[1]
		}
	}

	return result
}

// convertMapToKVStrings converts {"key": "value"} to ["key=value"]
func convertMapToKVStrings(values map[string]string) []string {
	result := make([]string, len(values))
	i := 0
	for key, value := range values {
		result[i] = key + "=" + value
		i++
	}
	return result
}

func mergePlugins(a, b types.PluginsInfo) types.PluginsInfo {

	var (
		found  bool
		i, j   int
		length int
	)

	// merge Volume of two PluginsInfos
	length = len(a.Volume)
	for i = 0; i < len(b.Volume); i++ {
		found = false
		for j = 0; j < length; j++ {
			if a.Volume[j] == b.Volume[i] {
				found = true
				break
			}
		}
		if found == false {
			a.Volume = append(a.Volume, b.Volume[i])
		}
	}

	// merge Network of two PluginsInfos
	length = len(a.Network)
	for i = 0; i < len(b.Network); i++ {
		found = false
		for j = 0; j < length; j++ {
			if a.Network[j] == b.Network[i] {
				found = true
				break
			}
		}
		if found == false {
			a.Network = append(a.Network, b.Network[i])
		}
	}

	// merge Authorization of two PluginsInfos
	length = len(a.Authorization)
	for i = 0; i < len(b.Authorization); i++ {
		found = false
		for j = 0; j < length; j++ {
			if a.Authorization[j] == b.Authorization[i] {
				found = true
				break
			}
		}
		if found == false {
			a.Authorization = append(a.Authorization, b.Authorization[i])
		}
	}

	return a
}
