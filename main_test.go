package main

import (
	"encoding/json"
	"math"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/lemon-mint/golang-json-benchmark/types"
	segmentio "github.com/segmentio/encoding/json"
	shamaton_msgpack "github.com/shamaton/msgpack/v2"
	shamaton_msgpackgen "github.com/shamaton/msgpackgen/msgpack"
	vmihailenco_msgpack "github.com/vmihailenco/msgpack/v5"
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

func BenchmarkVstructMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			// func types.New_JSONStructVstruct(Age int64, ID int64, Uint uint64, Int int64, Uint8 uint8, Uint16 uint16, Uint32 uint32, Uint64 uint64, Int8 int8, Int16 int16, Int32 int32, Int64 int64, Float32 float32, Float64 float64, Name string, Lang string) types.JSONStructVstruct
			types.New_JSONStructVstruct(
				int64(jsonStruct.Age),
				int64(jsonStruct.ID),
				uint64(jsonStruct.Uint),
				int64(jsonStruct.Int),
				jsonStruct.Uint8,
				jsonStruct.Uint16,
				jsonStruct.Uint32,
				jsonStruct.Uint64,
				jsonStruct.Int8,
				jsonStruct.Int16,
				jsonStruct.Int32,
				jsonStruct.Int64,
				jsonStruct.Float32,
				jsonStruct.Float64,
				jsonStruct.Name,
				jsonStruct.Lang,
			)
		}
	})
}

var jsonStructVstructBytes = types.New_JSONStructVstruct(
	int64(jsonStruct.Age),
	int64(jsonStruct.ID),
	uint64(jsonStruct.Uint),
	int64(jsonStruct.Int),
	jsonStruct.Uint8,
	jsonStruct.Uint16,
	jsonStruct.Uint32,
	jsonStruct.Uint64,
	jsonStruct.Int8,
	jsonStruct.Int16,
	jsonStruct.Int32,
	jsonStruct.Int64,
	jsonStruct.Float32,
	jsonStruct.Float64,
	jsonStruct.Name,
	jsonStruct.Lang,
)

func BenchmarkVstructUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		data := jsonStructVstructBytes
		for p.Next() {
			if !data.Vstruct_Validate() {
				b.Fatal("validation failed")
			}
			_ = data.Age()
			_ = data.ID()
			_ = data.Uint()
			_ = data.Int()
			_ = data.Uint8()
			_ = data.Uint16()
			_ = data.Uint32()
			_ = data.Uint64()
			_ = data.Int8()
			_ = data.Int16()
			_ = data.Int32()
			_ = data.Int64()
			_ = data.Float32()
			_ = data.Float64()
			_ = data.Name()
			_ = data.Lang()
		}
	})
}

func BenchmarkVmihailencoMsgpackMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := vmihailenco_msgpack.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkVmihailencoMsgpackUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		data, err := vmihailenco_msgpack.Marshal(jsonStruct)
		if err != nil {
			b.Fatal(err)
		}
		for p.Next() {
			var v JSONStruct
			err := vmihailenco_msgpack.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkShamatonMsgpackMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := shamaton_msgpack.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkShamatonMsgpackUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		data, err := shamaton_msgpack.Marshal(jsonStruct)
		if err != nil {
			b.Fatal(err)
		}
		for p.Next() {
			var v JSONStruct
			err := shamaton_msgpack.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

//go:generate go get github.com/shamaton/msgpackgen/internal/generator
//go:generate go run github.com/shamaton/msgpackgen
//go:generate go mod tidy

func BenchmarkShamatonMsgpackGenMarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			data, err := shamaton_msgpackgen.Marshal(jsonStruct)
			if err != nil {
				b.Fatal(err)
			}
			_ = data
		}
	})
}

func BenchmarkShamatonMsgpackGenUnmarshal(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		data, err := shamaton_msgpackgen.Marshal(jsonStruct)
		if err != nil {
			b.Fatal(err)
		}
		for p.Next() {
			var v JSONStruct
			err := shamaton_msgpackgen.Unmarshal(data, &v)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
