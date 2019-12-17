package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter = 0
	aux     = uint32(0)
	lock    sync.Mutex

	atomicCounter = AtomicInt{}
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *AtomicInt) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

func (i *AtomicInt) Value() int {
	return i.value
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(3)
		go updateCounterAtomically(&wg)
		go updateCounter(&wg)
		go updateAtomicCounter(&wg)
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(aux)
	fmt.Println(atomicCounter.Value())

}

func updateCounterAtomically(wg *sync.WaitGroup) {

	atomic.AddUint32(&aux, 1)
	wg.Done()
}

func updateAtomicCounter(wg *sync.WaitGroup) {
	atomicCounter.Increase()
	wg.Done()
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()

	counter++
	wg.Done()
}
