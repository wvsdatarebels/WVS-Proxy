FROM golang:alpine

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy the code into the container
COPY . .

# Export necessary port
EXPOSE 5024

# Command to run when starting the container
CMD ["/app/dist/linux/wvs_proxy"]
