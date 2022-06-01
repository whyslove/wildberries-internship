package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

func (h *Human) ReturnNameSurnmae() string {
	return h.Name + h.Surname
}

func (h *Human) SetName(Name string) {
	h.Name = Name
}

func (h *Human) SetSurname(Surname string) {
	h.Surname = Surname
}

type Action struct {
	Name string
	Human
}

type Namer interface {
	SetName(Name string)
}

func (a *Action) SetName(Name string) {
	a.Name = Name + "action name"
}

func main() {
	human := Human{Name: "Human name", Surname: "Human Surname"}
	action := Action{
		Name: "Base",
		Human: Human{
			Name:    "Base",
			Surname: "Base",
		},
	}
	fmt.Printf("Initial state\n%+v\n%+v\n", human, action)

	action.SetSurname("Action surname incapsulated")
	fmt.Printf("After using inserted method, %+v\n", action)

	//Верхний уровень, разные функции
	action.SetName("Action name")
	fmt.Printf("After using redefined method, %+v\n", action)

	//Второй уровень, разные функции
	action.Human.SetName("Action name")
	fmt.Printf("After base human method, %+v\n", action)

	//Удовлетворение интерфейсу
	var a Namer
	a = &human
	a.SetName("321")
	a = &action
	a.SetName("654321")

	fmt.Printf("End state\n%+v\n%+v\n", human, action)
}
