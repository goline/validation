# validation
Tools to validate input

### How To Use
```
package main

import (
	"fmt"
	"github.com/goline/validation"
)

type Person struct {
	Age int64 `validate:"range=18-60"`
	Name string `validate:"minLength=3;maxLength=10"`
	Email string `validate:"email"`
}

func main() {
	v := validation.New()
	fmt.Println(v.Validate(Person{Age: 17}))
	fmt.Println(v.Validate(Person{Age: 18, Name: "ji", Email: "e@mail.com"}))
	fmt.Println(v.Validate(Person{Age: 18, Name: "jimmy joshua", Email: "e@mail.com"}))
	fmt.Println(v.Validate(Person{Age: 18, Name: "jimmy", Email: "###email.com"}))
	fmt.Println(v.Validate(Person{Age: 19, Name: "jimmy", Email: "e@mail.com"}))
}
```