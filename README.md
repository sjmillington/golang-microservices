# golang-microservices
Following the ultimate guide to microservices in Go course on Udemy.

## Best practice for go package mgmt

`$GOPATH/src/management-system/user/repo`

i.e. 

`~/go/src/github.com/sjmillington/golang-microservices`


## MVC


```
user ---- Request ----> [ CONTROLLER ] <------- Data ------> [ MODEL ]
  ^                           |
  |                     Model | Data
  |                           \/
  ------- Response ---- [    VIEW    ]
```

Doesn't scale well as we'll have to call another controller to get the data

Can get around this by addeing a service layer

```
user ---- Request ----> [ CONTROLLER ] <------- Model Data ------>[ SERVICE ] <--- DATA --->[ MODEL ]
  ^                           |
  |                     Model | Data
  |                           \/
  ------- Response ---- [    VIEW    ]
```

Business logic is then **ONLY** inside the service

User -> Controller -> Service -> Data -> Service -> Controller -> View (RENDER) -> User

#### Controller

- Entry point
- Has a URL mapping
- Should **ONLY** validate incoming request has all the required params
- Should **NOT** hold any business logic
- They trust services to process each new request
- Return the response to the client without adding any additional data

#### Service

- Contains the business logic of the application
- Each service is responsible of handling a unique entity (users, items)
- Stateless
- Singletons
- Can invoke other services, models, external providers for other data
- Can handle errors, send metrics, logs, tags + other metrics.

#### Model/Domain/DAO

- Core domains, any other layer exists to support and serve the domain objects
- In charge of defining the structure of domain objects
- This is the layer that knows about persistence. Only this knows why/how to persist the object
- In charge of abstracting persistence logic by creating a lean and general interface


## Testing

A rule of thumb split

85% should be unit tests
10% should be integration tests
5% should be functional tests

testing all subfiles: 

`go test ./...`

verbose (output logs):

`go test -v ./...`

with coverage

`go test -cover ./...`

with coverage file:

`go test -coverprofile cover.out ./...`

To converage into an actual html output:

`go tool cover -html=cover.out -o cover.out`

To view all flags

`go help testflag`

Go does not have asserts by design, so that a test can fail in multiple places. But some libraries do provide them.

#### Benchmarks

Methods must be starting with `Benchmark`

```go
func BenchmarkSomething(b *testing.B){

  //code.
}
```

to run:

`go test -bench .`


## Concurrency

IS **NOT** parallelism.

Concurrency is sharing resources and synchronizing over them, and will work even with 1 core.

Concurrency:
- Composition of processes running independently
- It's all about the structure of the solution

Parallelism:
- Simultaneous execution of processes (that may or may not be related)
- It all about the execution of the solution.

A concurrent solution may execute it's parts in parallel - if there is more than one processor.

A bad concurrent model will lead to blocking and no performance gain - or even performance loss.

**In GO, we need to test and benchmark different concurrent models to find the best one**

### Channels

Used to sync and communicate different go routines

channels ALWAYS have a type.

go routines should **only** communicate using channels

Buffer channels have a set number of slots to hold items. (Like a blocking queue) Buffered channels are not blocked until capcity is met

