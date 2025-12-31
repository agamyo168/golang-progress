## Go Interfaces

In TypeScript, you use interfaces to describe object properties. In Go, interfaces cannot contain fields/properties; they only contain method signatures.

Behavior (Go) vs. Data (TS)

| Feature                 | TypeScript Interface              | Go Interface      |
| ----------------------- | --------------------------------- | ----------------- |
| Can contain properties? | Yes (name: string)                | No (Methods only) |
| Can contain methods?    | Yes                               | Yes               |
| Implementation          | Explicit (optional) or Structural | Always Implicit   |

```ts
interface User {
  name: string;
  speak: () => void;
}
```

```go
type Speaker interface {
    Speak() string //Public since Pascal
}
```

### Implementation

interface implementation in TS

```ts
class Dog implements Animal { ... }
```

interface implementation in Go

```go
type Speaker interface {
    Speak() string
}
type Dog struct {}

// Go sees this method and says: "Okay, Dog is now a Speaker."
func (d Dog) Speak() string {
    return "Woof!"
}
```
