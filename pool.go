// (c) 2022 Alex Suslov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package queue

import (
	"fmt"
	"sync"
)

type Pool struct {
	sync.Mutex
	Items []interface{}
	store IStore
}

type IStore interface {
	Push(value interface{}) (err error)
	Pop() (interface{}, error)
}

func (Pool *Pool) Push(value interface{}) {
	Pool.Lock()
	defer Pool.Unlock()
	Pool.Items = append(Pool.Items, value)
}

func (Pool *Pool) Pop() interface{} {
	if len(Pool.Items) == 0 {
		return nil
	}

	Pool.Lock()
	defer Pool.Unlock()

	value := Pool.Items[0]
	Pool.Items = Pool.Items[1:]
	return value
}

// Store

func (Pool *Pool) SetStore(store IStore) {
	Pool.store = store
}

func (Pool *Pool) PushS(item interface{}) error {
	if Pool.store != nil {
		return fmt.Errorf("no store")
	}
	return Pool.store.Push(item)
}

func (Pool *Pool) PopS() (interface{}, error) {
	if Pool.store != nil {
		return nil, fmt.Errorf("no store")
	}
	return Pool.store.Pop()
}
