package paasio

import (
	"io"
	"sync"
)

const testVersion = 3

type rCounter struct {
	rm    *sync.RWMutex
	rn    int64
	rnops int
	r     io.Reader
}

func (r *rCounter) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)

	r.rm.Lock()
	defer r.rm.Unlock()

	if err != nil {
		return n, err
	}

	r.rn += int64(n)
	r.rnops++

	return n, nil
}

func (r *rCounter) ReadCount() (int64, int) {
	r.rm.RLock()
	defer r.rm.RUnlock()
	return r.rn, r.rnops
}

type wCounter struct {
	wm    *sync.RWMutex
	wn    int64
	wnops int
	w     io.Writer
}

func (w *wCounter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)

	w.wm.Lock()
	defer w.wm.Unlock()

	if err != nil {
		return n, err
	}

	w.wn += int64(n)
	w.wnops++

	return n, err
}

func (w *wCounter) WriteCount() (int64, int) {
	w.wm.RLock()
	defer w.wm.RUnlock()
	return w.wn, w.wnops
}

type rwCounter struct {
	*rCounter
	*wCounter
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &wCounter{w: w, wm: &sync.RWMutex{}}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &rCounter{r: r, rm: &sync.RWMutex{}}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		rCounter: &rCounter{r: rw, rm: &sync.RWMutex{}},
		wCounter: &wCounter{w: rw, wm: &sync.RWMutex{}},
	}
}