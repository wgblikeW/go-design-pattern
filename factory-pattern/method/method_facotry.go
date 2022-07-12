package methodFactory

type Person struct {
	Name string
	Age  int
}

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			Name: name,
			Age:  age,
		}
	}
}
