FROM golang:1.20
WORKDIR /app
COPY ./MostrarPais.go ./
RUN go mod init mostrarpais && go mod tidy
CMD ["go", "run", "MostrarPais.go"]
EXPOSE 8082