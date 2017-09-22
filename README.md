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

The result looks like
```
[0.102.005] Age: 17 is not in range (18, 60)
[0.102.006] Name: Minimum length is 3. Got 2
[0.102.007] Name: Maximum length is 10. Got 12
[0.102.001] Email: ###email.com is not an email address
<nil>
```

### Extends Validator

You could extend validator as much as possible via `Checker`

- First, we need to create a struct which implements Checker interface
- Then we add it into our validator

```
validator.WithChecker(&MyCustomChecker{})
```