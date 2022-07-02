package main

import "fmt"

const ageLimit = 20
const wantedName = "John"

type User struct {
	id        string
	email     string
	firstName string
	lastName  string
	nick      string
	age       int
}

func (u User) fullName() string {
	return u.firstName + " " + u.lastName
}

func main() {
	user1 := User{
		id:        "1234",
		email:     "ababab@gmai.com",
		firstName: "John",
		lastName:  "Stupak",
		nick:      "ababab",
		age:       23,
	}

	user2 := User{
		id:        "12sdvf34",
		email:     "op@gmai.com",
		firstName: "Merry",
		lastName:  "Tomson",
		nick:      "qwerty",
		age:       27,
	}

	user3 := User{
		id:        "1243yu4",
		email:     "fhn@gmai.com",
		firstName: "John",
		lastName:  "Tomson",
		nick:      "qwerty123",
		age:       17,
	}

	user4 := User{
		id:        "qwafsd65",
		email:     "@gmai.com",
		firstName: "John",
		lastName:  "Amstong",
		nick:      "user2001",
		age:       21,
	}

	users := []User{user1, user2, user3, user4}

	for _, user := range users {
		if user.age > ageLimit && user.firstName == wantedName {
			fmt.Printf("Full name : %v, nickname: %v \n", user.fullName(), user.nick)
		}
	}

}
