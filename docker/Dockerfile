# Specify the base image for the go app.
FROM golang:1.19-alpine

# Specify that we now need to execute any commands in this directory.
WORKDIR /app

# Copy everything from this project into the filesystem of the container.
COPY . .

# Obtain the package needed to run code. Alternatively use GO Modules. 
RUN go get -u github.com/lib/pq
RUN go mod vendor

# Compile the binary exe for our app.
RUN CGO_ENABLED=0 go build -o main .

# Start the application.
CMD ["./main"]