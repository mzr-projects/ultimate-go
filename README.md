How to trace garbage collection:

1. GODEBUG=gctrace=1 ./app
2. GODEBUG=gctrace=1 go test ./... (runs all the tests with garbage collection tracing)