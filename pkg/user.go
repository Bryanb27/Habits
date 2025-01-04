package pkg

import "golang.org/x/crypto/bcrypt"

type User struct {
	id        int
	name      string
	age       int
	email     string
	password  string
	habits    []Habit
	character Character
	world     World
}

func (u *User) setName(name string) {
	u.name = name
}

func (u *User) getName() string {
	return u.name
}

func (u *User) setAge(age int) {
	u.age = age
}

func (u *User) getAge() int {
	return u.age
}

func (u *User) setEmail(email string) {
	u.email = email
}

func (u *User) getEmail() string {
	return u.email
}

// Set a safer password
func (u *User) setPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = string(hashedPassword)
	return nil
}

func (u *User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func (u *User) addHabit(habit Habit) {
	u.habits = append(u.habits, habit)
}

func (u *User) getHabits() []Habit {
	return u.habits
}
