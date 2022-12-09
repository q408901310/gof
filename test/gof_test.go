package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"reflect"
	"sync"
	"testing"

	. "github.com/gogf/gf/v2/test/gtest"
)

type (
	RWMutex struct {
		*sync.RWMutex
	}

	async struct {
		mu RWMutex
	}

	obj struct {
		as *async
	}
)

func Test_Async(t *testing.T) {
	C(t, func(t *T) {
		//mm := RWMutex{}
		//mm.RWMutex = new(sync.RWMutex)
		//var o = obj{
		//	as: &async{mu: mm},
		//}
		//if o.as.mu.RWMutex != nil {
		//	g.Dump("lock")
		//	return
		//}
		//g.Dump("unlock")

		type MyInt int
		var x MyInt = 7
		v := reflect.ValueOf(x)
		g.Dump(v.Type())
	})
}
