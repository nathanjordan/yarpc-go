// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpcclient

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/yarpc/v2"
)

func TestProvider(t *testing.T) {
	p := NewProvider()
	require.NotNil(t, p)

	client, ok := p.Client("foo")
	require.False(t, ok)
	require.Equal(t, yarpc.Client{}, client)

	p.Register("foo", yarpc.Client{
		Caller: "foo-caller",
	})

	client, ok = p.Client("foo")
	require.True(t, ok)

	assert.Equal(t, client, yarpc.Client{
		Caller: "foo-caller",
	})
}

func TestProviderConccurent(t *testing.T) {
	// this test is intended to be run with the race detector

	p := NewProvider()
	require.NotNil(t, p)

	var wg sync.WaitGroup
	start := make(chan struct{}, 0)

	// initialize reads (even) and writes (odds)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			<-start
			if i%2 == 0 {
				_, _ = p.Client("foo")
			} else {
				p.Register("foo", yarpc.Client{})
			}

			wg.Done()
		}(i)
	}

	close(start)
	wg.Wait()
}
