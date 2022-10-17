package rootcloser

import "fmt"

var closeFuncs []func()

func Register(closeFn func()) {
	closeFuncs = append(closeFuncs, closeFn)
}

func Execute() {
	for _, closeFn := range closeFuncs {
		defer recoverRootCloser()
		closeFn()
	}
}

func recoverRootCloser() {
	if err := recover(); err == nil {
		return
	} else {
		var (
			err error
			ok  bool
		)
		err, ok = err.(error)
		if !ok {
			err = fmt.Errorf("%v", err)
		}
		fmt.Printf("execute root defer failed | err=%s", err.Error())
	}
}
