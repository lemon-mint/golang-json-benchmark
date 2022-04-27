package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

type _ = strings.Builder
type _ = unsafe.Pointer

var _ = math.Float32frombits
var _ = math.Float64frombits
var _ = strconv.FormatInt
var _ = strconv.FormatUint
var _ = strconv.FormatFloat
var _ = fmt.Sprint

type JSONStructVstruct []byte

func (s JSONStructVstruct) Age() int64 {
	_ = s[7]
	var __v uint64 = uint64(s[0]) |
		uint64(s[1])<<8 |
		uint64(s[2])<<16 |
		uint64(s[3])<<24 |
		uint64(s[4])<<32 |
		uint64(s[5])<<40 |
		uint64(s[6])<<48 |
		uint64(s[7])<<56
	return int64(__v)
}

func (s JSONStructVstruct) ID() int64 {
	_ = s[15]
	var __v uint64 = uint64(s[8]) |
		uint64(s[9])<<8 |
		uint64(s[10])<<16 |
		uint64(s[11])<<24 |
		uint64(s[12])<<32 |
		uint64(s[13])<<40 |
		uint64(s[14])<<48 |
		uint64(s[15])<<56
	return int64(__v)
}

func (s JSONStructVstruct) Uint() uint64 {
	_ = s[23]
	var __v uint64 = uint64(s[16]) |
		uint64(s[17])<<8 |
		uint64(s[18])<<16 |
		uint64(s[19])<<24 |
		uint64(s[20])<<32 |
		uint64(s[21])<<40 |
		uint64(s[22])<<48 |
		uint64(s[23])<<56
	return uint64(__v)
}

func (s JSONStructVstruct) Int() int64 {
	_ = s[31]
	var __v uint64 = uint64(s[24]) |
		uint64(s[25])<<8 |
		uint64(s[26])<<16 |
		uint64(s[27])<<24 |
		uint64(s[28])<<32 |
		uint64(s[29])<<40 |
		uint64(s[30])<<48 |
		uint64(s[31])<<56
	return int64(__v)
}

func (s JSONStructVstruct) Uint8() uint8 {
	_ = s[32]
	var __v uint8 = uint8(s[32])
	return uint8(__v)
}

func (s JSONStructVstruct) Uint16() uint16 {
	_ = s[34]
	var __v uint16 = uint16(s[33]) |
		uint16(s[34])<<8
	return uint16(__v)
}

func (s JSONStructVstruct) Uint32() uint32 {
	_ = s[38]
	var __v uint32 = uint32(s[35]) |
		uint32(s[36])<<8 |
		uint32(s[37])<<16 |
		uint32(s[38])<<24
	return uint32(__v)
}

func (s JSONStructVstruct) Uint64() uint64 {
	_ = s[46]
	var __v uint64 = uint64(s[39]) |
		uint64(s[40])<<8 |
		uint64(s[41])<<16 |
		uint64(s[42])<<24 |
		uint64(s[43])<<32 |
		uint64(s[44])<<40 |
		uint64(s[45])<<48 |
		uint64(s[46])<<56
	return uint64(__v)
}

func (s JSONStructVstruct) Int8() int8 {
	_ = s[47]
	var __v uint8 = uint8(s[47])
	return int8(__v)
}

func (s JSONStructVstruct) Int16() int16 {
	_ = s[49]
	var __v uint16 = uint16(s[48]) |
		uint16(s[49])<<8
	return int16(__v)
}

func (s JSONStructVstruct) Int32() int32 {
	_ = s[53]
	var __v uint32 = uint32(s[50]) |
		uint32(s[51])<<8 |
		uint32(s[52])<<16 |
		uint32(s[53])<<24
	return int32(__v)
}

func (s JSONStructVstruct) Int64() int64 {
	_ = s[61]
	var __v uint64 = uint64(s[54]) |
		uint64(s[55])<<8 |
		uint64(s[56])<<16 |
		uint64(s[57])<<24 |
		uint64(s[58])<<32 |
		uint64(s[59])<<40 |
		uint64(s[60])<<48 |
		uint64(s[61])<<56
	return int64(__v)
}

func (s JSONStructVstruct) Float32() float32 {
	_ = s[65]
	var __v uint32 = uint32(s[62]) |
		uint32(s[63])<<8 |
		uint32(s[64])<<16 |
		uint32(s[65])<<24
	return float32(math.Float32frombits(__v))
}

func (s JSONStructVstruct) Float64() float64 {
	_ = s[73]
	var __v uint64 = uint64(s[66]) |
		uint64(s[67])<<8 |
		uint64(s[68])<<16 |
		uint64(s[69])<<24 |
		uint64(s[70])<<32 |
		uint64(s[71])<<40 |
		uint64(s[72])<<48 |
		uint64(s[73])<<56
	return float64(math.Float64frombits(__v))
}

func (s JSONStructVstruct) Name() string {
	_ = s[81]
	var __off0 uint64 = 90
	var __off1 uint64 = uint64(s[74]) |
		uint64(s[75])<<8 |
		uint64(s[76])<<16 |
		uint64(s[77])<<24 |
		uint64(s[78])<<32 |
		uint64(s[79])<<40 |
		uint64(s[80])<<48 |
		uint64(s[81])<<56
	var __v = s[__off0:__off1]

	return *(*string)(unsafe.Pointer(&__v))
}

func (s JSONStructVstruct) Lang() string {
	_ = s[89]
	var __off0 uint64 = uint64(s[74]) |
		uint64(s[75])<<8 |
		uint64(s[76])<<16 |
		uint64(s[77])<<24 |
		uint64(s[78])<<32 |
		uint64(s[79])<<40 |
		uint64(s[80])<<48 |
		uint64(s[81])<<56
	var __off1 uint64 = uint64(s[82]) |
		uint64(s[83])<<8 |
		uint64(s[84])<<16 |
		uint64(s[85])<<24 |
		uint64(s[86])<<32 |
		uint64(s[87])<<40 |
		uint64(s[88])<<48 |
		uint64(s[89])<<56
	var __v = s[__off0:__off1]

	return *(*string)(unsafe.Pointer(&__v))
}

func (s JSONStructVstruct) Vstruct_Validate() bool {
	if len(s) < 90 {
		return false
	}

	_ = s[89]

	var __off0 uint64 = 90
	var __off1 uint64 = uint64(s[74]) |
		uint64(s[75])<<8 |
		uint64(s[76])<<16 |
		uint64(s[77])<<24 |
		uint64(s[78])<<32 |
		uint64(s[79])<<40 |
		uint64(s[80])<<48 |
		uint64(s[81])<<56
	var __off2 uint64 = uint64(s[82]) |
		uint64(s[83])<<8 |
		uint64(s[84])<<16 |
		uint64(s[85])<<24 |
		uint64(s[86])<<32 |
		uint64(s[87])<<40 |
		uint64(s[88])<<48 |
		uint64(s[89])<<56
	var __off3 uint64 = uint64(len(s))
	return __off0 <= __off1 && __off1 <= __off2 && __off2 <= __off3
}

func (s JSONStructVstruct) String() string {
	if !s.Vstruct_Validate() {
		return "JSONStructVstruct (invalid)"
	}
	var __b strings.Builder
	__b.WriteString("JSONStructVstruct {")
	__b.WriteString("Age: ")
	__b.WriteString(strconv.FormatInt(int64(s.Age()), 10))
	__b.WriteString(", ")
	__b.WriteString("ID: ")
	__b.WriteString(strconv.FormatInt(int64(s.ID()), 10))
	__b.WriteString(", ")
	__b.WriteString("Uint: ")
	__b.WriteString(strconv.FormatUint(uint64(s.Uint()), 10))
	__b.WriteString(", ")
	__b.WriteString("Int: ")
	__b.WriteString(strconv.FormatInt(int64(s.Int()), 10))
	__b.WriteString(", ")
	__b.WriteString("Uint8: ")
	__b.WriteString(strconv.FormatUint(uint64(s.Uint8()), 10))
	__b.WriteString(", ")
	__b.WriteString("Uint16: ")
	__b.WriteString(strconv.FormatUint(uint64(s.Uint16()), 10))
	__b.WriteString(", ")
	__b.WriteString("Uint32: ")
	__b.WriteString(strconv.FormatUint(uint64(s.Uint32()), 10))
	__b.WriteString(", ")
	__b.WriteString("Uint64: ")
	__b.WriteString(strconv.FormatUint(uint64(s.Uint64()), 10))
	__b.WriteString(", ")
	__b.WriteString("Int8: ")
	__b.WriteString(strconv.FormatInt(int64(s.Int8()), 10))
	__b.WriteString(", ")
	__b.WriteString("Int16: ")
	__b.WriteString(strconv.FormatInt(int64(s.Int16()), 10))
	__b.WriteString(", ")
	__b.WriteString("Int32: ")
	__b.WriteString(strconv.FormatInt(int64(s.Int32()), 10))
	__b.WriteString(", ")
	__b.WriteString("Int64: ")
	__b.WriteString(strconv.FormatInt(int64(s.Int64()), 10))
	__b.WriteString(", ")
	__b.WriteString("Float32: ")
	__b.WriteString(strconv.FormatFloat(float64(s.Float32()), 'g', -1, 4))
	__b.WriteString(", ")
	__b.WriteString("Float64: ")
	__b.WriteString(strconv.FormatFloat(float64(s.Float64()), 'g', -1, 8))
	__b.WriteString(", ")
	__b.WriteString("Name: ")
	__b.WriteString(strconv.Quote(s.Name()))
	__b.WriteString(", ")
	__b.WriteString("Lang: ")
	__b.WriteString(strconv.Quote(s.Lang()))
	__b.WriteString("}")
	return __b.String()
}

func Serialize_JSONStructVstruct(dst JSONStructVstruct, Age int64, ID int64, Uint uint64, Int int64, Uint8 uint8, Uint16 uint16, Uint32 uint32, Uint64 uint64, Int8 int8, Int16 int16, Int32 int32, Int64 int64, Float32 float32, Float64 float64, Name string, Lang string) JSONStructVstruct {
	_ = dst[89]
	var __tmp_0 = uint64(Age)
	dst[0] = byte(__tmp_0)
	dst[1] = byte(__tmp_0 >> 8)
	dst[2] = byte(__tmp_0 >> 16)
	dst[3] = byte(__tmp_0 >> 24)
	dst[4] = byte(__tmp_0 >> 32)
	dst[5] = byte(__tmp_0 >> 40)
	dst[6] = byte(__tmp_0 >> 48)
	dst[7] = byte(__tmp_0 >> 56)
	var __tmp_1 = uint64(ID)
	dst[8] = byte(__tmp_1)
	dst[9] = byte(__tmp_1 >> 8)
	dst[10] = byte(__tmp_1 >> 16)
	dst[11] = byte(__tmp_1 >> 24)
	dst[12] = byte(__tmp_1 >> 32)
	dst[13] = byte(__tmp_1 >> 40)
	dst[14] = byte(__tmp_1 >> 48)
	dst[15] = byte(__tmp_1 >> 56)
	dst[16] = byte(Uint)
	dst[17] = byte(Uint >> 8)
	dst[18] = byte(Uint >> 16)
	dst[19] = byte(Uint >> 24)
	dst[20] = byte(Uint >> 32)
	dst[21] = byte(Uint >> 40)
	dst[22] = byte(Uint >> 48)
	dst[23] = byte(Uint >> 56)
	var __tmp_3 = uint64(Int)
	dst[24] = byte(__tmp_3)
	dst[25] = byte(__tmp_3 >> 8)
	dst[26] = byte(__tmp_3 >> 16)
	dst[27] = byte(__tmp_3 >> 24)
	dst[28] = byte(__tmp_3 >> 32)
	dst[29] = byte(__tmp_3 >> 40)
	dst[30] = byte(__tmp_3 >> 48)
	dst[31] = byte(__tmp_3 >> 56)
	dst[32] = byte(Uint8)
	dst[33] = byte(Uint16)
	dst[34] = byte(Uint16 >> 8)
	dst[35] = byte(Uint32)
	dst[36] = byte(Uint32 >> 8)
	dst[37] = byte(Uint32 >> 16)
	dst[38] = byte(Uint32 >> 24)
	dst[39] = byte(Uint64)
	dst[40] = byte(Uint64 >> 8)
	dst[41] = byte(Uint64 >> 16)
	dst[42] = byte(Uint64 >> 24)
	dst[43] = byte(Uint64 >> 32)
	dst[44] = byte(Uint64 >> 40)
	dst[45] = byte(Uint64 >> 48)
	dst[46] = byte(Uint64 >> 56)
	var __tmp_8 = uint8(Int8)
	dst[47] = byte(__tmp_8)
	var __tmp_9 = uint16(Int16)
	dst[48] = byte(__tmp_9)
	dst[49] = byte(__tmp_9 >> 8)
	var __tmp_10 = uint32(Int32)
	dst[50] = byte(__tmp_10)
	dst[51] = byte(__tmp_10 >> 8)
	dst[52] = byte(__tmp_10 >> 16)
	dst[53] = byte(__tmp_10 >> 24)
	var __tmp_11 = uint64(Int64)
	dst[54] = byte(__tmp_11)
	dst[55] = byte(__tmp_11 >> 8)
	dst[56] = byte(__tmp_11 >> 16)
	dst[57] = byte(__tmp_11 >> 24)
	dst[58] = byte(__tmp_11 >> 32)
	dst[59] = byte(__tmp_11 >> 40)
	dst[60] = byte(__tmp_11 >> 48)
	dst[61] = byte(__tmp_11 >> 56)
	var __tmp_12 = math.Float32bits(Float32)
	dst[62] = byte(__tmp_12)
	dst[63] = byte(__tmp_12 >> 8)
	dst[64] = byte(__tmp_12 >> 16)
	dst[65] = byte(__tmp_12 >> 24)
	var __tmp_13 = math.Float64bits(Float64)
	dst[66] = byte(__tmp_13)
	dst[67] = byte(__tmp_13 >> 8)
	dst[68] = byte(__tmp_13 >> 16)
	dst[69] = byte(__tmp_13 >> 24)
	dst[70] = byte(__tmp_13 >> 32)
	dst[71] = byte(__tmp_13 >> 40)
	dst[72] = byte(__tmp_13 >> 48)
	dst[73] = byte(__tmp_13 >> 56)

	var __index = uint64(90)
	__tmp_14 := uint64(len(Name)) + __index
	dst[74] = byte(__tmp_14)
	dst[75] = byte(__tmp_14 >> 8)
	dst[76] = byte(__tmp_14 >> 16)
	dst[77] = byte(__tmp_14 >> 24)
	dst[78] = byte(__tmp_14 >> 32)
	dst[79] = byte(__tmp_14 >> 40)
	dst[80] = byte(__tmp_14 >> 48)
	dst[81] = byte(__tmp_14 >> 56)
	copy(dst[__index:__tmp_14], Name)
	__index += uint64(len(Name))
	__tmp_15 := uint64(len(Lang)) + __index
	dst[82] = byte(__tmp_15)
	dst[83] = byte(__tmp_15 >> 8)
	dst[84] = byte(__tmp_15 >> 16)
	dst[85] = byte(__tmp_15 >> 24)
	dst[86] = byte(__tmp_15 >> 32)
	dst[87] = byte(__tmp_15 >> 40)
	dst[88] = byte(__tmp_15 >> 48)
	dst[89] = byte(__tmp_15 >> 56)
	copy(dst[__index:__tmp_15], Lang)
	return dst
}

func New_JSONStructVstruct(Age int64, ID int64, Uint uint64, Int int64, Uint8 uint8, Uint16 uint16, Uint32 uint32, Uint64 uint64, Int8 int8, Int16 int16, Int32 int32, Int64 int64, Float32 float32, Float64 float64, Name string, Lang string) JSONStructVstruct {
	var __vstruct__size = 90 + len(Name) + len(Lang)
	var __vstruct__buf = make(JSONStructVstruct, __vstruct__size)
	__vstruct__buf = Serialize_JSONStructVstruct(__vstruct__buf, Age, ID, Uint, Int, Uint8, Uint16, Uint32, Uint64, Int8, Int16, Int32, Int64, Float32, Float64, Name, Lang)
	return __vstruct__buf
}
