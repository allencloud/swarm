package swarm

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertKVStringsToMap(t *testing.T) {
	result := convertKVStringsToMap([]string{"HELLO=WORLD", "a=b=c=d", "e"})
	expected := map[string]string{"HELLO": "WORLD", "a": "b=c=d", "e": ""}
	assert.Equal(t, expected, result)
}

func TestConvertMapToKVStrings(t *testing.T) {
	result := convertMapToKVStrings(map[string]string{"HELLO": "WORLD", "a": "b=c=d", "e": ""})
	sort.Strings(result)
	expected := []string{"HELLO=WORLD", "a=b=c=d", "e="}
	assert.Equal(t, expected, result)
}

func TestMergePlugins(t *testing.T) {
	a := types.PluginsInfo{
		Volume:        []string{"local"},
		Network:       []string{"null", "bridge", "host"},
		Authorization: []string{},
	}

	b := types.PluginsInfo{
		Volume:        []string{"local"},
		Network:       []string{"null", "bridge", "host"},
		Authorization: []string{},
	}

	c := types.PluginsInfo{
		Volume:        []string{"local", "aaa"},
		Network:       []string{"null", "bridge", "host", "overlay"},
		Authorization: []string{"bbb"},
	}

	d := mergePlugins(a, b)
	assert.Equal(t, d.Volume, []string{"local"})
	assert.Equal(t, d.Network, []string{"null", "bridge", "host"})
	assert.Equal(t, d.Authorization, []string{})

	e := mergePlugins(a, c)
	assert.Equal(t, e.Volume, []string{"local", "aaa"})
	assert.Equal(t, e.Network, []string{"null", "bridge", "host", "overlay"})
	assert.Equal(t, e.Authorization, []string{"bbb"})
}
