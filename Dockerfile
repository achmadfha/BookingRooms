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

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_DATABASE=booking-roomsdb
ENV DB_USERNAME=postgres
ENV DB_PASSWORD=StrongPassword123!

ENV MAX_IDLE=1
ENV MAX_CONN=2
ENV MAX_LIFE_TIME=1h
ENV PORT=8080
ENV LOG_MODE=1

ENV SALT=10
ENV ROLES=admin
ENV SECRET_TOKEN=0e92db198435842f3eee72cf785ab737
ENV TOKEN_EXPIRED=24

ENTRYPOINT ["/app/booking-room"]