FROM golang:1.22.1

# Inicializa um módulo Go
RUN go mod init hello-go

# Copia o código-fonte para o contêiner
COPY . .

# Compila o servidor
RUN go build -o server .

# Define o comando padrão para iniciar o servidor
CMD ["./server"]
