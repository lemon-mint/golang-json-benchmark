// Code generated by msgpackgen. DO NOT EDIT.

package main

import (
	"fmt"
	types "github.com/lemon-mint/golang-json-benchmark/types"
	msgpack "github.com/shamaton/msgpackgen/msgpack"
	dec "github.com/shamaton/msgpackgen/msgpack/dec"
	enc "github.com/shamaton/msgpackgen/msgpack/enc"
)

// RegisterGeneratedResolver registers generated resolver.
func RegisterGeneratedResolver() {
	msgpack.SetResolver(___encodeAsMap, ___encodeAsArray, ___decodeAsMap, ___decodeAsArray)
}

// encode
func ___encode(i interface{}) ([]byte, error) {
	if msgpack.StructAsArray() {
		return ___encodeAsArray(i)
	} else {
		return ___encodeAsMap(i)
	}
}

// encodeAsArray
func ___encodeAsArray(i interface{}) ([]byte, error) {
	switch v := i.(type) {
	case types.JSONStruct:
		encoder := enc.NewEncoder()
		size, err := ___calcArraySizeJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeArrayJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "github.com/lemon-mint/golang-json-benchmark/types.JSONStruct", size, offset)
		}
		return b, err
	case *types.JSONStruct:
		encoder := enc.NewEncoder()
		size, err := ___calcArraySizeJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(*v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeArrayJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(*v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "github.com/lemon-mint/golang-json-benchmark/types.JSONStruct", size, offset)
		}
		return b, err
	}
	return nil, nil
}

// encodeAsMap
func ___encodeAsMap(i interface{}) ([]byte, error) {
	switch v := i.(type) {
	case types.JSONStruct:
		encoder := enc.NewEncoder()
		size, err := ___calcMapSizeJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeMapJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "github.com/lemon-mint/golang-json-benchmark/types.JSONStruct", size, offset)
		}
		return b, err
	case *types.JSONStruct:
		encoder := enc.NewEncoder()
		size, err := ___calcMapSizeJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(*v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeMapJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(*v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "github.com/lemon-mint/golang-json-benchmark/types.JSONStruct", size, offset)
		}
		return b, err
	}
	return nil, nil
}

// decode
func ___decode(data []byte, i interface{}) (bool, error) {
	if msgpack.StructAsArray() {
		return ___decodeAsArray(data, i)
	} else {
		return ___decodeAsMap(data, i)
	}
}

// decodeAsArray
func ___decodeAsArray(data []byte, i interface{}) (bool, error) {
	switch v := i.(type) {
	case *types.JSONStruct:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeArrayJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	case **types.JSONStruct:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeArrayJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(*v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	}
	return false, nil
}

// decodeAsMap
func ___decodeAsMap(data []byte, i interface{}) (bool, error) {
	switch v := i.(type) {
	case *types.JSONStruct:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeMapJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	case **types.JSONStruct:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeMapJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(*v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	}
	return false, nil
}

// calculate size from github.com/lemon-mint/golang-json-benchmark/types.JSONStruct
func ___calcArraySizeJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v types.JSONStruct, encoder *enc.Encoder) (int, error) {
	size := 0
	size += encoder.CalcStructHeader16(16)
	size += encoder.CalcString(v.Name)
	size += encoder.CalcInt(v.Age)
	size += encoder.CalcString(v.Lang)
	size += encoder.CalcInt(v.ID)
	size += encoder.CalcUint(v.Uint)
	size += encoder.CalcInt(v.Int)
	size += encoder.CalcUint8(v.Uint8)
	size += encoder.CalcUint16(v.Uint16)
	size += encoder.CalcUint32(v.Uint32)
	size += encoder.CalcUint64(v.Uint64)
	size += encoder.CalcInt8(v.Int8)
	size += encoder.CalcInt16(v.Int16)
	size += encoder.CalcInt32(v.Int32)
	size += encoder.CalcInt64(v.Int64)
	size += encoder.CalcFloat32(v.Float32)
	size += encoder.CalcFloat64(v.Float64)
	return size, nil
}

// calculate size from github.com/lemon-mint/golang-json-benchmark/types.JSONStruct
func ___calcMapSizeJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v types.JSONStruct, encoder *enc.Encoder) (int, error) {
	size := 0
	size += encoder.CalcStructHeader16(16)
	size += encoder.CalcStringFix(4)
	size += encoder.CalcString(v.Name)
	size += encoder.CalcStringFix(3)
	size += encoder.CalcInt(v.Age)
	size += encoder.CalcStringFix(4)
	size += encoder.CalcString(v.Lang)
	size += encoder.CalcStringFix(2)
	size += encoder.CalcInt(v.ID)
	size += encoder.CalcStringFix(4)
	size += encoder.CalcUint(v.Uint)
	size += encoder.CalcStringFix(3)
	size += encoder.CalcInt(v.Int)
	size += encoder.CalcStringFix(5)
	size += encoder.CalcUint8(v.Uint8)
	size += encoder.CalcStringFix(6)
	size += encoder.CalcUint16(v.Uint16)
	size += encoder.CalcStringFix(6)
	size += encoder.CalcUint32(v.Uint32)
	size += encoder.CalcStringFix(6)
	size += encoder.CalcUint64(v.Uint64)
	size += encoder.CalcStringFix(4)
	size += encoder.CalcInt8(v.Int8)
	size += encoder.CalcStringFix(5)
	size += encoder.CalcInt16(v.Int16)
	size += encoder.CalcStringFix(5)
	size += encoder.CalcInt32(v.Int32)
	size += encoder.CalcStringFix(5)
	size += encoder.CalcInt64(v.Int64)
	size += encoder.CalcStringFix(7)
	size += encoder.CalcFloat32(v.Float32)
	size += encoder.CalcStringFix(7)
	size += encoder.CalcFloat64(v.Float64)
	return size, nil
}

// encode from github.com/lemon-mint/golang-json-benchmark/types.JSONStruct
func ___encodeArrayJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v types.JSONStruct, encoder *enc.Encoder, offset int) ([]byte, int, error) {
	var err error
	offset = encoder.WriteStructHeader16AsArray(16, offset)
	offset = encoder.WriteString(v.Name, offset)
	offset = encoder.WriteInt(v.Age, offset)
	offset = encoder.WriteString(v.Lang, offset)
	offset = encoder.WriteInt(v.ID, offset)
	offset = encoder.WriteUint(v.Uint, offset)
	offset = encoder.WriteInt(v.Int, offset)
	offset = encoder.WriteUint8(v.Uint8, offset)
	offset = encoder.WriteUint16(v.Uint16, offset)
	offset = encoder.WriteUint32(v.Uint32, offset)
	offset = encoder.WriteUint64(v.Uint64, offset)
	offset = encoder.WriteInt8(v.Int8, offset)
	offset = encoder.WriteInt16(v.Int16, offset)
	offset = encoder.WriteInt32(v.Int32, offset)
	offset = encoder.WriteInt64(v.Int64, offset)
	offset = encoder.WriteFloat32(v.Float32, offset)
	offset = encoder.WriteFloat64(v.Float64, offset)
	return encoder.EncodedBytes(), offset, err
}

// encode from github.com/lemon-mint/golang-json-benchmark/types.JSONStruct
func ___encodeMapJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v types.JSONStruct, encoder *enc.Encoder, offset int) ([]byte, int, error) {
	var err error
	offset = encoder.WriteStructHeader16AsMap(16, offset)
	offset = encoder.WriteStringFix("Name", 4, offset)
	offset = encoder.WriteString(v.Name, offset)
	offset = encoder.WriteStringFix("Age", 3, offset)
	offset = encoder.WriteInt(v.Age, offset)
	offset = encoder.WriteStringFix("Lang", 4, offset)
	offset = encoder.WriteString(v.Lang, offset)
	offset = encoder.WriteStringFix("ID", 2, offset)
	offset = encoder.WriteInt(v.ID, offset)
	offset = encoder.WriteStringFix("Uint", 4, offset)
	offset = encoder.WriteUint(v.Uint, offset)
	offset = encoder.WriteStringFix("Int", 3, offset)
	offset = encoder.WriteInt(v.Int, offset)
	offset = encoder.WriteStringFix("Uint8", 5, offset)
	offset = encoder.WriteUint8(v.Uint8, offset)
	offset = encoder.WriteStringFix("Uint16", 6, offset)
	offset = encoder.WriteUint16(v.Uint16, offset)
	offset = encoder.WriteStringFix("Uint32", 6, offset)
	offset = encoder.WriteUint32(v.Uint32, offset)
	offset = encoder.WriteStringFix("Uint64", 6, offset)
	offset = encoder.WriteUint64(v.Uint64, offset)
	offset = encoder.WriteStringFix("Int8", 4, offset)
	offset = encoder.WriteInt8(v.Int8, offset)
	offset = encoder.WriteStringFix("Int16", 5, offset)
	offset = encoder.WriteInt16(v.Int16, offset)
	offset = encoder.WriteStringFix("Int32", 5, offset)
	offset = encoder.WriteInt32(v.Int32, offset)
	offset = encoder.WriteStringFix("Int64", 5, offset)
	offset = encoder.WriteInt64(v.Int64, offset)
	offset = encoder.WriteStringFix("Float32", 7, offset)
	offset = encoder.WriteFloat32(v.Float32, offset)
	offset = encoder.WriteStringFix("Float64", 7, offset)
	offset = encoder.WriteFloat64(v.Float64, offset)
	return encoder.EncodedBytes(), offset, err
}

// decode to github.com/lemon-mint/golang-json-benchmark/types.JSONStruct
func ___decodeArrayJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v *types.JSONStruct, decoder *dec.Decoder, offset int) (int, error) {
	offset, err := decoder.CheckStructHeader(16, offset)
	if err != nil {
		return 0, err
	}
	{
		var vv string
		vv, offset, err = decoder.AsString(offset)
		if err != nil {
			return 0, err
		}
		v.Name = vv
	}
	{
		var vv int
		vv, offset, err = decoder.AsInt(offset)
		if err != nil {
			return 0, err
		}
		v.Age = vv
	}
	{
		var vv string
		vv, offset, err = decoder.AsString(offset)
		if err != nil {
			return 0, err
		}
		v.Lang = vv
	}
	{
		var vv int
		vv, offset, err = decoder.AsInt(offset)
		if err != nil {
			return 0, err
		}
		v.ID = vv
	}
	{
		var vv uint
		vv, offset, err = decoder.AsUint(offset)
		if err != nil {
			return 0, err
		}
		v.Uint = vv
	}
	{
		var vv int
		vv, offset, err = decoder.AsInt(offset)
		if err != nil {
			return 0, err
		}
		v.Int = vv
	}
	{
		var vv uint8
		vv, offset, err = decoder.AsUint8(offset)
		if err != nil {
			return 0, err
		}
		v.Uint8 = vv
	}
	{
		var vv uint16
		vv, offset, err = decoder.AsUint16(offset)
		if err != nil {
			return 0, err
		}
		v.Uint16 = vv
	}
	{
		var vv uint32
		vv, offset, err = decoder.AsUint32(offset)
		if err != nil {
			return 0, err
		}
		v.Uint32 = vv
	}
	{
		var vv uint64
		vv, offset, err = decoder.AsUint64(offset)
		if err != nil {
			return 0, err
		}
		v.Uint64 = vv
	}
	{
		var vv int8
		vv, offset, err = decoder.AsInt8(offset)
		if err != nil {
			return 0, err
		}
		v.Int8 = vv
	}
	{
		var vv int16
		vv, offset, err = decoder.AsInt16(offset)
		if err != nil {
			return 0, err
		}
		v.Int16 = vv
	}
	{
		var vv int32
		vv, offset, err = decoder.AsInt32(offset)
		if err != nil {
			return 0, err
		}
		v.Int32 = vv
	}
	{
		var vv int64
		vv, offset, err = decoder.AsInt64(offset)
		if err != nil {
			return 0, err
		}
		v.Int64 = vv
	}
	{
		var vv float32
		vv, offset, err = decoder.AsFloat32(offset)
		if err != nil {
			return 0, err
		}
		v.Float32 = vv
	}
	{
		var vv float64
		vv, offset, err = decoder.AsFloat64(offset)
		if err != nil {
			return 0, err
		}
		v.Float64 = vv
	}
	return offset, err
}

// decode to github.com/lemon-mint/golang-json-benchmark/types.JSONStruct
func ___decodeMapJSONStruct_f54ceef57ff0a7b2d54f427103e01334f2f3f184d7886f94018db35014044b06(v *types.JSONStruct, decoder *dec.Decoder, offset int) (int, error) {
	keys := [][]byte{
		{uint8(0x4e), uint8(0x61), uint8(0x6d), uint8(0x65)},                                        // Name
		{uint8(0x41), uint8(0x67), uint8(0x65)},                                                     // Age
		{uint8(0x4c), uint8(0x61), uint8(0x6e), uint8(0x67)},                                        // Lang
		{uint8(0x49), uint8(0x44)},                                                                  // ID
		{uint8(0x55), uint8(0x69), uint8(0x6e), uint8(0x74)},                                        // Uint
		{uint8(0x49), uint8(0x6e), uint8(0x74)},                                                     // Int
		{uint8(0x55), uint8(0x69), uint8(0x6e), uint8(0x74), uint8(0x38)},                           // Uint8
		{uint8(0x55), uint8(0x69), uint8(0x6e), uint8(0x74), uint8(0x31), uint8(0x36)},              // Uint16
		{uint8(0x55), uint8(0x69), uint8(0x6e), uint8(0x74), uint8(0x33), uint8(0x32)},              // Uint32
		{uint8(0x55), uint8(0x69), uint8(0x6e), uint8(0x74), uint8(0x36), uint8(0x34)},              // Uint64
		{uint8(0x49), uint8(0x6e), uint8(0x74), uint8(0x38)},                                        // Int8
		{uint8(0x49), uint8(0x6e), uint8(0x74), uint8(0x31), uint8(0x36)},                           // Int16
		{uint8(0x49), uint8(0x6e), uint8(0x74), uint8(0x33), uint8(0x32)},                           // Int32
		{uint8(0x49), uint8(0x6e), uint8(0x74), uint8(0x36), uint8(0x34)},                           // Int64
		{uint8(0x46), uint8(0x6c), uint8(0x6f), uint8(0x61), uint8(0x74), uint8(0x33), uint8(0x32)}, // Float32
		{uint8(0x46), uint8(0x6c), uint8(0x6f), uint8(0x61), uint8(0x74), uint8(0x36), uint8(0x34)}, // Float64
	}
	offset, err := decoder.CheckStructHeader(16, offset)
	if err != nil {
		return 0, err
	}
	count := 0
	for count < 16 {
		var dataKey []byte
		dataKey, offset, err = decoder.AsStringBytes(offset)
		if err != nil {
			return 0, err
		}
		fieldIndex := -1
		for i, key := range keys {
			if len(dataKey) != len(key) {
				continue
			}
			fieldIndex = i
			for dataKeyIndex := range dataKey {
				if dataKey[dataKeyIndex] != key[dataKeyIndex] {
					fieldIndex = -1
					break
				}
			}
			if fieldIndex >= 0 {
				break
			}
		}
		switch fieldIndex {
		case 0:
			{
				var vv string
				vv, offset, err = decoder.AsString(offset)
				if err != nil {
					return 0, err
				}
				v.Name = vv
			}
			count++
		case 1:
			{
				var vv int
				vv, offset, err = decoder.AsInt(offset)
				if err != nil {
					return 0, err
				}
				v.Age = vv
			}
			count++
		case 2:
			{
				var vv string
				vv, offset, err = decoder.AsString(offset)
				if err != nil {
					return 0, err
				}
				v.Lang = vv
			}
			count++
		case 3:
			{
				var vv int
				vv, offset, err = decoder.AsInt(offset)
				if err != nil {
					return 0, err
				}
				v.ID = vv
			}
			count++
		case 4:
			{
				var vv uint
				vv, offset, err = decoder.AsUint(offset)
				if err != nil {
					return 0, err
				}
				v.Uint = vv
			}
			count++
		case 5:
			{
				var vv int
				vv, offset, err = decoder.AsInt(offset)
				if err != nil {
					return 0, err
				}
				v.Int = vv
			}
			count++
		case 6:
			{
				var vv uint8
				vv, offset, err = decoder.AsUint8(offset)
				if err != nil {
					return 0, err
				}
				v.Uint8 = vv
			}
			count++
		case 7:
			{
				var vv uint16
				vv, offset, err = decoder.AsUint16(offset)
				if err != nil {
					return 0, err
				}
				v.Uint16 = vv
			}
			count++
		case 8:
			{
				var vv uint32
				vv, offset, err = decoder.AsUint32(offset)
				if err != nil {
					return 0, err
				}
				v.Uint32 = vv
			}
			count++
		case 9:
			{
				var vv uint64
				vv, offset, err = decoder.AsUint64(offset)
				if err != nil {
					return 0, err
				}
				v.Uint64 = vv
			}
			count++
		case 10:
			{
				var vv int8
				vv, offset, err = decoder.AsInt8(offset)
				if err != nil {
					return 0, err
				}
				v.Int8 = vv
			}
			count++
		case 11:
			{
				var vv int16
				vv, offset, err = decoder.AsInt16(offset)
				if err != nil {
					return 0, err
				}
				v.Int16 = vv
			}
			count++
		case 12:
			{
				var vv int32
				vv, offset, err = decoder.AsInt32(offset)
				if err != nil {
					return 0, err
				}
				v.Int32 = vv
			}
			count++
		case 13:
			{
				var vv int64
				vv, offset, err = decoder.AsInt64(offset)
				if err != nil {
					return 0, err
				}
				v.Int64 = vv
			}
			count++
		case 14:
			{
				var vv float32
				vv, offset, err = decoder.AsFloat32(offset)
				if err != nil {
					return 0, err
				}
				v.Float32 = vv
			}
			count++
		case 15:
			{
				var vv float64
				vv, offset, err = decoder.AsFloat64(offset)
				if err != nil {
					return 0, err
				}
				v.Float64 = vv
			}
			count++
		default:
			return 0, fmt.Errorf("unknown key[%s] found", string(dataKey))
		}
	}
	return offset, err
}
