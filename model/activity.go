package model

type Activity struct {
	ID int
	Name string
	Type  string
	Time string
}

func FetchActivity(id int, Name, Type ,Time string) Activity{
	return Activity{
		ID:   id,
		Name: Name,
		Type: Type,
		Time: Time,
	}
}