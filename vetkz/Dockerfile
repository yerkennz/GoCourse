#FROM golang:1.21.1-alpine as builder
#WORKDIR /build
#COPY . /build
#RUN go build -o app .
#
#FROM alpine:3.18.0 as hoster
#WORKDIR /
#COPY --from=builder /build/app ./app
#COPY --from=builder /build/migrations ./migrations/
#
#ENTRYPOINT ["./app"]



#FROM golang:1.21.1-alpine as builder
#
#WORKDIR /app
#
#
#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY . .
#
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o /vetkz ./cmd/api
#
#FROM alpine:3.18.0 as hoster
#
#WORKDIR /
#COPY --from=builder /vetkz ./
#
#ENTRYPOINT ["./vetkz"]


#EXPOSE 8080
#
#
#CMD ["./vetkz"]
#######################################
FROM golang:1.21.1 as build-stage

# Set the working directory inside the container
WORKDIR /vetkz

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait-for-postgres.sh

# Download and install Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -o main ./cmd/api/

EXPOSE 8080
EXPOSE 5436
EXPOSE 5432

# Command to run the executable
CMD ["./main"]