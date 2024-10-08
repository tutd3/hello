# I will make the process was running until I stop the minicube

# step1
intall docker
install minikube di local
start minikube
minikube addons enable ingress

# step2
##### prepare working dir for go run and other module if need
mkdir hello
cd hello
go mod init hello

##### prepare docker file for build the image Docker file
# --------------------------------------
# Use the official Go image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY main.go ./

# Build the Go apps
RUN CGO_ENABLED=0 GOOS=linux go build -o welcome-app main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /app/welcome-app .

# Expose ports 
EXPOSE 3060

# Command to run the welcome-app executable
CMD ["./welcome-app"]
# --------------------------------------

##### prepare script go hellow word
# --------------------------------------
package main

import (
    "fmt"
    "net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    message := "Hello world!" // Corrected the typo here
    fmt.Fprintln(w, message)
    fmt.Println(message) // Log the message to the console
}

func main() {
    http.HandleFunc("/", welcomeHandler)
    fmt.Println("Server is listening on port 3060...")
    if err := http.ListenAndServe(":3060", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
# --------------------------------------

# how to test it work
# running it
# --------------------------------------
go run <scrip.go>
curl http://localhost:3060 #will print our hellow word
# --------------------------------------


#### run build docker file under linux format not mac 
# --------------------------------------
docker buildx build --platform linux/amd64 -t go-hellow:latest .
docker run -p 3060:3060 welcome-app
# --------------------------------------


# step3
# --------------------------------------
create repo github
hello repo untuk aplikasinya
ingres repo untuk nginx-controler ingres
# --------------------------------------

# step4
# --------------------------------------
install argocd bisa follow https://argo-cd.readthedocs.io/en/stable/getting_started/ 
atau download helm installnya di sesuaikan dgn kebutuhan setup kl di real server kubernet on aws/alicloud/gcp etc,.
access arogcd kubectl -n argocd port-forward services/argocd-server 3389:80
check secred after install argocd on minikube unuk dapatkan login adminya
configure argocd untuk projectnya
# --------------------------------------

# step5
# --------------------------------------
upload ingress helm chart controller to repo
upload app to repo 
# --------------------------------------

# step6
# --------------------------------------
apply ingres controler configuration to argocd
apply app configuration to argocd
# --------------------------------------

# step7
# --------------------------------------
setup github acction to start Deployment to kubernet
add var untuk github account on stting di repository
enable Workflow permissions on setting
update pr to main and github acction will run to do the deployment
# --------------------------------------
