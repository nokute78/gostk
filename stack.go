/*
   Copyright 2020 Takahiro Yamashita

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package stack

import (
	"errors"
	"sync"
)

var ErrEmpty error = errors.New("stack is empty")

// Stack represents stack.
type Stack struct {
	stk    []interface{}
	pos    int
	stklen int
	mu     sync.RWMutex
}

// New returns a pointer of Stack struct.
func New() *Stack {
	return NewLen(16)
}

// NewLen returns a pointer of Stack struct.
// len can set the default size of stack.
func NewLen(len int) *Stack {
	return &Stack{stk: make([]interface{}, len), stklen: len}
}

// Reset resets the stack.
// The internal storage is not release for use by future writes.
func (s *Stack) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stk = s.stk[:0]
	s.pos = 0
}

// Len returns the current size of stack.
func (s *Stack) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.pos
}

// Push saves the v at the top of stack.
func (s *Stack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stklen <= s.pos {
		s.stk = append(s.stk, v)
		s.stklen++
	} else {
		s.stk[s.pos] = v
	}
	s.pos++
}

// Pop returns the value of top of stack.
func (s *Stack) Pop() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.pos == 0 {
		return nil, ErrEmpty
	}
	s.pos--
	return s.stk[s.pos], nil
}

// Copy copies stack and returns its pointer.
func (s Stack) Copy() *Stack {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ret := &Stack{pos: s.pos, stklen: s.stklen}
	ret.stk = make([]interface{}, len(s.stk), cap(s.stk))
	copy(ret.stk, s.stk)
	return ret
}
