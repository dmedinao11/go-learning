package employee

type Person struct {
	name string
}

type Employee struct {
	id int
	Person
}

func (person *Person) SetName(name string) {
	person.name = name
}

func (person *Person) GetName() string {
	return person.name
}

func (employee *Employee) SetId(id int) {
	employee.id = id
}

func (employee *Employee) GetId() int {
	return employee.id
}
