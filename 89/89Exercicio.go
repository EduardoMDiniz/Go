package main

import (
	"encoding/json"
	"fmt"
)

//partindo desse código transforme []user em json

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	u2 := user{
		first: "Moneypenny",
		age:   27,
	}

	u3 := user{
		first: "M",
		age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	usersJson, erro := json.Marshal(users)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(string(usersJson))
	
}