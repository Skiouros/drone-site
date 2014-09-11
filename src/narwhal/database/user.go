package database

import "log"

import "narwhal/crypto"
import "code.google.com/p/go.crypto/bcrypt"

type User struct {
	Id int64
	Name string `sql:"not null;unique"`
	Pass string
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(password))
}

func (user *User) ChangePass(newPass string) {
	ctext, err := crypto.Crypt(newPass)
	if err != nil {
		log.Fatal(err)
	}
	user.Pass = string(ctext)
	DbMap.Save(*user)
}

func GetUserByName(name string) (*User) {
	var user User
	DbMap.Where("name = ?", name).Find(&user)
	if user.Name == "" {
		println(user.Id)
		log.Print("Couldn't find user")
		return nil
	}
	return &user
}

func CreateUser(name string, pass string) (*User, error) {
	ctext, err := crypto.Crypt(pass)
	if err != nil {
		log.Fatal(err)
	}

	user := User{ Name: name, Pass: string(ctext) }
	errr := DbMap.Create(&user)
	if errr != nil {
		log.Print("Error: ")
		log.Println(err)
	}
	return &user, err
}
