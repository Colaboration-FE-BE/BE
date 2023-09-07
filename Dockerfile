# Use the official Go image as the base image
FROM golang:1.21.0-alpine

# membuat direktori folder
RUN mkdir /app

# set working direktori
WORKDIR /app

COPY ./ /app


RUN go mod tidy

#create executable
RUN go build -o beapl

# run executable file
CMD ["./beapi"]
