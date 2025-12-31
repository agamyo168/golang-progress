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
