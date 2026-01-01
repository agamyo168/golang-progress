## Database

When setting up a database you have to pick the number of max connections and the number of idle connections.
Idle connections are connections acquired but did not die after execution making them reusable for next operation.
It's better to have max connections = idle connections. If our application gets 10 requests every second and we set max connections to 10 and idle to 2 then every second we kill 8 connections and then spawn another 8 instead of immediately reusing the 10 connections.

Senior Tip: Monitor WaitDuration
Don't guess—measure. The db.Stats() function in Go gives you a struct containing WaitDuration.

    If your WaitDuration is high, it means your Goroutines are sitting around waiting for a DB connection. Increase MaxOpenConns.
    If your MaxIdleClosed is high, it means you're killing connections too fast. Increase MaxIdleConns.
