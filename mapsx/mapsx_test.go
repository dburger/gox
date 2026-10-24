package mapsx

import (
	"testing"

	"github.com/dburger/gox/internal/testutil"
)

func TestMerge(t *testing.T) {
	map1 := map[string]string{"key1": "value1", "key2": "value2"}
	map2 := map[string]string{"key2": "value3", "key4": "value4"}
	map3 := map[string]string{"key5": "value5"}
	merged := Merge(map1, map2, map3)
	testutil.Equal(t, len(merged), 4)
	testutil.Equal(t, merged["key1"], "value1")
	testutil.Equal(t, merged["key2"], "value3")
	testutil.Equal(t, merged["key4"], "value4")
	testutil.Equal(t, merged["key5"], "value5")
	// TODO(dburger): could assert original maps unchanged
}
