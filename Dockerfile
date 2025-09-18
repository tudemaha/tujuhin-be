FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .
RUN touch /app/.env
RUN go build -v -o /app/tujuhin /app/cmd/main

FROM gcr.io/distroless/static-debian12

COPY --from=build /app/.env /
COPY --from=build /app/tujuhin /

CMD ["/tujuhin"]