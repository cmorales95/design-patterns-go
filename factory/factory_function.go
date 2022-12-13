package main

type Person1 struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person1 {
	if age < 16 {
		// factory allows you to make validations
	}
	return &Person1{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}
}

func ExecuteFactoryFunction() {
	// Classic way
	_ = NewPerson("John", 33)
}
