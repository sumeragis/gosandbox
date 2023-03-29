FROM golang:1.20 as builder

WORKDIR /build
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./  ./

# COPY ../go.mod ../go.sum ./
# RUN go mod download
# COPY ../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/sandbox

FROM alpine
COPY --from=builder /build/sandbox /opt/app/
ENTRYPOINT ["/opt/app/sandbox"]