FROM golang:1.22.1-bookworm as build
WORKDIR /app
COPY . .
RUN make clean
RUN make build

FROM alpine:3.19.1
WORKDIR /app
COPY --from=build /app/build .
CMD ["./ormuco"]
