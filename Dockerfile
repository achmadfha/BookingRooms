FROM golang:alpine as build

#build stage
#create folder app
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o booking-room

# Final Stage
FROM alpine:latest
WORKDIR /app

COPY --from=build /app/booking-room /app/booking-room

ENTRYPOINT ["/app/booking-room"]