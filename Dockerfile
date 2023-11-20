# FROM golang:alpine

# RUN mkdir /app-transactions

# WORKDIR /app-transactions

# COPY . .

# RUN go mod download

# RUN go get github.com/githubnemo/CompileDaemon

# RUN go install github.com/githubnemo/CompileDaemon

# ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
FROM golang:alpine

ENV PROJECT_DIR=/app-transactions \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app-transactions
RUN mkdir "/build"
COPY . .
# RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -build="go build -o /build/app-transactions" -command="/build/app-transactions"

# ENV PROJECT_DIR=/app-transactions \
#     GO111MODULE=on \
#     CGO_ENABLED=0

# WORKDIR /app-transactions
# RUN mkdir "/build"
# COPY . .
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go install github.com/githubnemo/CompileDaemon
# ENTRYPOINT CompileDaemon -build="go build -o /build/app-transactions" -command="/build/app-transactions"
