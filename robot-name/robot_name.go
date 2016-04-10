package robotname

import (
	"fmt"
	"math/rand"
	"sync"
)

var usedMu sync.RWMutex
var usedNames map[string]struct{} = make(map[string]struct{})

type Robot struct {
	sync.RWMutex
	once sync.Once
	name string
}

func (r *Robot) Name() string {
	r.once.Do(func() {
		r.Reset()
	})

	r.RLock()
	defer r.RUnlock()
	return r.name
}

func (r *Robot) Reset() {
	r.Lock()
	defer r.Unlock()

	ok := true
	var name string

	for ok {
		name = generateName()
		usedMu.RLock()
		_, ok = usedNames[name]
		usedMu.RUnlock()
	}

	r.name = name
	usedMu.Lock()
	usedNames[name] = struct{}{}
	usedMu.Unlock()
}

func generateName() string {
	digits := rand.Perm(10)[:3]

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	first := letters[rand.Intn(len(letters))]
	second := letters[rand.Intn(len(letters))]

	return fmt.Sprintf("%c%c%d%d%d", first, second, digits[0], digits[1], digits[2])
}
