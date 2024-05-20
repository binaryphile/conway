package must

func AssertNil(err error) {
	if err != nil {
		panic(err)
	}
}
