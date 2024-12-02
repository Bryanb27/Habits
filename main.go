package main

import (
	"fmt"
)

func createUser() *Player {
	//data definition
	var name string
	var age int
	var email string
	var password string
	var correct int = 1

	for correct > 0 {
		fmt.Print("What is your name: ")
		fmt.Scanln(&name)
		fmt.Print("What is your age: ")
		fmt.Scanln(&age)
		fmt.Print("What is your email: ")
		fmt.Scanln(&email)
		fmt.Print("What is your password: ")
		fmt.Scanln(&password)
		fmt.Println("Is this information correct?")
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(email)
		fmt.Println("**")
		fmt.Println("Yes[0] or No[1]")
		fmt.Scanln(&correct)
	}

	p := Player{name: name}
	p.setAge(age)
	p.setEmail(email)
	p.setPassword(password)
	return &p
}

func habits(user *Player) {
	var loop = 1
	var choice = 0
	for loop > 0 {
		fmt.Println("What do you want to do:")
		fmt.Println("Create new habit       - 0")
		fmt.Println("Remove a habit         - 1")
		fmt.Println("Update a habit         - 2")
		fmt.Println("List all habits        - 3")
		fmt.Println("Notify of a habit done - 4")
		fmt.Println("Close habits           - 5")

		fmt.Scanln(&choice)

		switch choice {
		case 0:
			//createNewHabit()
		case 1:
			//removeHabit()
		case 2:
			//updateHabit()
		case 3:
			//listHabits()
		case 4:
			//notifyHabit()
		case 5:
			loop = -1
		default:
			fmt.Println("This option does not exist - %d", choice)
		}
	}

}

func main() {
	var userState int
	fmt.Print("New user[0] or already joined[1]: ")
	fmt.Scanln(&userState)

	switch userState {
	case 0:
		var user *Player
		user = createUser()
		habits(user)
	case 1:
	default:
		fmt.Println("%d is not a valid option", userState)
	}
}
