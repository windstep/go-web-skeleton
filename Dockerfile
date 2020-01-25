FROM golang:1.13-alpine as build_env
COPY . /app
WORKDIR /app
RUN go build main.go

FROM golang:1.13-alpine as production
COPY --from=build_env /app/.env /.env
COPY --from=build_env /app/main /main
COPY --from=build_env /app/static /static
VOLUME /static
EXPOSE 80
CMD ["/main"]