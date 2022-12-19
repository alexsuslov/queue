//
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
	"context"
	"sync"
)

type Avoid struct {
	sync.Mutex
	fn map[string]context.CancelFunc
}

func (Avoid *Avoid) Push(key string, fn context.CancelFunc) {
	Avoid.Lock()
	defer Avoid.Unlock()
	if Avoid.fn == nil {
		Avoid.fn = map[string]context.CancelFunc{}
	}

	if fn, ok := Avoid.fn[key]; ok {
		fn()
	}
	Avoid.fn[key] = fn
}

func (Avoid *Avoid) Remove(key string) {
	Avoid.Lock()
	defer Avoid.Unlock()
	delete(Avoid.fn, key)
}
