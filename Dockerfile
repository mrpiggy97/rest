ARG VARIANT="1.18-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}
WORKDIR /rest
COPY . /rest
RUN go mod tidy
RUN go build main.go
CMD ["./main"]