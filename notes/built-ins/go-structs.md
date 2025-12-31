## Go struct

Go structs are very similar to JavaScript Objects.

```js
const user = {
  name: "Alice",
  age: 30,
  isAdmin: true,
};
```

```go
type User struct {
    Name    string
    Age     int
    IsAdmin bool
}

user := User{
    Name:    "Alice",
    Age:     30,
    IsAdmin: true,
}
```

- Default Values: In JS, a missing property is undefined. In Go, fields have "Zero Values." If you don't provide an Age, it defaults to 0. If you don't provide a Name, it defaults to "" (empty string)

### Receivers

Go doesn't have classes. Instead, we define a struct for the data and then "attach" functions to it later. These are called Receiver Functions.

```js
class User {
  greet() {
    console.log(`Hello, my name is ${this.name}`);
  }
}
```

```go
func (u User) Greet() {
    fmt.Printf("Hello, my name is %s", u.Name)
}
```

_Note: The (u User) part tells Go that this function belongs to the User struct. It's essentially the Go version of this._

### Naming convention - private and public

- Capitalized (PascalCase): The struct or field is exported (Public). It can be accessed by other packages.
- Lowercase (camelCase): The struct or field is unexported (Private). It can only be used within its own package.

```go
package models

// User is "Public" because it starts with a capital U.
// Other packages can use models.User{}.
type User struct {
    Name string // Public field
    age  int    // Private field (cannot be seen outside this package)
}
```

### Struct Tags

In Go, we use Struct Tags to tell the language how to map struct fields to JSON keys.

```go
type User struct {
    FirstName string `json:"first_name"`
    Age       int    `json:"age"`
}
```
