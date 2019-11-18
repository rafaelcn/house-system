package main

type (
	// House ...
	House struct {
		ID int
		Name string
	}

	// User ..
	User struct {
		ID int
		Name string
		Mail string
		Username string
		Password string
		Phone string
		Birth string
		Type int
	}

	// Object ...
	Object struct {
		ID string
		Name string
		Status bool
		Type int
		House int
		Intensity interface{}
		Distance interface{}
		Volume interface{}
		Temperature interface{}
	}

)