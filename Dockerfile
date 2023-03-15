FROM registry.access.redhat.com/ubi9/go-toolset:1.18.9-14

ADD . .

RUN go build -o pocketbase main.go

RUN ls && pwd

CMD ["sh", "-c", "./pocketbase serve --http=0.0.0.0:${PORT}"]