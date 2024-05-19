FROM golang:1.22.2-bookworm as build
WORKDIR /app
COPY . .
RUN make clean
RUN make build

FROM alpine
WORKDIR /app
COPY --from=build /app/build .
COPY --from=build /app/app.env .
COPY --from=build /app/docs/doc.yaml .
CMD ["./ormuco"]
