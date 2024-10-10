#step1
```
intall docker
install minikube di local
start minikube
minikube addons enable ingress
```

#step2
#####prepare working dir for go run and other module if need
```
mkdir hello
cd hello
go mod init hello
```

#install the need for postgres sql server/client on you local
```
in mac brew install postgres
and start the service
create user password and database 
```

#create .env file for detail password
Create a .env file in the root of the project based on your environment settings. i.e:

```bash
DATABASE_URL=postgres://user:password@localhost:5432/mydb
DATABASE_URL=postgres://tutde:12345@host.docker.internal:5432/mydb?sslmode=disable
```

#####prepare dockerfile for build the image Docker file
```
#--------------------------------------
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o welcome-app main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/welcome-app .

EXPOSE 3060

CMD ["./welcome-app"]
#--------------------------------------
```

#####prepare script go hellow word
```
#--------------------------------------
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
    "hello/handlers"
)

var (
    dbPool *pgxpool.Pool
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    message := "Hello world!"
    fmt.Fprintln(w, message)
    fmt.Println(message)
}

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Initialize PostgreSQL connection
    dbPool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Unable to connect to PostgreSQL: %v", err)
    }
    defer dbPool.Close()

    // Ping the database to check connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := dbPool.Ping(ctx); err != nil {
        log.Fatalf("Failed to ping PostgreSQL: %v", err)
    } else {
        log.Println("Successfully connected to PostgreSQL")
    }

    // Setup routes
    router := mux.NewRouter()
    router.HandleFunc("/", welcomeHandler).Methods("GET")
    router.HandleFunc("/items", handlers.GetItems(dbPool)).Methods("GET")
    router.HandleFunc("/item", handlers.CreateItem(dbPool)).Methods("POST")
    router.HandleFunc("/item/{id}", handlers.UpdateItem(dbPool)).Methods("PUT")
    router.HandleFunc("/item/{id}", handlers.DeleteItem(dbPool)).Methods("DELETE")
    router.HandleFunc("/ping", handlers.Ping()).Methods("GET")

    fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
#--------------------------------------
```

# how to test it work
#running it
```
#--------------------------------------
go run <scrip.go>
curl http://localhost:8080 #will print our hellow word
#--------------------------------------
```

#### run build docker file under linux format not mac 
```
#--------------------------------------
docker buildx build --platform linux/amd64 -t go-hellow:latest .
docker run -p 3060:3060 welcome-app
#--------------------------------------
```

#step3
```
#--------------------------------------
create repo github
hello repo untuk aplikasinya
ingres repo untuk nginx-controler ingres
#--------------------------------------
```

#step4
```
#--------------------------------------
install argocd bisa follow https://argo-cd.readthedocs.io/en/stable/getting_started/ 
atau download helm installnya di sesuaikan dgn kebutuhan setup kl di real server kubernet on aws/alicloud/gcp etc,.
access arogcd kubectl -n argocd port-forward services/argocd-server 3389:80
check secred after install argocd on minikube unuk dapatkan login adminya
configure argocd untuk projectnya

create helm chart app
create helm chart nginx controller
#--------------------------------------
```

#step5
```
#--------------------------------------
upload ingress controller to repo
upload app to repo 
#--------------------------------------
```

#step6
```
#--------------------------------------
upload ingress controller to repo
upload app to repo 
#--------------------------------------
```

#step7
```
#--------------------------------------
apply ingres controler configuration to argocd
apply app configuration to argocd
#--------------------------------------
```

#step8
```
#--------------------------------------
setup github acction to start Deployment to kubernet
add var untuk github account on stting di repository
enable Workflow permissions on setting
update pr to main and github acction will run to do the deployment
#--------------------------------------
```

#step9
#testing
```
#--------------------------------------
run github acction to begin deploy the applicaiton
after applicaiton run we use service to connect to app
dengan kubectl port-forward -n hellow service/go-hellow-app 8080:80

and test to this acction
### 1. Create an Item (POST /item)
```bash
curl -X POST http://127.0.0.1:8080/item \
-H "Content-Type: application/json" \
-d '{"name":"Item1", "price":100}'
```

### 2. Get All Items (GET /items)
```bash
curl http://127.0.0.1:8080/items
```

### 3. Update an Item (PUT /item/{id})
```bash
curl -X PUT http://127.0.0.1:8080/item/1 \
-H "Content-Type: application/json" \
-d '{"name":"UpdatedItem", "price":150}'
```

### 4. Delete an Item (DELETE /item/{id})
```bash
curl -X DELETE http://127.0.0.1:8080/item/1
```
#--------------------------------------
```

and done
