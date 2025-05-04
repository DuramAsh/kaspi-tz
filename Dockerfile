FROM golang:1.23-alpine as builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o kaspi-tz .

FROM alpine:3.18 as hoster

RUN apk add --no-cache tzdata
ENV TZ=Asia/Qyzylorda

COPY --from=builder /src/.env ./.env
COPY --from=builder /src/migrations ./migrations
COPY --from=builder /src/kaspi-tz ./kaspi-tz

ENTRYPOINT [ "./kaspi-tz" ]