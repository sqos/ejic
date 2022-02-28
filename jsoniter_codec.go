// Copyright 2015-2022 sqos. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ejic

import (
	"github.com/henrylee2cn/erpc/v6/codec"
	jsoniter "github.com/json-iterator/go"
)

// jsoniter codec name and id
const (
	NAME_JSONITER = "jsoniter"
	ID_JSONITER   = 'J'
)

func init() {
	codec.Reg(new(JsoniterCodec))
}

// JsoniterCodec jsoniter codec
type JsoniterCodec struct{}

// Name returns codec name.
func (JsoniterCodec) Name() string {
	return NAME_JSONITER
}

// ID returns codec id.
func (JsoniterCodec) ID() byte {
	return ID_JSONITER
}

// Marshal returns the JSON encoding of v.
func (JsoniterCodec) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func (JsoniterCodec) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}
