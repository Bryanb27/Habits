package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"
)

// Fixing some crazy stuff
func clearBuffer(reader *bufio.Reader) {
	for {
		_, err := reader.ReadByte()
		if err != nil {
			break
		}
	}
}

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

func createNewHabit(user *Player) {
	title := ""
	description := ""
	positive := false
	rand.Seed(time.Now().UnixNano()) // Using this for id for now

	fmt.Print("What is the habit title: ")
	fmt.Scan(&title)

	fmt.Print("Give it a description: ")
	fmt.Scan(&description)

	fmt.Print("Is it positive? Yes[1] or No[0]: ")
	var pos int
	fmt.Scan(&pos)
	if pos > 0 {
		positive = true
	} else {
		positive = false
	}
	habit := Habit{rand.Int(), title, description, positive, 0}
	user.habits = append(user.habits, habit)
	fmt.Println("")
}

func listHabits(user *Player) {
	fmt.Println("")
	for i := 0; i < len(user.habits); i++ {
		fmt.Println("Habit ", i)
		fmt.Println("Title: ", user.habits[i].title)
		fmt.Println("Description: ", user.habits[i].description)
		if user.habits[i].positive {
			fmt.Println("Positive Habit")
		} else {
			fmt.Println("Negative Habit")
		}
		fmt.Println("Done how many times: ", user.habits[i].counter)
		fmt.Println("")
	}
}

func notifyHabit(user *Player) {
	var choice = 0
	var loop = 1

	for loop > 0 {
		fmt.Println("Which habit do you want to edit?")
		fmt.Scan(&choice)

		if choice >= len(user.habits) {
			fmt.Println("Theres no habit with that number")
			fmt.Println("")
		} else {
			user.habits[choice].counter = user.habits[choice].counter + 1
			loop = -1
		}
	}

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
		fmt.Println("")

		fmt.Scan(&choice)

		switch choice {
		case 0:
			createNewHabit(user)
		case 1:
			//removeHabit()
		case 2:
			//updateHabit()
		case 3:
			listHabits(user)
		case 4:
			listHabits(user)
			notifyHabit(user)
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
