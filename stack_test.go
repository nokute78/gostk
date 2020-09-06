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

package stack_test

import (
	"github.com/nokute78/gostk"
	"testing"
)

func TestPushPop(t *testing.T) {
	s := stack.New()
	len := 32

	for i := 0; i < len; i++ {
		s.Push(i)
	}

	for i := 0; i < len; i++ {
		v, err := s.Pop()
		if err != nil {
			t.Fatalf("%d:err=%s", i, err)
		}
		if v.(int) != len-i-1 {
			t.Errorf("%d:given=%d expect=%d", i, v.(int), len-i-1)
		}
	}
}

func checkPopErr(t *testing.T, given interface{}, expect interface{}, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("err=%s", err)
	}

	switch v := given.(type) {
	case int:
		ex, ok := expect.(int)
		if !ok {
			t.Fatalf("can not convert given")
		}
		if ex != v {
			t.Fatalf("value mismatch given=%d expect=%d", v, ex)
		}

	default:
		t.Fatalf("not supported given=%v", given)
	}

}

func TestPushPop2(t *testing.T) {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Pop()
	checkPopErr(t, v, 3, err)
	v, err = s.Pop()
	checkPopErr(t, v, 2, err)

	s.Push(4)
	v, err = s.Pop()
	checkPopErr(t, v, 4, err)
	v, err = s.Pop()
	checkPopErr(t, v, 1, err)

	v, err = s.Pop()
	if err == nil {
		t.Fatal("err is nil?")
	}
	v, err = s.Pop()
	if err == nil {
		t.Fatal("err is nil?")
	}

	s.Push(5)
	s.Push(6)

	v, err = s.Pop()
	checkPopErr(t, v, 6, err)
	v, err = s.Pop()
	checkPopErr(t, v, 5, err)

}

func TestReset(t *testing.T) {
	s := stack.New()
	len := 32
	for i := 0; i < len; i++ {
		s.Push(i)
	}
	if s.Len() != len {
		t.Errorf("size mismatch: given=%d expect=%d", s.Len(), len)
	}
	s.Reset()
	if s.Len() != 0 {
		t.Errorf("size mismatch: given=%d expect=%d", s.Len(), 0)
	}

}

func TestEmptyPop(t *testing.T) {
	s := stack.New()

	for i := 0; i < 10; i++ {
		_, err := s.Pop()
		if err == nil {
			t.Fatalf("%d: err is nil?", i)
		}
	}
}

func TestLen(t *testing.T) {
	s := stack.New()

	if s.Len() != 0 {
		t.Errorf("stack size is not 0 given=%d", s.Len())
	}

	count := 4
	for i := 0; i < count; i++ {
		s.Push(i)
	}

	if s.Len() != count {
		t.Errorf("stack size error: given=%d expected=%d", s.Len(), count)
	}

	s.Pop()
	s.Pop()

	if s.Len() != count-2 {
		t.Errorf("stack size error: given=%d expected=%d", s.Len(), count-2)
	}

}

func TestCopy(t *testing.T) {
	s := stack.New()
	size := 4

	for i := 0; i < size; i++ {
		s.Push(i)
	}

	cpy := s.Copy()
	cpy2 := s.Copy()

	len := s.Len()
	for i := 0; i < len; i++ {
		sv, err := s.Pop()
		if err != nil {
			t.Fatalf("%d: err=%s", i, err)
		}
		cpyv, err := cpy.Pop()
		if err != nil {
			t.Fatalf("%d: err=%s", i, err)
		}
		if sv.(int) != cpyv.(int) {
			t.Fatalf("%d:mismatch src=%d cpy=%d", i, sv.(int), cpyv.(int))
		}
	}

	if cpy2.Len() != size {
		t.Fatalf("size error given=%d expect=%d", cpy2.Len(), size)
	}
}
