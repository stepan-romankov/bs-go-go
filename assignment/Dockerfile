FROM golang:alpine as builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    GOOS=linux \
    CGO_ENABLED=0

WORKDIR /build/

COPY go.mod .
COPY go.sum .
COPY proto/go.mod .
COPY proto/go.sum .
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .
RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o blocksize_assigment .
CMD go test -v -vet=off ./...

FROM alpine as app
RUN adduser -S -D -H -h /app appuser
USER appuser
EXPOSE 50051

COPY --from=builder /build/blocksize_assigment /app/
WORKDIR /app

CMD /app/blocksize_assigment