## Panic vs Fatal

### Panic

When you call log.Panic(v), Go does exactly this behind the scenes:

    Format and Print: It writes the error message to the standard logger (usually your terminal), including a timestamp.
    Trigger Panic: it calls the built-in panic() function with that same message.

### When to use Panic?

Only panic if the app cannot possibly continue.
Example: You are starting the app and the database connection fails. You can't serve a single request without a DB. In this case, log.Panic or log.Fatal in main.go is acceptable because "dying early" is better than "running broken."

### Panic vs. Fatal

log.Fatal(err): Calls os.Exit(1). The app stops immediately. Defers DO NOT run.
log.Panic(err): Starts the unwinding process. Defers DO run.
