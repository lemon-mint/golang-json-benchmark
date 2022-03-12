package main

import (
	"encoding/json"
	"math"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

type JSONStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Lang string `json:"lang"`
	ID   int    `json:"id"`

	Uint   uint   `json:"uint"`
	Int    int    `json:"int"`
	Uint8  uint8  `json:"uint8"`
	Uint16 uint16 `json:"uint16"`
	Uint32 uint32 `json:"uint32"`
	Uint64 uint64 `json:"uint64"`
	Int8   int8   `json:"int8"`
	Int16  int16  `json:"int16"`
	Int32  int32  `json:"int32"`
	Int64  int64  `json:"int64"`

	Float32 float32 `json:"float32"`
	Float64 float64 `json:"float64"`
}

var jsonStruct = func() JSONStruct {
	v := JSONStruct{}

	v.Name = "John Doe"
	v.Age = 30
	v.Lang = "en"
	v.ID = 1
	v.Uint = 1<<32 - 1
	v.Int = -(1 << 31)
	v.Uint8 = 1<<8 - 1
	v.Uint16 = 1<<16 - 1
	v.Uint32 = 1<<32 - 1
	v.Uint64 = 1<<64 - 1
	v.Int8 = -(1 << 7)
	v.Int16 = -(1 << 15)
	v.Int32 = -(1 << 31)
	v.Int64 = -(1 << 63)
	v.Float32 = math.Pi
	v.Float64 = math.Pi
	return v
}()

var jsonStructBytes = func() []byte {
	b, _ := json.Marshal(jsonStruct)
	return b
}

func BenchmarkStdJSONMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := json.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkStdJSONUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data := jsonStructBytes()
			var v JSONStruct
			err := json.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkJsoniterMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := jsoniter.ConfigFastest.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkJsoniterUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data := jsonStructBytes()
			var v JSONStruct
			err := jsoniter.ConfigFastest.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGoJsonMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := gojson.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkGoJsonUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data := jsonStructBytes()
			var v JSONStruct
			err := gojson.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
