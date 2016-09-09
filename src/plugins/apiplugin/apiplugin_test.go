package apiplugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	name1 := "test1"
	dependency1 := []string{"db", "cache"}
	name2 := "test2"
	dependency2 := []string{"cache"}
	pluginTest1 := &ApiPlugin{
		Name:         name1,
		Dependencies: dependency1,
	}

	pluginTest2 := &ApiPlugin{
		Name:         name2,
		Dependencies: dependency2,
	}

	Add(pluginTest1)
	defer delete(ApiPlugins, name1)
	Add(pluginTest2)
	defer delete(ApiPlugins, name2)

	assert.NotNil(t, ApiPlugins)
	assert.Equal(t, pluginTest1, ApiPlugins[name1])
	assert.Equal(t, pluginTest2, ApiPlugins[name2])
}
