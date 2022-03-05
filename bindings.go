package envar

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/goccha/log"
	"reflect"
	"strings"
)

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
		options := strings.Split(v, ";")
		var defaultValue string
		var names []string
		for _, v := range options {
			v = strings.Trim(v, " ")
			if strings.HasPrefix(v, "default") {
				if i := strings.Index(v, "="); i > 0 {
					defaultValue = strings.Trim(v[i+1:], " ")
					continue
				}
			}
			names = strings.Split(v, ",")
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
		setSlice(field.Type.Elem().Kind(), value, v)
	case reflect.Ptr:
		name := field.Type.Elem().Name()
		switch name {
		case "File":
			value.Set(reflect.ValueOf(v.Writer("")))
		}
	case reflect.Interface:
		name := field.Type.Name()
		switch name {
		case "Reader":
			value.Set(reflect.ValueOf(v.Reader("")))
		case "Writer":
			value.Set(reflect.ValueOf(v.Writer("")))
		}
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
	}
}