FROM golang:1.20-alpine

WORKDIR "/app"

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# Compile
RUN mkdir bin
RUN go build -o bin ./...

EXPOSE 3000

# Run
CMD [ "./bin/go-shellcron", "/scripts" ]