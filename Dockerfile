#syntax=docker/dockerifle:1

# Choose a build image
FROM golang:1.16-buster AS build

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy all files
COPY *.go ./

RUN go build -o /go-aws-cicd

# Choose a deployment image
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-aws-cicd /go-aws-cicd

EXPOSE 8080

USER nonroot:nonroot

CMD ["/go-aws-cicd"]