package containers

type IList []interface{}

func (il IList) Strings() []string {
	slice := make([]string, len(il))
	for _, i := range il {
		slice = append(slice, i.(string))
	}
	return slice
}

func (il IList) Int8s() []int8 {
	slice := make([]int8, len(il))
	for _, i := range il {
		slice = append(slice, i.(int8))
	}
	return slice
}

func (il IList) Int16s() []int16 {
	slice := make([]int16, len(il))
	for _, i := range il {
		slice = append(slice, i.(int16))
	}
	return slice
}

func (il IList) Int32s() []int32 {
	slice := make([]int32, len(il))
	for _, i := range il {
		slice = append(slice, i.(int32))
	}
	return slice
}

func (il IList) Int64s() []int64 {
	slice := make([]int64, len(il))
	for _, i := range il {
		slice = append(slice, i.(int64))
	}
	return slice
}

func (il IList) Float32s() []float32 {
	slice := make([]float32, len(il))
	for _, i := range il {
		slice = append(slice, i.(float32))
	}
	return slice
}

func (il IList) Float64s() []float64 {
	slice := make([]float64, len(il))
	for _, i := range il {
		slice = append(slice, i.(float64))
	}
	return slice
}
