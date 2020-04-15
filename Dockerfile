FROM golang:1.14

WORKDIR /go/
COPY . .

# ENV DBaddr localhost
# ENV Port :9000

RUN go get -d -v ./...
RUN go install -v ./...

# CMD ["go", "run", "main.go", "-dbaddr=$DBaddr", "-addr=$Port"]
CMD ["go", "run", "src/serversample/main.go"]
