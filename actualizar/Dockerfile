FROM golang:1.20
WORKDIR /app
COPY ./ActualizarPais.go ./
RUN go mod init actualizarpais && go mod tidy
CMD ["go", "run", "ActualizarPais.go"]
EXPOSE 8083