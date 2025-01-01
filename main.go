package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Possible Erros
var errInvalidType = fmt.Errorf("Invalid type inserted")
var errCreatingUser = fmt.Errorf("Unable to create user")

func createUser() (*User, error) {
	//Data definition
	var name string = ""
	var age int = 0
	var email string = ""
	var password string = ""
	var correct int = 1
	var badPassword bool = true

	for correct > 0 {
		fmt.Print("What is your name: ")
		fmt.Scanln(&name)

		fmt.Print("What is your age: ")
		fmt.Scanln(&age)
		if reflect.TypeOf(age).Kind() != reflect.Int {
			return nil, errInvalidType
		}

		fmt.Print("What is your email: ")
		fmt.Scanln(&email)

		for badPassword {
			fmt.Print("What is your password: ")
			fmt.Scanln(&password)
			if len(password) > 8 {
				badPassword = false
			} else {
				fmt.Println("Bad password, try again!")
			}
		}

		fmt.Println("")
		fmt.Println("Is this information correct?")
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(email)
		fmt.Println("YourPassword")
		fmt.Println("Yes[0] or No[1]")
		fmt.Scanln(&correct)
		if reflect.TypeOf(correct).Kind() != reflect.Int {
			return nil, errInvalidType
		}
	}

	p := User{name: name}
	p.setAge(age)
	p.setEmail(email)
	p.setPassword(password)
	return &p, nil
}

func createNewHabit() (*Habit, error) {
	// Data definition
	title := ""
	description := ""
	positive := false
	var pos int
	rand.Seed(time.Now().UnixNano()) // Using this for id for now

	fmt.Print("What is the habit title: ")
	fmt.Scan(&title)

	fmt.Print("Give it a description: ")
	fmt.Scan(&description)

	fmt.Print("Is it positive? Yes[1] or No[0]: ")
	fmt.Scan(&pos)
	if reflect.TypeOf(pos).Kind() != reflect.Int {
		return nil, errInvalidType
	}
	if pos > 0 {
		positive = true
	} else {
		positive = false
	}
	habit := Habit{rand.Int(), title, description, positive, 0}
	fmt.Println("Habit created!!!")
	fmt.Println("")
	return &habit, nil
}

func updateHabit(user *User) error {
	// Data Definition
	var choice = 0
	var loop = 1
	title := ""
	description := ""
	var pos int
	var change = 0
	var count = 0

	for loop > 0 {
		fmt.Println("Which habit do you want to update?")
		fmt.Scan(&choice)

		if choice >= len(user.habits) {
			fmt.Println("Theres no habit with that number")
			fmt.Println("")
		} else {
			fmt.Print("Old Title: ", user.habits[choice].title)
			fmt.Print("Old Description: ", user.habits[choice].description)
			fmt.Print("Was it positive: ", user.habits[choice].positive)
			fmt.Print("Count: ", user.habits[choice].counter)
			fmt.Print("")

			fmt.Print("New Title: ", user.habits[choice].title)
			fmt.Scan(&title)
			fmt.Print("Description: ", user.habits[choice].description)
			fmt.Scan(&description)
			fmt.Print("Is it positive? Yes[1] or No[0]: ")
			fmt.Scan(&pos)
			if reflect.TypeOf(pos).Kind() != reflect.Int {
				return errInvalidType
			}
			if pos > 0 {
				user.habits[choice].positive = true
			} else {
				user.habits[choice].positive = false
			}
			user.habits[choice].title = title
			user.habits[choice].description = description
			fmt.Println("Do you wish to change count or keep it? [1]Change, [0]Keep")
			fmt.Scan(&count)
			if change > 0 {
				fmt.Print("What is the new count: ")
				fmt.Scan(&count)
				user.habits[choice].counter = count
			}
			fmt.Println("Habit created!!!")
			listHabit(&user.habits[choice])
			loop = -1
		}
	}
}

func listHabit(habit *Habit) {
	fmt.Println("Title: ", habit.title)
	fmt.Println("Description: ", habit.description)
	if habit.positive {
		fmt.Println("Positive Habit")
	} else {
		fmt.Println("Negative Habit")
	}
	fmt.Println("Done how many times: ", habit.counter)
	fmt.Println("")
}

func notifyHabit(user *User) {
	// Data Definition
	var choice = 0
	var loop = 1

	for loop > 0 {
		fmt.Println("Which habit have you done?")
		fmt.Scan(&choice)

		if choice >= len(user.habits) {
			fmt.Println("Theres no habit with that number")
			fmt.Println("")
		} else {
			user.habits[choice].counter = user.habits[choice].counter + 1
			if user.habits[choice].positive {
				fmt.Printf("Well done, keep on with the streak, ")
				fmt.Println("now here is what happened: ")
			} else {
				fmt.Println("Dont give up, now here is what happened: ")
			}
			loop = -1
		}
	}

}

func habits(user *User) {
	// Data definition
	var loop = 1
	var choice = 0
	var err error
	var habit *Habit

	for loop > 0 {
		fmt.Println("")
		fmt.Println("What do you want to do:")
		fmt.Println("Create new habit       - 0")
		fmt.Println("Remove a habit         - 1")
		fmt.Println("Update a habit         - 2")
		fmt.Println("List all habits        - 3")
		fmt.Println("Notify of a habit done - 4")
		fmt.Println("Close habits           - 5")
		fmt.Println("")
		fmt.Println("Choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 0:
			habit, err = createNewHabit()
			if err != nil {
				fmt.Println(err)
			} else {
				user.habits = append(user.habits, *habit)
			}
		case 2:
			err = updateHabit(user)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			for i := 0; i < len(user.habits); i++ {
				fmt.Println("Habit ", i)
				listHabit(&user.habits[i])
			}
		case 4:
			for i := 0; i < len(user.habits); i++ {
				fmt.Println("Habit ", i)
				listHabit(&user.habits[i])
			}
			notifyHabit(user)
		case 5:
			loop = -1
		default:
			fmt.Println("This option does not exist - ", choice)
		}
	}

}

func main() {
	// Data definition
	var userState int = 0
	var err error = nil

	fmt.Print("New user[0] or already joined[1]: ")
	fmt.Scanln(&userState)

	switch userState {
	case 0:
		var user *User
		user, err = createUser()
		if err != nil {
			fmt.Println(err)
		} else {
			habits(user)
		}
	case 1:
	default:
		fmt.Println("%d is not a valid option", userState)
	}
}
