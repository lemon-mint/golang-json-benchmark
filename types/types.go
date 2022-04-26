package types

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
