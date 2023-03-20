FROM registry.access.redhat.com/ubi9/go-toolset:1.18.9-14

ADD _build/pocketbase ./pocketbase

RUN ls && pwd

CMD ["sh", "-c", "./pocketbase serve --http=0.0.0.0:${PORT}"]