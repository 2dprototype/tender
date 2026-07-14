package stdlib

import (
	"sync"
	// "runtime"
	"github.com/2dprototype/tender"
)

var syncModule = map[string]tender.Object{
	"mutex": &tender.NativeFunction{Name: "mutex", Value: syncNewMutex},
	// "lock_os_thread": &tender.NativeFunction{
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
			"lock": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					mu.Lock()
					return nil, nil
				},
			},
			"unlock": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					mu.Unlock()
					return nil, nil
				},
			},
		},
	}
}