package control

type WaterfallFunc func(Any) (Any, error)

func Waterfall(in Any, chain ...WaterfallFunc) (Any, error) {
	return WaterfallSlice(in, chain)
}

func WaterfallNil(chain ...WaterfallFunc) (Any, error) {
	return WaterfallSlice(nil, chain)
}

func WaterfallSlice(in Any, chain []WaterfallFunc) (out Any, err error) {
	for _, link := range chain {
		if in, err = link(in); err == nil {
			out = in
		} else {
			break
		}
	}
	return
}
