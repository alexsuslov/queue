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

import "sync"

type Relation struct {
	sync.RWMutex
	m map[string]string
}

func (Relation *Relation) Put(id, id1 string) {
	Relation.Lock()
	defer Relation.Unlock()
	Relation.m[id] = id1
}

func (Relation *Relation) Get(id string) (string, bool) {
	Relation.RLock()
	defer Relation.RUnlock()
	id1, ok := Relation.m[id]
	return id1, ok
}

func (Relation *Relation) Remove(id string) {
	Relation.Lock()
	defer Relation.Unlock()
	delete(Relation.m, id)
}
