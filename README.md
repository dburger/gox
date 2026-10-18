# gox

Go extras: packages for what the standard library leaves out, named after the standard package
each extends.

- `slicesx`: `Map`, `Filter`, `Reduce`, `Partition`, `Distinct` and `Generate`.
- `mapsx`: `Merge`.
- `coll`: the collection types `Set`, `MultiSet` and `Deque`.

```go
import (
	"github.com/dburger/gox/coll"
	"github.com/dburger/gox/slicesx"
)

names := slicesx.Map(players, func(p Player) string { return p.Name })
seen := coll.NewSet(names...)
```

Formerly gocoll; gofunc's functions moved here too (2026-10-18).
