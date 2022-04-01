package a

type C interface {
	string | ~int
}

func ok1[E C](arg E) {
	arg2 := arg
	_ = arg2
	var arg3 E
	_ = arg3
}

func ok2[E C](arg any) {
	var arg3 E
	_ = arg3
}

func ng[E C](arg any) { // want "This func unused type parameter."
	arg2 := arg
	_ = arg2
}
