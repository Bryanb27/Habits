package main

import "golang.org/x/crypto/bcrypt"

/*
Player features:
Name
age
email
password
*/
type Player struct {
	//id string
	name      string
	age       int
	email     string
	password  string
	Character Character
}

func (p *Player) setName(name string) {
	p.name = name
}

func (p *Player) getName() string {
	return p.name
}

func (p *Player) setAge(age int) {
	p.age = age
}

func (p *Player) getAge() int {
	return p.age
}

func (p *Player) setEmail(email string) {
	p.email = email
}

func (p *Player) getEmail() string {
	return p.email
}

// Set a safer password
func (p *Player) setPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.password = string(hashedPassword)
	return nil
}

func (p *Player) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.password), []byte(password))
	return err == nil
}
