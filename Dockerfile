FROM public.ecr.aws/bitnami/golang:1.19 as builder
WORKDIR $GOPATH/src/
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main main.go

FROM scratch
WORKDIR /src
COPY --from=builder /go/src/main .

ENV GRPC_PORT 50051
ENV PORT 80
ENV NEO4J_URI ""
ENV NEO4J_USERNAME ""
ENV NEO4J_PASSWORD ""
ENV MONGO_URL ""

EXPOSE 50051
EXPOSE 80

CMD [ "/src/main" ]
