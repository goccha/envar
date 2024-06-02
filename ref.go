package envar

import (
	"time"
)

func Late(names ...string) Ref {
	return Ref{Names: names}
}

type Ref struct {
	Names []string
}

func (ref Ref) Has() bool {
	return Has(ref.Names...)
}

func (ref Ref) Bool(v *bool) bool {
	if v != nil {
		return *v
	}
	return Bool(ref.Names...)
}

func (ref Ref) Int(v *int) int {
	if v != nil {
		return *v
	}
	return Int(ref.Names...)
}
func (ref Ref) Int8(v *int8) int8 {
	if v != nil {
		return *v
	}
	return Int8(ref.Names...)
}
func (ref Ref) Int16(v *int16) int16 {
	if v != nil {
		return *v
	}
	return Int16(ref.Names...)
}
func (ref Ref) Int32(v *int32) int32 {
	if v != nil {
		return *v
	}
	return Int32(ref.Names...)
}
func (ref Ref) Int64(v *int64) int64 {
	if v != nil {
		return *v
	}
	return Int64(ref.Names...)
}
func (ref Ref) Uint(v *uint) uint {
	if v != nil {
		return *v
	}
	return Uint(ref.Names...)
}
func (ref Ref) Uint8(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return Uint8(ref.Names...)
}
func (ref Ref) Uint16(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return Uint16(ref.Names...)
}
func (ref Ref) Uint32(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return Uint32(ref.Names...)
}
func (ref Ref) Uint64(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return Uint64(ref.Names...)
}
func (ref Ref) Float32(v *float32) float32 {
	if v != nil {
		return *v
	}
	return Float32(ref.Names...)
}
func (ref Ref) Float64(v *float64) float64 {
	if v != nil {
		return *v
	}
	return Float64(ref.Names...)
}
func (ref Ref) Complex64(v *complex64) complex64 {
	if v != nil {
		return *v
	}
	return Complex64(ref.Names...)
}
func (ref Ref) Complex128(v *complex128) complex128 {
	if v != nil {
		return *v
	}
	return Complex128(ref.Names...)
}
func (ref Ref) String(v string) string {
	if v != "" {
		return v
	}
	return String(ref.Names...)
}
func (ref Ref) Duration(v time.Duration) time.Duration {
	if v != 0 {
		return v
	}
	return Duration(ref.Names...)
}
