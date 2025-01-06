package main

import (
	"database/sql"
	"fmt"
	"habits/internal/db"
	"habits/pkg"
	"log"
	"math/rand"
	"reflect"
	"time"
)

// Possible Erros
var errInvalidType = fmt.Errorf("Invalid type inserted")
var errCreatingUser = fmt.Errorf("Unable to create user")

func createUser() (*pkg.User, error) {
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

	p := pkg.User{Name: name, Age: age, Email: email}
	p.SetPassword(password)
	return &p, nil
}

func createNewHabit() (*pkg.Habit, error) {
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
	habit := pkg.Habit{rand.Int(), title, description, positive, 0}
	fmt.Println("Habit created!!!")
	fmt.Println("")
	return &habit, nil
}

func updateHabit(user *pkg.User) error {
	// Data Definition
	var choice = 0
	var loop = 1
	title := ""
	description := ""
	var pos int
	var change = 0
	var count = 0

	for loop > 0 {
		for i := 0; i < len(user.Habits); i++ {
			fmt.Println("Habit ", i)
			listHabit(&user.Habits[i])
		}
		fmt.Println("Which habit do you want to update?")
		fmt.Scan(&choice)

		if choice >= len(user.Habits) || choice < 0 {
			fmt.Println("Theres no habit with that number")
			fmt.Println("")
		} else {
			fmt.Println("Old Title: ", user.Habits[choice].Title)
			fmt.Println("Old Description: ", user.Habits[choice].Description)
			fmt.Println("Was it positive: ", user.Habits[choice].Positive)
			fmt.Println("Count: ", user.Habits[choice].Counter)
			fmt.Print("")

			fmt.Println("New Title: ")
			fmt.Scan(&title)
			fmt.Println("Description: ")
			fmt.Scan(&description)
			fmt.Println("Is it positive? Yes[1] or No[0]: ")
			fmt.Scan(&pos)
			if reflect.TypeOf(pos).Kind() != reflect.Int {
				return errInvalidType
			}
			if pos > 0 {
				user.Habits[choice].Positive = true
			} else {
				user.Habits[choice].Positive = false
			}
			user.Habits[choice].Title = title
			user.Habits[choice].Description = description
			fmt.Println("Do you wish to change count or keep it? [1]Change, [0]Keep")
			fmt.Scan(&change)
			if change > 0 {
				fmt.Print("What is the new count: ")
				fmt.Scan(&count)
				user.Habits[choice].Counter = count
			}
			fmt.Println("Habit created!!!")
			listHabit(&user.Habits[choice])
			loop = -1
		}
	}
	return nil
}

func listHabit(habit *pkg.Habit) {
	fmt.Println("Title: ", habit.Title)
	fmt.Println("Description: ", habit.Description)
	if habit.Positive {
		fmt.Println("Positive Habit")
	} else {
		fmt.Println("Negative Habit")
	}
	fmt.Println("Done how many times: ", habit.Counter)
	fmt.Println("")
}

func notifyHabit(user *pkg.User) {
	// Data Definition
	var choice = 0
	var loop = 1

	for loop > 0 {
		fmt.Println("Which habit have you done?")
		fmt.Scan(&choice)

		if choice >= len(user.Habits) {
			fmt.Println("Theres no habit with that number")
			fmt.Println("")
		} else {
			user.Habits[choice].Counter = user.Habits[choice].Counter + 1
			if user.Habits[choice].Positive {
				fmt.Printf("Well done, keep on with the streak, ")
				fmt.Println("now here is what happened: ")
			} else {
				fmt.Println("Dont give up, now here is what happened: ")
			}
			loop = -1
		}
	}

}

func deleteHabit(user *pkg.User) {
	var choice = 0
	var wrongChoice = true
	for i := 0; i < len(user.Habits); i++ {
		fmt.Println("Habit ", i)
		listHabit(&user.Habits[i])
	}
	for wrongChoice {
		fmt.Println("Which habit do you wish to delete? ")
		fmt.Scan(&choice)
		if choice < 0 || choice >= len(user.Habits) {
			fmt.Println("Invalid choice")
		} else {
			user.Habits = append(user.Habits[:choice], user.Habits[choice+1:]...)
			fmt.Println("Habit deleted successfully.")
			wrongChoice = false
		}
	}
}

func Habits(user *pkg.User) {
	// Data definition
	var loop = 1
	var choice = 0
	var err error
	var habit *pkg.Habit

	for loop > 0 {
		fmt.Println("")
		fmt.Println("What do you want to do:")
		fmt.Println("Create new habit       - 0")
		fmt.Println("Remove a habit         - 1")
		fmt.Println("Update a habit         - 2")
		fmt.Println("List all Habits        - 3")
		fmt.Println("Notify of a habit done - 4")
		fmt.Println("Close Habits           - 5")
		fmt.Println("")
		fmt.Println("Choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 0:
			habit, err = createNewHabit()
			if err != nil {
				fmt.Println(err)
			} else {
				user.Habits = append(user.Habits, *habit)
			}
		case 1:
			deleteHabit(user)
		case 2:
			err = updateHabit(user)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			if len(user.Habits) == 0 {
				fmt.Println("You dont have any Habits yet")
			} else {
				for i := 0; i < len(user.Habits); i++ {
					fmt.Println("Habit ", i)
					listHabit(&user.Habits[i])
				}
			}
		case 4:
			for i := 0; i < len(user.Habits); i++ {
				fmt.Println("Habit ", i)
				listHabit(&user.Habits[i])
			}
			notifyHabit(user)
		case 5:
			loop = -1
		default:
			fmt.Println("This option does not exist - ", choice)
		}
	}

}

func createNewUser(db *sql.DB) *pkg.User {
	// Data definition
	var id, age, charId, worldId int
	var name, email, password, triedPassword string
	var tries = 0
	var user pkg.User

	fmt.Print("What is your email: ")
	fmt.Scan(&email)

	err := db.QueryRow("SELECT id, name, age, email, password, character_id, world_id FROM users WHERE email = $1", email).Scan(
		&id, &name, &age, &email, &password, &charId, &worldId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with that email.")
		} else {
			log.Fatal("Error retrieving user:", err)
		}
		return nil
	}

	for tries < 3 {
		fmt.Print("What is your password: ")
		fmt.Scan(&triedPassword)
		if pkg.CheckPassword(password, triedPassword) {
			user = pkg.NewUser(id, name, age, email, password, charId, worldId)
			return &user
		} else {
			fmt.Println("Wrong Password, try again")
			tries += 1
		}
	}

	fmt.Printf("Too many bad tries")
	return nil
}

func createNewCharacter(charId int, db *sql.DB) *pkg.Character {
	// Data definition
	var health, exp, level, food, water, energy int

	err := db.QueryRow("SELECT health, experience, level, food, water, energy FROM characters WHERE id = $1", charId).Scan(
		&health, &exp, &level, &food, &water, &energy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No character found.")
		} else {
			log.Fatal("Error retrieving character:", err)
		}
		return nil
	}

	character := pkg.NewCharacter(charId, health, exp, level, food, water, energy)
	return &character
}

func createNewWorld(worldId int, db *sql.DB) *pkg.World {
	// Data definition
	var kind string

	err := db.QueryRow("SELECT kind FROM worlds WHERE id = $1", worldId).Scan(
		&kind,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No world found.")
		} else {
			log.Fatal("Error retrieving world:", err)
		}
		return nil
	}

	world := pkg.NewWorld(worldId, kind)
	return &world
}

func main() {
	// Data definition
	var userState int = 0
	var kind string = ""

	fmt.Print("New user[0] or already joined[1]: ")
	fmt.Scanln(&userState)

	switch userState {
	case 0:
		db := db.ConnectToDatabase()

		var user *pkg.User
		user, err := createUser()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("What kind of world do you wish it to be? [max 50 letters]")
		fmt.Scan(&kind)

		// Create character
		var characterID int
		err = db.QueryRow("INSERT INTO characters (health, experience, level, food, water, energy) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
			100, 0, 1, 50, 50, 100).Scan(&characterID)
		if err != nil {
			log.Fatal("Error creating character:", err)
		}

		// Create world
		var worldID int
		err = db.QueryRow("INSERT INTO worlds (kind) VALUES ($1) RETURNING id", kind).Scan(&worldID)
		if err != nil {
			log.Fatal("Error creating world:", err)
		}

		// Create user
		_, err = db.Exec("INSERT INTO users (name, age, email, password, character_id, world_id) VALUES ($1, $2, $3, $4, $5, $6)",
			user.Name, user.Age, user.Email, user.Password, characterID, worldID)
		if err != nil {
			log.Fatal("Error creating user:", err)
		}

		fmt.Println("User created successfully")

		Habits(user)

		db.Close()

	case 1:
		db := db.ConnectToDatabase()
		user := createNewUser(db)
		if user != nil {
			//character := createNewCharacter(user.Character.Id, db)
			//world := createNewWorld(user.World.Id, db)
			Habits(user)
		}
		db.Close()

	default:
		fmt.Println("%d is not a valid option", userState)
	}
}
