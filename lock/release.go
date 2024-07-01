package lock

import (
	"context"

	"gopkg.in/errgo.v1"
)

func (l *EtcdLock) Release() error {
	if l == nil {
		return errgo.New("nil lock")
	}
	l.Lock()
	defer l.Unlock()
	if l.released {
		return nil
	}
	defer func() {
		l.released = true
	}()
	defer l.session.Close()
	return l.mutex.Unlock(context.Background())
}
