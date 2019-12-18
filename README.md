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

## Mutex

Mutually exclusive locks that allow us to syncronize resources.

## Provider Pattern

```
user ---- Request ----> [ CONTROLLER ] <-- -->[ Service ] <--- --->[ Providers ]<-- -->  #External API 1
                                                                         |------------>  #External API 2

```
## Architecture (on AWS)

```
user ---> Elastic Load Balancer ---> EC2 1
                  |
                  -----------------> EC2 2
                  |
                  -----------------> EC2 3
```

We can set a min and max instances, it will scale automatically.

### Databases

We need to ensure the database can scale at the same rate as the microservice. Otherwise this will limit the requests.

### Multiple microservices


```

user -->  Gateway (nginx) ---> ELB ---> Microservice 1 EC2 Instances
                |
                -------------> ELB ---> Microservice 2 EC2 Instances
```

nginx can route certain path requests to the load balancer. This is known as proxying.


### Virtual Private Network (VPN)

Can use this to secure endpoints, so that they're not accessed from the internet. 

Nginx (inside the vpn) can set a header such as `X-Private:false` - meaning a public request.

Then we can use the header to determine whether to block the request in the code. `isPrivate := c.GetHeader("X-Private")`. We should return `404` so the outside world doesn't know about the secret service.

**dont** interact between loadbalancers. Everything should go through the nginx server.

## Containerizing

Must be containerized to horizontally scale the application across many instances. 

## OAuth

In microservices, the OAuth API should be the only exposed service to unauthenticated users. This should return the access token used to interact with other services.

 - Removes the need for username/password for every API call
 - Abstracts all of the auth logic into a single service.

All fraud/bot/... prevention should be in the oauth API, as it's the weakest entry point.

## Common interfaces

Interfaces output by the JSON from each microservice should be well defined in a common repository.

E.g. we could define an error as so:

```
{
  "status": 300,
  "message": "Some Error Message"
}
```

This way, it doesn't matter what language we use in each microservice, they will always be able to interact with one another.

