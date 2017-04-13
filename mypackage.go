package mypackage
import (
 "sync/atomic"
)
var count unit64

func Count() unit64 {
atomic.AddUnit64(&count, 1)
return count
}
