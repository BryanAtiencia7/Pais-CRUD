FROM golang:1.20
WORKDIR /app
COPY ./CrearPais.go ./
RUN go mod init crearpais && go mod tidy
CMD ["go", "run", "CrearPais.go"]
EXPOSE 8081