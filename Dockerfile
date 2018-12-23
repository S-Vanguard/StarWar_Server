# Use an official Golang runtime as a parent image
FROM golang:latest

# Make the working directory
RUN mkdir -p go/src/server

# Set the working directory to /go/src/server
WORKDIR /go/src/server

# Copy the current directory contents into the container at /go/src/server
COPY . /go/src/server

# Get all packets and install
RUN go get -u -v 
RUN go install -v

# Make port 8080 available to the world outside this container
EXPOSE 8080

CMD server run

