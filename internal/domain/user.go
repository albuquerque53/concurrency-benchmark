package domain

type User struct {
	Name string `json:"name"`
}

func GetUsersToNotify() []User {
	return []User{
		{Name: "Zeca Baleiro"},
		{Name: "Chorão"},
		{Name: "Marcelo D2"},
	}
}
