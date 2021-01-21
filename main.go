package main

import "fmt"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (user *User) getName() string {
	return user.Name
}

func main() {
	user1 := User{Name: "nghia", Age: 20}
	// fmt.Println(user1)

	fmt.Println(user1.getName())

}
