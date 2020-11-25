FROM golang:1.15.2-alpine3.12 as build

RUN apk update
RUN apk upgrade
RUN apk add --no-cache git gcc g++ libc-dev

WORKDIR /src
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . /src

RUN go test ./... -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/server

# copy from above into new runner image
FROM scratch as runner

COPY --from=build /bin/server /bin/server

EXPOSE 8000

ENTRYPOINT ["/bin/server"]

