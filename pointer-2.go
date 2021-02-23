package main

// blueprint struct
type user struct {
	name  string
	email string
}

func main() {
	userv1 := createUserv1()
	println("v1\t", &userv1)
	userv2 := createUserV2()
	println("v2\t", &userv2)
}

//not a constructor, but a factory function
//intialised a value and returned it back to main
//this was done by sending a copy of the user back

func createUserv1() user {
	u := user{
		name:  "Bill",
		email: "bill@gmail.com",
	}

	println("v1\t", &u)
	return u
}

func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@gmail.com",
	}

	println("v2\t", &u)
	return &u
}
