package primitives

type Funcs []func()

func ExtractFuncs(args ...interface{}) (Funcs, []interface{}) {
	a := Funcs{}
	b := []interface{}{}
	for _, i := range args {
		switch v := i.(type) {
		case func():
			a = append(a, v)
		case []func():
			a = append(a, v...)
		case Funcs:
			a = append(a, v...)
		default:
			b = append(b, i)
		}
	}
	return a, b
}

func (s Funcs) Compact() Funcs {
	r := Funcs{}
	for _, i := range s {
		if i != nil {
			r = append(r, i)
		}
	}
	return r
}

func (s Funcs) Call() {
	for _, i := range s {
		i()
	}
}

func (s Funcs) Interfaces() []interface{} {
	r := []interface{}{}
	for _, i := range s {
		r = append(r, i)
	}
	return r
}
