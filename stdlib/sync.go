package stdlib

import (
	"sync"
	// "runtime"
	"github.com/2dprototype/tender"
)

var syncModule = map[string]tender.Object{
	"mutex": &tender.UserFunction{Name: "mutex", Value: syncNewMutex},
	// "lock_os_thread": &tender.BuiltinFunction{
		// Name: "lock_os_thread",
		// Value: func(args ...tender.Object) (tender.Object, error) {
			// if len(args) != 0 {
				// return nil, tender.ErrInvalidArgCount
			// }
			// runtime.LockOSThread()
			// return tender.NullValue, nil
		// },
	// },
}

func syncNewMutex(args ...tender.Object) (tender.Object, error) {
	if len(args) != 0 {
		return nil, tender.ErrWrongNumArguments
	}
	return makeMutexObj(&sync.Mutex{}), nil
}

func makeMutexObj(mu *sync.Mutex) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"lock": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					mu.Lock()
					return nil, nil
				},
			},
			"unlock": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					mu.Unlock()
					return nil, nil
				},
			},
		},
	}
}