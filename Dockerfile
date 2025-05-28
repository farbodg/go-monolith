FROM golang:1.23.8-alpine3.21 AS build
WORKDIR /src
COPY ./ ./
RUN go build -o /bin/go-monolith ./cmd

FROM scratch
WORKDIR /app
COPY --from=build /bin/go-monolith ./
COPY --from=build /src/db/migrations ./db/migrations
CMD [ "./go-monolith" ]
