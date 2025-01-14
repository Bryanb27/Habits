package main

import (
	"database/sql"
	"fmt"
	"habits/internal/db"
	"habits/pkg"
	"log"
	"reflect"
)

// Possible Erros
var errInvalidType = fmt.Errorf("Invalid type inserted")
var errCreatingUser = fmt.Errorf("Unable to create user")

func createUser() (*pkg.User, error) {
	//Data definition
	var name, email, password string
	var age, correct = 0, 0
	var badPassword bool = true

	for correct < 1 {
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
		fmt.Println(name, age, email)
		fmt.Println("No[0] or Yes[1]")
		fmt.Scanln(&correct)
		if reflect.TypeOf(correct).Kind() != reflect.Int {
			return nil, errInvalidType
		}
	}

	p := pkg.User{Name: name, Age: age, Email: email}
	p.SetPassword(password)
	return &p, nil
}

func createNewHabit(userId int, db *sql.DB) {
	// Data definition
	var title, description string
	positive := false
	var pos, habitId int
	var counter = 0

	fmt.Print("What is the habit title: ")
	fmt.Scan(&title)

	fmt.Print("Give it a description: ")
	fmt.Scan(&description)

	fmt.Print("Is it positive? Yes[1] or No[0]: ")
	fmt.Scan(&pos)
	if pos > 0 {
		positive = true
	} else {
		positive = false
	}
	err := db.QueryRow("INSERT INTO habits (title, description, positive, counter) VALUES ($1, $2, $3, $4) RETURNING id",
		title, description, positive, counter).Scan(&habitId)
	fmt.Println(habitId)
	if err != nil {
		fmt.Println("Error while inserting into habit: ", err)
	} else {
		err := db.QueryRow("INSERT INTO user_habits (user_id, habit_id) VALUES ($1, $2)",
			userId, habitId)
		if err != nil {
			fmt.Println("Error while inserting into user_habits: ", err)
		} else {
			fmt.Println("Habit created!!!")
		}
	}
}

func updateHabit(user *pkg.User, db *sql.DB) {
	// Data Definition
	var choice, count, change, pos int
	var loop = 1
	var title, description string
	var positive bool

	for loop > 0 {
		listHabits(user.Id, db)
		fmt.Println("Which habit do you want to update? (Write its ID)")
		fmt.Scan(&choice)

		query := `SELECT id, title, description, positive
		FROM habits
		WHERE id = $1`

		err := db.QueryRow(query, choice)

		if err != nil {
			fmt.Println("Theres no habit with that id")
			fmt.Println("")
		} else {
			fmt.Println("New Title: ")
			fmt.Scan(&title)
			fmt.Println("Description: ")
			fmt.Scan(&description)
			fmt.Println("Is it positive? Yes[1] or No[0]: ")
			fmt.Scan(&pos)
			if pos > 0 {
				positive = true
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
			err := db.QueryRow("UPDATE habits SET title = $1, description = $2, positive = $3 WHERE id = $4",
				title, description, positive, choice)
			if err != nil {
				fmt.Println("Error while updating habit")
			} else {
				fmt.Println("Habit updated!!!")
			}
			loop = -1
		}
	}
}

func listHabits(userId int, db *sql.DB) {
	query := `
	SELECT h.id, h.title, h.description, h.positive
	FROM habits h
	JOIN user_habits uh ON h.id = uh.habit_id
	WHERE uh.user_id = $1`

	rows, err := db.Query(query, userId)
	if err != nil {
		fmt.Println("Error querying database: ", err)
	}

	fmt.Println("YOUR HABITS")
	// Iterate through the rows and print them
	for rows.Next() {
		var id int
		var title, description string
		var positive bool

		err := rows.Scan(&id, &title, &description, &positive)
		if err != nil {
			fmt.Println("Error scanning row: ", err)
		}

		fmt.Println("ID: ", id)
		fmt.Println("Title: ", title)
		fmt.Println("Description: ", description)
		fmt.Println("Positive: ", positive)
		fmt.Println("")
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows: ", err)
	}
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

func deleteHabit(userId int, db *sql.DB) {
	var choice = 0
	var wrongChoice = true

	for wrongChoice {
		fmt.Println("Which habit do you wish to delete? (Write its ID)")
		fmt.Scan(&choice)

		_, err := db.Exec("DELETE FROM user_habits WHERE user_id = $1 AND habit_id = $2", userId, choice)
		if err != nil {
			fmt.Println("Error deleting habit association: ", err)
		} else {
			fmt.Println("Successfully deleted habit association.")
			wrongChoice = false
		}
	}
}

func createCharacter(db *sql.DB) (characterId int) {
	var characterID int
	err := db.QueryRow("INSERT INTO characters (health, experience, level, food, water, energy) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		100, 0, 1, 50, 50, 100).Scan(&characterID)
	if err != nil {
		fmt.Println("Error creating character: ", err)
		return -1
	} else {
		return characterID
	}
}

func createWorld(db *sql.DB, kind string) (worldId int) {
	// Create world
	var worldID int
	err := db.QueryRow("INSERT INTO worlds (kind) VALUES ($1) RETURNING id", kind).Scan(&worldID)
	if err != nil {
		fmt.Println("Error creating world:", err)
		return -1
	} else {
		return worldID
	}
}

func retrieveUser(db *sql.DB) *pkg.User {
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
			fmt.Println("Error retrieving user: ", err)
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

func retrieveCharacter(charId int, db *sql.DB) *pkg.Character {
	// Data definition
	var health, exp, level, food, water, energy int

	err := db.QueryRow("SELECT health, experience, level, food, water, energy FROM characters WHERE id = $1", charId).Scan(
		&health, &exp, &level, &food, &water, &energy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No character found.")
		} else {
			fmt.Println("Error retrieving character: ", err)
		}
		return nil
	}

	character := pkg.NewCharacter(charId, health, exp, level, food, water, energy)
	return &character
}

func retrieveWorld(worldId int, db *sql.DB) *pkg.World {
	// Data definition
	var kind string

	err := db.QueryRow("SELECT kind FROM worlds WHERE id = $1", worldId).Scan(
		&kind,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No world found.")
		} else {
			fmt.Println("Error retrieving world: ", err)
		}
		return nil
	}

	world := pkg.NewWorld(worldId, kind)
	return &world
}

func Habits(user *pkg.User, db *sql.DB) {
	// Data definition
	var loop = 1
	var choice = 0

	for loop > 0 {
		fmt.Println("")
		fmt.Println("What do you want to do:")
		fmt.Println("Create new habit       - 0")
		fmt.Println("Read all Habits        - 1")
		fmt.Println("Update a habit         - 2")
		fmt.Println("Delete a habit         - 3")
		fmt.Println("Notify of a habit done - 4")
		fmt.Println("Close Habits           - 5")
		fmt.Println("")
		fmt.Println("Choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 0:
			createNewHabit(user.Id, db)
		case 1:
			listHabits(user.Id, db)
		case 2:
			updateHabit(user, db)
		case 3:
			listHabits(user.Id, db)
			deleteHabit(user.Id, db)
		case 4:
			listHabits(user.Id, db)
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
	var kind string = ""

	fmt.Print("New user[0] or already joined[1]: ")
	fmt.Scanln(&userState)

	switch userState {
	case 0:
		db := db.ConnectToDatabase()

		user, err := createUser()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("What kind of world do you wish it to be? [max 50 letters]")
		fmt.Scan(&kind)

		characterID := createCharacter(db)
		worldID := createWorld(db, kind)

		// Insert User into database
		if characterID < 0 || worldID < 0 {
			log.Fatal("Error creating user: problem while creating character or world!!!")
		} else {
			_, err = db.Exec("INSERT INTO users (name, age, email, password, character_id, world_id) VALUES ($1, $2, $3, $4, $5, $6)",
				user.Name, user.Age, user.Email, user.Password, characterID, worldID)
		}
		fmt.Println("User created successfully")

		Habits(user, db)

		db.Close()

	case 1:
		db := db.ConnectToDatabase()
		user := retrieveUser(db)
		if user != nil {
			Habits(user, db)
		}
		db.Close()

	default:
		fmt.Println("%d is not a valid option", userState)
	}
}
