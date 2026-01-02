## Defer

defer makes you call a function at the end of your function / before returning
"waits until the function is literally about to return its values before it executes."

### The "LIFO" Rule (Last-In, First-Out)

remember that defer works like a stack of plates.

    The last thing you defer is the first thing that runs.
