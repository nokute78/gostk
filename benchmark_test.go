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

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack.New()
	}
}
func BenchmarkPush4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stack.New()
		for i := 0; i < 4; i++ {
			s.Push(i)
		}
	}
}

func BenchmarkPush16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stack.New()
		for i := 0; i < 16; i++ {
			s.Push(i)
		}
	}
}

func BenchmarkPush32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stack.New()
		for i := 0; i < 32; i++ {
			s.Push(i)
		}
	}
}

func BenchmarkPush256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stack.New()
		for i := 0; i < 256; i++ {
			s.Push(i)
		}
	}
}

func BenchmarkPush512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stack.New()
		for i := 0; i < 256; i++ {
			s.Push(i)
		}
	}
}

func BenchmarkPush1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stack.New()
		for i := 0; i < 1024; i++ {
			s.Push(i)
		}
	}
}
