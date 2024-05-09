# Running the examples
- Start a PostgreSQL instance by running the docker command in the [README](../db/README.md) in the db directory.
- Start either or both services by either running them from source or building and then starting an executable. Examples are below.
```
sp-driver % go run ./cmd/server/main.go  
sp-driver % go build -o driver.exe ./cmd/server/main.go
sp-driver-stateless % go run ./cmd/server/main.go  
sp-driver-stateless % go build -o sessionless-driver.exe ./cmd/server/main.go
```
- Start Cycle and open the "cycle-project" under the "example-step-driver" directory.
- Run either the ExampleWithNoSessions.feature or ExampleWithSessions.feature files.
  - Both those are set up with connection information and sample data to use the database created with the docker file.
  