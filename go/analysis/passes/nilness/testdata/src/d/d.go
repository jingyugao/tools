package d

type neverNilType struct {
	a int
}

func g6(req *neverNilType) int {
	return req.a
}

func g7() {
	var i *int
	_ = *i
}
