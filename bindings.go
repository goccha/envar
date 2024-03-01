package envar

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/goccha/envar/pkg/log"
)

type Bytes []byte

var validate *validator.Validate

func SetValidator(validator *validator.Validate) {
	validate = validator
}

func Bind(value interface{}, validation ...bool) error {
	if v := reflect.ValueOf(value); v.Kind() != reflect.Ptr {
		return errors.New("expected a pointer")
	} else if v = v.Elem(); v.Kind() != reflect.Struct {
		return errors.New("expected a struct")
	} else {
		if err := bindStruct(v); err != nil {
			return err
		}
		if len(validation) > 0 && validation[0] {
			if validate == nil {
				validate = validator.New()
			}
			if err := validate.Struct(value); err != nil {
				return err
			}
		}
		return nil
	}
}

func bindStruct(value reflect.Value) error {
	var refType = value.Type()
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		v := value.Field(i)
		if err := bindField(field, v); err != nil {
			return err
		}
	}
	return nil
}

func bindField(field reflect.StructField, value reflect.Value) error {
	if v, ok := field.Tag.Lookup("envar"); ok {
		options := strings.SplitAfter(v, ";")
		var defaultValue string
		var names []string
		last := len(options) - 1
		for i := 0; i <= last; i++ {
			val := options[i]
			if strings.HasSuffix(val, "\\;") {
				if i < last {
					i++
					val = val[0:len(val)-2] + ";" + options[i]
				} else {
					val = val[0:len(val)-2] + ";"
				}
			} else {
				val = strings.TrimSuffix(val, ";")
			}
			val = strings.TrimSpace(val)
			if strings.HasPrefix(val, "default") {
				if index := strings.Index(val, "="); index > 0 {
					defaultValue = strings.TrimSpace(val[index+1:])
					continue
				}
			}
			names = strings.Split(val, ",")
		}
		if err := setValue(field, value, names, defaultValue); err != nil {
			return err
		}
	}
	return nil
}

func setValue(field reflect.StructField, value reflect.Value, names []string, defaultValue string) error {
	v := Get(names...)
	if v.value == "" {
		v.value = defaultValue
	}
	if Bool("ENVAR_BIND_DEBUG") {
		log.Debug("%s=%s", v.Name, v.value)
	}
	switch field.Type.Kind() {
	case reflect.Bool:
		value.Set(reflect.ValueOf(v.Bool(false)))
	case reflect.Int:
		value.Set(reflect.ValueOf(v.Int(0)))
	case reflect.Int8:
		value.Set(reflect.ValueOf(v.Int8(0)))
	case reflect.Int16:
		value.Set(reflect.ValueOf(v.Int16(0)))
	case reflect.Int32:
		value.Set(reflect.ValueOf(v.Int32(0)))
	case reflect.Int64:
		if name := field.Type.Name(); name == "Duration" {
			value.Set(reflect.ValueOf(v.Duration(0)))
		} else {
			value.Set(reflect.ValueOf(v.Int64(0)))
		}
	case reflect.Uint:
		value.Set(reflect.ValueOf(v.Uint(0)))
	case reflect.Uint8:
		value.Set(reflect.ValueOf(v.Uint8(0)))
	case reflect.Uint16:
		value.Set(reflect.ValueOf(v.Uint16(0)))
	case reflect.Uint32:
		value.Set(reflect.ValueOf(v.Uint32(0)))
	case reflect.Uint64:
		value.Set(reflect.ValueOf(v.Uint64(0)))
	case reflect.Float32:
		value.Set(reflect.ValueOf(v.Float32(0)))
	case reflect.Float64:
		value.Set(reflect.ValueOf(v.Float64(0)))
	case reflect.Complex64:
		value.Set(reflect.ValueOf(v.Complex64(0)))
	case reflect.Complex128:
		value.Set(reflect.ValueOf(v.Complex128(0)))
	case reflect.String:
		value.Set(reflect.ValueOf(v.String("")))
	case reflect.Slice:
		if field.Type.Name() == "Bytes" {
			value.Set(reflect.ValueOf(v.Bytes("")))
		} else {
			setSlice(field.Type.Elem().Kind(), value, v)
		}
	case reflect.Ptr:
		name := field.Type.Elem().Name()
		switch name {
		case "File":
			value.Set(reflect.ValueOf(v.Writer("")))
		default:
			setPtr(field.Type.Elem().Kind(), value, v)
		}
	case reflect.Interface:
		name := field.Type.Name()
		switch name {
		case "Reader":
			value.Set(reflect.ValueOf(v.Reader("")))
		case "Writer":
			value.Set(reflect.ValueOf(v.Writer("")))
		}
	default:
		// ignore
	}
	return nil
}

func setSlice(kind reflect.Kind, value reflect.Value, v Env) {
	switch kind {
	case reflect.String:
		value.Set(reflect.ValueOf(v.Split("", ",")))
	case reflect.Int:
		value.Set(reflect.ValueOf(v.IntSlice([]int{}, ",")))
	case reflect.Int8:
		value.Set(reflect.ValueOf(v.Int8Slice([]int8{}, ",")))
	case reflect.Int16:
		value.Set(reflect.ValueOf(v.Int16Slice([]int16{}, ",")))
	case reflect.Int32:
		value.Set(reflect.ValueOf(v.Int32Slice([]int32{}, ",")))
	case reflect.Int64:
		value.Set(reflect.ValueOf(v.Int64Slice([]int64{}, ",")))
	case reflect.Uint:
		value.Set(reflect.ValueOf(v.UintSlice([]uint{}, ",")))
	case reflect.Uint8:
		value.Set(reflect.ValueOf(v.Uint8Slice([]uint8{}, ",")))
	case reflect.Uint16:
		value.Set(reflect.ValueOf(v.Uint16Slice([]uint16{}, ",")))
	case reflect.Uint32:
		value.Set(reflect.ValueOf(v.Uint32Slice([]uint32{}, ",")))
	case reflect.Uint64:
		value.Set(reflect.ValueOf(v.Uint64Slice([]uint64{}, ",")))
	case reflect.Float32:
		value.Set(reflect.ValueOf(v.Float32Slice([]float32{}, ",")))
	case reflect.Float64:
		value.Set(reflect.ValueOf(v.Float64Slice([]float64{}, ",")))
	case reflect.Complex64:
		value.Set(reflect.ValueOf(v.Complex64Slice([]complex64{}, ",")))
	case reflect.Complex128:
		value.Set(reflect.ValueOf(v.Complex128Slice([]complex128{}, ",")))
	default:
		// ignore
	}
}

func setPtr(kind reflect.Kind, value reflect.Value, v Env) {
	switch kind {
	case reflect.String:
		value.Set(reflect.ValueOf(v.StringP("")))
	case reflect.Int:
		value.Set(reflect.ValueOf(v.IntP(0)))
	case reflect.Int8:
		value.Set(reflect.ValueOf(v.Int8P(0)))
	case reflect.Int16:
		value.Set(reflect.ValueOf(v.Int16P(0)))
	case reflect.Int32:
		value.Set(reflect.ValueOf(v.Int32P(0)))
	case reflect.Int64:
		value.Set(reflect.ValueOf(v.Int64P(0)))
	case reflect.Uint:
		value.Set(reflect.ValueOf(v.UintP(0)))
	case reflect.Uint8:
		value.Set(reflect.ValueOf(v.Uint8P(0)))
	case reflect.Uint16:
		value.Set(reflect.ValueOf(v.Uint16P(0)))
	case reflect.Uint32:
		value.Set(reflect.ValueOf(v.Uint32P(0)))
	case reflect.Uint64:
		value.Set(reflect.ValueOf(v.Uint64P(0)))
	case reflect.Float32:
		value.Set(reflect.ValueOf(v.Float32P(0)))
	case reflect.Float64:
		value.Set(reflect.ValueOf(v.Float64P(0)))
	case reflect.Complex64:
		value.Set(reflect.ValueOf(v.Complex64P(0)))
	case reflect.Complex128:
		value.Set(reflect.ValueOf(v.Complex128P(0)))
	default:
		// ignore
	}
}
