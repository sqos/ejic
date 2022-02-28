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
	"reflect"
	"testing"
)

func TestJsoniterCodec_ID(t *testing.T) {
	tests := []struct {
		name string
		want byte
	}{
		{name: "JsoniterCodecNameTest", want: 'J'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsoniterCodec{}
			if got := js.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsoniterCodec_Marshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "JsonterCodecMarshalTest",
			args:    args{v: map[string]string{"a": "hello", "b": "world"}},
			want:    []byte(`{"a":"hello","b":"world"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsoniterCodec{}
			got, err := js.Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestJsoniterCodec_Name(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "JsonterCodecNameTest", want: NAME_JSONITER},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsoniterCodec{}
			if got := js.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsoniterCodec_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "JsonterCodecUnmarshalTest",
			args: args{
				data: []byte(`{"a":"hello","b":"world"}`),
				v:    &map[string]string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsoniterCodec{}
			if err := js.Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
