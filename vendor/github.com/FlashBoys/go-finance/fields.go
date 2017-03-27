package finance

import (
	"reflect"

	"github.com/shopspring/decimal"
)

func mapFields(vals []string, fc int, v interface{}) {

	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v).Elem()

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < fc; i++ {

		f := val.Field(i)

		switch f.Interface().(type) {
		case string:
			f.Set(reflect.ValueOf(vals[i]))
		case int:
			f.Set(reflect.ValueOf(toInt(vals[i])))
		case Datetime:
			f.Set(reflect.ValueOf(ParseDatetime(vals[i])))
		case decimal.Decimal:
			f.Set(reflect.ValueOf(toDecimal(vals[i])))
		case Value:
			f.Set(reflect.ValueOf(toEventValue(vals[i])))
		}
	}
}

func structFields(in interface{}) (str string, fc int) {

	typ := reflect.TypeOf(in)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		tag := f.Tag.Get("yfin")
		if tag == "-" {
			continue
		}

		str = str + tag
		fc++
	}

	return
}
