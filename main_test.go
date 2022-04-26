package main

import (
	"encoding/json"
	"math"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	segmentio "github.com/segmentio/encoding/json"
	"github.com/wI2L/jettison"
)

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
		data := jsonStructBytes()
		for p.Next() {
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
		data := jsonStructBytes()
		for p.Next() {
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
		data := jsonStructBytes()
		for p.Next() {
			var v JSONStruct
			err := gojson.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkJettisonMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := jettison.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkSegmentioMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := segmentio.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkSegmentioUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		data := jsonStructBytes()
		for p.Next() {
			var v JSONStruct
			err := segmentio.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
