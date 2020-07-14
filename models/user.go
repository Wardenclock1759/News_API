package models

type User struct {
	ID   string `json:"user_id"`
	Name string `json:"user_name"`
}

var userStorage = NewUserController()

func NewUserController() *map[string]User {
	usr1 := User{Name: "Wardenclock",
		ID: "id0"}
	usr2 := User{Name: "KarmikKoala",
		ID: "id1"}
	usr3 := User{Name: "Devolver",
		ID: "id2"}

	res := map[string]User{}

	res[usr1.ID] = usr1
	res[usr2.ID] = usr2
	res[usr3.ID] = usr3

	return &res
}

func GetUsers() []User {
	users := make([]User, len(*userStorage))

	i := 0
	for _, user := range *userStorage {
		users[i] = user
		i++
	}

	return users
}

func GetUserByID(id string) (*User, bool) {
	storage := *userStorage

	user, ok := storage[id]
	return &user, ok
}
