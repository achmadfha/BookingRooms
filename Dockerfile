FROM golang:alpine as build

#build stage
#create folder app
WORKDIR /app

COPY . .

COPY .env .

RUN go mod download

RUN go build -o booking-room

# Final Stage
FROM alpine:latest
WORKDIR /app

COPY --from=build /app/booking-room /app/booking-room
RUN ls -la /app
RUN cat /app/.env

ENTRYPOINT ["/app/booking-room"]