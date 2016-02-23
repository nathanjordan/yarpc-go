// Copyright (c) 2016 Uber Technologies, Inc.
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

package main

import (
	"fmt"

	"github.com/yarpc/yarpc-go"
	"github.com/yarpc/yarpc-go/encoding/thrift"
	"github.com/yarpc/yarpc-go/examples/thrift/keyvalue"
	"github.com/yarpc/yarpc-go/transport"
	"github.com/yarpc/yarpc-go/transport/http"

	"golang.org/x/net/context"
)

type handler struct {
	items map[string]string
}

func (h handler) GetValue(ctx context.Context, meta yarpc.Meta, key string) (string, yarpc.Meta, error) {
	if value, ok := h.items[key]; ok {
		return value, nil, nil
	}

	return "", nil, &keyvalue.ResourceDoesNotExist{Key: key}
}

func (h handler) SetValue(ctx context.Context, meta yarpc.Meta, key string, value string) (yarpc.Meta, error) {
	h.items[key] = value
	return nil, nil
}

func main() {
	yarpc := yarpc.New(yarpc.Config{
		Name:     "keyvalue",
		Inbounds: []transport.Inbound{http.NewInbound(":8080")},
	})

	handler := handler{items: make(map[string]string)}
	thrift.Register(yarpc, keyvalue.NewKeyValueHandler(handler))

	if err := yarpc.Start(); err != nil {
		fmt.Println("error:", err.Error())
	}
}