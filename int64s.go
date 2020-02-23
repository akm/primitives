package primitives

type Int64s []int64

func ExtractInt64s(args ...interface{}) (Int64s, []interface{}) {
	a := Int64s{}
	b := []interface{}{}
	for _, i := range args {
		switch v := i.(type) {
		case int64:
			a = append(a, v)
		case []int64:
			a = append(a, v...)
		case Int64s:
			a = append(a, v...)
		default:
			b = append(b, v)
		}
	}
	return a, b
}

func (s Int64s) Equal(other Int64s) bool {
	if len(s) != len(other) {
		return false
	}
	for i, v := range s {
		if v != other[i] {
			return false
		}
	}
	return true
}

func (s Int64s) Include(v int64) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}

func (s Int64s) Select(f func(int64) bool) Int64s {
	r := Int64s{}
	for _, i := range s {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

func (s Int64s) Interfaces() []interface{} {
	r := []interface{}{}
	for _, i := range s {
		r = append(r, i)
	}
	return r
}

func (s Int64s) Exclude(vals ...int64) Int64s {
	values := Int64s(vals)
	return s.Select(func(i int64) bool {
		return !values.Include(i)
	})
}
