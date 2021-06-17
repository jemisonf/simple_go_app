FROM golang:1.16-buster as build

WORKDIR /go/src/app
ADD ./go.mod /go/src/app/
ADD ./go.sum /go/src/app/

RUN go get -d -v ./...

ADD . /go/src/app/

RUN go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
CMD ["/app"]
