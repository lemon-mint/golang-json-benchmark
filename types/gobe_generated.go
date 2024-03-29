// Code generated by "gobe -allow-unsafe -allow-zero-copy-string ./types"; DO NOT EDIT.
// versions:
//     gobe: v0.2.1

package types

import (
	ns25532 "math"
	ns25537 "unsafe"
)

func (ns25519 *JSONStruct) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{Name string "json:\"name\""; Age int "json:\"age\""; Lang string "json:\"lang\""; ID int "json:\"id\""; Uint uint "json:\"uint\""; Int int "json:\"int\""; Uint8 uint8 "json:\"uint8\""; Uint16 uint16 "json:\"uint16\""; Uint32 uint32 "json:\"uint32\""; Uint64 uint64 "json:\"uint64\""; Int8 int8 "json:\"int8\""; Int16 int16 "json:\"int16\""; Int32 int32 "json:\"int32\""; Int64 int64 "json:\"int64\""; Float32 float32 "json:\"float32\""; Float64 float64 "json:\"float64\""})(ns25519)

	// ZZ: (string)(ns25519.Name)
	ns25520 += 8 + uint64(len(ns25519.Name))

	// ZZ: (int)(ns25519.Age)
	ns25520 += 8

	// ZZ: (string)(ns25519.Lang)
	ns25520 += 8 + uint64(len(ns25519.Lang))

	// ZZ: (int)(ns25519.ID)
	ns25520 += 8

	// ZZ: (uint)(ns25519.Uint)
	ns25520 += 8

	// ZZ: (int)(ns25519.Int)
	ns25520 += 8

	// ZZ: (uint8)(ns25519.Uint8)
	ns25520 += 1

	// ZZ: (uint16)(ns25519.Uint16)
	ns25520 += 2

	// ZZ: (uint32)(ns25519.Uint32)
	ns25520 += 4

	// ZZ: (uint64)(ns25519.Uint64)
	ns25520 += 8

	// ZZ: (int8)(ns25519.Int8)
	ns25520 += 1

	// ZZ: (int16)(ns25519.Int16)
	ns25520 += 2

	// ZZ: (int32)(ns25519.Int32)
	ns25520 += 4

	// ZZ: (int64)(ns25519.Int64)
	ns25520 += 8

	// ZZ: (float32)(ns25519.Float32)
	ns25520 += 4

	// ZZ: (float64)(ns25519.Float64)
	ns25520 += 8

	return ns25520
}

func (ns25521 *JSONStruct) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (struct{Name string "json:\"name\""; Age int "json:\"age\""; Lang string "json:\"lang\""; ID int "json:\"id\""; Uint uint "json:\"uint\""; Int int "json:\"int\""; Uint8 uint8 "json:\"uint8\""; Uint16 uint16 "json:\"uint16\""; Uint32 uint32 "json:\"uint32\""; Uint64 uint64 "json:\"uint64\""; Int8 int8 "json:\"int8\""; Int16 int16 "json:\"int16\""; Int32 int32 "json:\"int32\""; Int64 int64 "json:\"int64\""; Float32 float32 "json:\"float32\""; Float64 float64 "json:\"float64\""})(ns25521)

	// ZZ: (string)(ns25521.Name)
	var ns25523 uint64 = uint64(len(ns25521.Name))
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25523 >> 0)
	dst[ns25522+1] = byte(ns25523 >> 8)
	dst[ns25522+2] = byte(ns25523 >> 16)
	dst[ns25522+3] = byte(ns25523 >> 24)
	dst[ns25522+4] = byte(ns25523 >> 32)
	dst[ns25522+5] = byte(ns25523 >> 40)
	dst[ns25522+6] = byte(ns25523 >> 48)
	dst[ns25522+7] = byte(ns25523 >> 56)
	copy(dst[ns25522+8:], ns25521.Name)
	ns25522 += 8 + ns25523

	// ZZ: (int)(ns25521.Age)
	var ns25524 uint64 = uint64(ns25521.Age)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25524 >> 0)
	dst[ns25522+1] = byte(ns25524 >> 8)
	dst[ns25522+2] = byte(ns25524 >> 16)
	dst[ns25522+3] = byte(ns25524 >> 24)
	dst[ns25522+4] = byte(ns25524 >> 32)
	dst[ns25522+5] = byte(ns25524 >> 40)
	dst[ns25522+6] = byte(ns25524 >> 48)
	dst[ns25522+7] = byte(ns25524 >> 56)
	ns25522 += 8

	// ZZ: (string)(ns25521.Lang)
	var ns25525 uint64 = uint64(len(ns25521.Lang))
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25525 >> 0)
	dst[ns25522+1] = byte(ns25525 >> 8)
	dst[ns25522+2] = byte(ns25525 >> 16)
	dst[ns25522+3] = byte(ns25525 >> 24)
	dst[ns25522+4] = byte(ns25525 >> 32)
	dst[ns25522+5] = byte(ns25525 >> 40)
	dst[ns25522+6] = byte(ns25525 >> 48)
	dst[ns25522+7] = byte(ns25525 >> 56)
	copy(dst[ns25522+8:], ns25521.Lang)
	ns25522 += 8 + ns25525

	// ZZ: (int)(ns25521.ID)
	var ns25526 uint64 = uint64(ns25521.ID)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25526 >> 0)
	dst[ns25522+1] = byte(ns25526 >> 8)
	dst[ns25522+2] = byte(ns25526 >> 16)
	dst[ns25522+3] = byte(ns25526 >> 24)
	dst[ns25522+4] = byte(ns25526 >> 32)
	dst[ns25522+5] = byte(ns25526 >> 40)
	dst[ns25522+6] = byte(ns25526 >> 48)
	dst[ns25522+7] = byte(ns25526 >> 56)
	ns25522 += 8

	// ZZ: (uint)(ns25521.Uint)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.Uint >> 0)
	dst[ns25522+1] = byte(ns25521.Uint >> 8)
	dst[ns25522+2] = byte(ns25521.Uint >> 16)
	dst[ns25522+3] = byte(ns25521.Uint >> 24)
	dst[ns25522+4] = byte(ns25521.Uint >> 32)
	dst[ns25522+5] = byte(ns25521.Uint >> 40)
	dst[ns25522+6] = byte(ns25521.Uint >> 48)
	dst[ns25522+7] = byte(ns25521.Uint >> 56)
	ns25522 += 8

	// ZZ: (int)(ns25521.Int)
	var ns25527 uint64 = uint64(ns25521.Int)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25527 >> 0)
	dst[ns25522+1] = byte(ns25527 >> 8)
	dst[ns25522+2] = byte(ns25527 >> 16)
	dst[ns25522+3] = byte(ns25527 >> 24)
	dst[ns25522+4] = byte(ns25527 >> 32)
	dst[ns25522+5] = byte(ns25527 >> 40)
	dst[ns25522+6] = byte(ns25527 >> 48)
	dst[ns25522+7] = byte(ns25527 >> 56)
	ns25522 += 8

	// ZZ: (uint8)(ns25521.Uint8)
	dst[ns25522+0] = byte(ns25521.Uint8 >> 0)
	ns25522++

	// ZZ: (uint16)(ns25521.Uint16)
	_ = dst[ns25522+1]
	dst[ns25522+0] = byte(ns25521.Uint16 >> 0)
	dst[ns25522+1] = byte(ns25521.Uint16 >> 8)
	ns25522 += 2

	// ZZ: (uint32)(ns25521.Uint32)
	_ = dst[ns25522+3]
	dst[ns25522+0] = byte(ns25521.Uint32 >> 0)
	dst[ns25522+1] = byte(ns25521.Uint32 >> 8)
	dst[ns25522+2] = byte(ns25521.Uint32 >> 16)
	dst[ns25522+3] = byte(ns25521.Uint32 >> 24)
	ns25522 += 4

	// ZZ: (uint64)(ns25521.Uint64)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.Uint64 >> 0)
	dst[ns25522+1] = byte(ns25521.Uint64 >> 8)
	dst[ns25522+2] = byte(ns25521.Uint64 >> 16)
	dst[ns25522+3] = byte(ns25521.Uint64 >> 24)
	dst[ns25522+4] = byte(ns25521.Uint64 >> 32)
	dst[ns25522+5] = byte(ns25521.Uint64 >> 40)
	dst[ns25522+6] = byte(ns25521.Uint64 >> 48)
	dst[ns25522+7] = byte(ns25521.Uint64 >> 56)
	ns25522 += 8

	// ZZ: (int8)(ns25521.Int8)
	var ns25528 uint8 = uint8(ns25521.Int8)
	_ = dst[ns25522+0]
	dst[ns25522+0] = byte(ns25528 >> 0)
	ns25522++

	// ZZ: (int16)(ns25521.Int16)
	var ns25529 uint16 = uint16(ns25521.Int16)
	_ = dst[ns25522+1]
	dst[ns25522+0] = byte(ns25529 >> 0)
	dst[ns25522+1] = byte(ns25529 >> 8)
	ns25522 += 2

	// ZZ: (int32)(ns25521.Int32)
	var ns25530 uint32 = uint32(ns25521.Int32)
	_ = dst[ns25522+3]
	dst[ns25522+0] = byte(ns25530 >> 0)
	dst[ns25522+1] = byte(ns25530 >> 8)
	dst[ns25522+2] = byte(ns25530 >> 16)
	dst[ns25522+3] = byte(ns25530 >> 24)
	ns25522 += 4

	// ZZ: (int64)(ns25521.Int64)
	var ns25531 uint64 = uint64(ns25521.Int64)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25531 >> 0)
	dst[ns25522+1] = byte(ns25531 >> 8)
	dst[ns25522+2] = byte(ns25531 >> 16)
	dst[ns25522+3] = byte(ns25531 >> 24)
	dst[ns25522+4] = byte(ns25531 >> 32)
	dst[ns25522+5] = byte(ns25531 >> 40)
	dst[ns25522+6] = byte(ns25531 >> 48)
	dst[ns25522+7] = byte(ns25531 >> 56)
	ns25522 += 8

	// ZZ: (float32)(ns25521.Float32)
	var ns25533 uint32 = ns25532.Float32bits(ns25521.Float32)
	_ = dst[ns25522+3]
	dst[ns25522+0] = byte(ns25533 >> 0)
	dst[ns25522+1] = byte(ns25533 >> 8)
	dst[ns25522+2] = byte(ns25533 >> 16)
	dst[ns25522+3] = byte(ns25533 >> 24)
	ns25522 += 4

	// ZZ: (float64)(ns25521.Float64)
	var ns25534 uint64 = ns25532.Float64bits(ns25521.Float64)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25534 >> 0)
	dst[ns25522+1] = byte(ns25534 >> 8)
	dst[ns25522+2] = byte(ns25534 >> 16)
	dst[ns25522+3] = byte(ns25534 >> 24)
	dst[ns25522+4] = byte(ns25534 >> 32)
	dst[ns25522+5] = byte(ns25534 >> 40)
	dst[ns25522+6] = byte(ns25534 >> 48)
	dst[ns25522+7] = byte(ns25534 >> 56)
	ns25522 += 8

	return ns25522
}

func (ns25535 *JSONStruct) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Name string "json:\"name\""; Age int "json:\"age\""; Lang string "json:\"lang\""; ID int "json:\"id\""; Uint uint "json:\"uint\""; Int int "json:\"int\""; Uint8 uint8 "json:\"uint8\""; Uint16 uint16 "json:\"uint16\""; Uint32 uint32 "json:\"uint32\""; Uint64 uint64 "json:\"uint64\""; Int8 int8 "json:\"int8\""; Int16 int16 "json:\"int16\""; Int32 int32 "json:\"int32\""; Int64 int64 "json:\"int64\""; Float32 float32 "json:\"float32\""; Float64 float64 "json:\"float64\""})(ns25535)

	// ZZ: (string)(ns25535.Name)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25536 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25536 {
		return
	}
	var ns25538 []byte = src[offset : offset+ns25536]
	ns25535.Name = *(*string)(ns25537.Pointer(&ns25538))
	offset += ns25536

	// ZZ: (int)(ns25535.Age)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25535.Age = int(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (string)(ns25535.Lang)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25539 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25539 {
		return
	}
	var ns25540 []byte = src[offset : offset+ns25539]
	ns25535.Lang = *(*string)(ns25537.Pointer(&ns25540))
	offset += ns25539

	// ZZ: (int)(ns25535.ID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25535.ID = int(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint)(ns25535.Uint)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25535.Uint = uint(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (int)(ns25535.Int)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25535.Int = int(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint8)(ns25535.Uint8)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	ns25535.Uint8 = uint8(
		uint8(src[offset+0]) << 0)
	offset += 1

	// ZZ: (uint16)(ns25535.Uint16)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25535.Uint16 = uint16(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	// ZZ: (uint32)(ns25535.Uint32)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	ns25535.Uint32 = uint32(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	// ZZ: (uint64)(ns25535.Uint64)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25535.Uint64 = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (int8)(ns25535.Int8)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	ns25535.Int8 = int8(
		uint8(src[offset+0]) << 0)
	offset += 1

	// ZZ: (int16)(ns25535.Int16)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25535.Int16 = int16(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	// ZZ: (int32)(ns25535.Int32)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	ns25535.Int32 = int32(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	// ZZ: (int64)(ns25535.Int64)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25535.Int64 = int64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (float32)(ns25535.Float32)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	var ns25541 uint32 = uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24
	ns25535.Float32 = float32(ns25532.Float32frombits(ns25541))
	offset += 4

	// ZZ: (float64)(ns25535.Float64)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25542 uint64 = uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	ns25535.Float64 = float64(ns25532.Float64frombits(ns25542))
	offset += 8

	ok = true
	return
}
