package yaml

import (
	"fmt"
	"testing"

	yaml "github.com/cloudfoundry-incubator/candiedyaml"

	"github.com/stretchr/testify/assert"
)

type StructStringorslice struct {
	Foo Stringorslice
}

func TestStringorsliceYaml(t *testing.T) {
	str := `{foo: [bar, baz]}`

	s := StructStringorslice{}
	yaml.Unmarshal([]byte(str), &s)

	assert.Equal(t, Stringorslice{"bar", "baz"}, s.Foo)

	d, err := yaml.Marshal(&s)
	assert.Nil(t, err)

	s2 := StructStringorslice{}
	yaml.Unmarshal(d, &s2)

	assert.Equal(t, Stringorslice{"bar", "baz"}, s2.Foo)
}

type StructSliceorMap struct {
	Foos SliceorMap `yaml:"foos,omitempty"`
	Bars []string   `yaml:"bars"`
}

func TestSliceOrMapYaml(t *testing.T) {
	str := `{foos: [bar=baz, far=faz]}`

	s := StructSliceorMap{}
	yaml.Unmarshal([]byte(str), &s)

	assert.Equal(t, SliceorMap{"bar": "baz", "far": "faz"}, s.Foos)

	d, err := yaml.Marshal(&s)
	assert.Nil(t, err)

	s2 := StructSliceorMap{}
	yaml.Unmarshal(d, &s2)

	assert.Equal(t, SliceorMap{"bar": "baz", "far": "faz"}, s2.Foos)
}

var sampleStructSliceorMap = `
foos:
  io.rancher.os.bar: baz
  io.rancher.os.far: true
bars: []
`

func TestUnmarshalSliceOrMap(t *testing.T) {
	s := StructSliceorMap{}
	err := yaml.Unmarshal([]byte(sampleStructSliceorMap), &s)
	assert.Equal(t, fmt.Errorf("Cannot unmarshal 'true' of type bool into a string value"), err)
}

func TestStr2SliceOrMapPtrMap(t *testing.T) {
	s := map[string]*StructSliceorMap{"udav": {
		Foos: SliceorMap{"io.rancher.os.bar": "baz", "io.rancher.os.far": "true"},
		Bars: []string{},
	}}
	d, err := yaml.Marshal(&s)
	assert.Nil(t, err)

	s2 := map[string]*StructSliceorMap{}
	yaml.Unmarshal(d, &s2)

	assert.Equal(t, s, s2)
}

type StructMaporslice struct {
	Foo MaporEqualSlice
}

func contains(list []string, item string) bool {
	for _, test := range list {
		if test == item {
			return true
		}
	}
	return false
}

func TestMaporsliceYaml(t *testing.T) {
	str := `{foo: {bar: baz, far: 1, qux: null}}`

	s := StructMaporslice{}
	yaml.Unmarshal([]byte(str), &s)

	assert.Equal(t, 3, len(s.Foo))
	assert.True(t, contains(s.Foo, "bar=baz"))
	assert.True(t, contains(s.Foo, "far=1"))
	assert.True(t, contains(s.Foo, "qux"))

	d, err := yaml.Marshal(&s)
	assert.Nil(t, err)

	s2 := StructMaporslice{}
	yaml.Unmarshal(d, &s2)

	assert.Equal(t, 3, len(s2.Foo))
	assert.True(t, contains(s2.Foo, "bar=baz"))
	assert.True(t, contains(s2.Foo, "far=1"))
	assert.True(t, contains(s2.Foo, "qux"))
}
