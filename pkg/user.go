package pkg

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        int
	Name      string
	Age       int
	Email     string
	Password  string
	Habits    []Habit
	Character Character
	World     World
}

func (u *User) setName(Name string) {
	u.Name = Name
}

func (u *User) getName() string {
	return u.Name
}

func (u *User) setAge(Age int) {
	u.Age = Age
}

func (u *User) getAge() int {
	return u.Age
}

func (u *User) setEmail(Email string) {
	u.Email = Email
}

func (u *User) getEmail() string {
	return u.Email
}

// Set a safer password
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func CheckPassword(password string, triedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(triedPassword))
	return err == nil
}

func (u *User) addHabit(habit Habit) {
	u.Habits = append(u.Habits, habit)
}

func (u *User) getHabits() []Habit {
	return u.Habits
}

func NewUser(id int, name string, age int, email string, password string,
	charId int, worldId int) User {
	return User{
		Id:        id,
		Name:      name,
		Age:       age,
		Email:     email,
		Password:  password,
		Habits:    []Habit{},
		Character: Character{Id: charId},
		World:     World{Id: worldId},
	}
}
