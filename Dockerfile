FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

RUN go get github.com/patrickmn/go-cache
RUN go get github.com/victorspringer/http-cache

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o wvs_proxy .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/wvs_proxy .

# Export necessary port
EXPOSE 5024

# Command to run when starting the container
CMD ["/dist/wvs_proxy"]
