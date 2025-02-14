FROM golang:1.20
WORKDIR /app
COPY ./EliminarPais.go ./
RUN go mod init eliminarpais && go mod tidy
CMD ["go", "run", "EliminarPais.go"]
EXPOSE 8084