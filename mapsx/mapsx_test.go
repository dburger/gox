package mapsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	map1 := map[string]string{"key1": "value1", "key2": "value2"}
	map2 := map[string]string{"key2": "value3", "key4": "value4"}
	map3 := map[string]string{"key5": "value5"}
	merged := Merge(map1, map2, map3)
	assert.Equal(t, 4, len(merged))
	assert.Equal(t, "value1", merged["key1"])
	assert.Equal(t, "value3", merged["key2"])
	assert.Equal(t, "value4", merged["key4"])
	assert.Equal(t, "value5", merged["key5"])
	// TODO(dburger): could assert original maps unchanged
}
