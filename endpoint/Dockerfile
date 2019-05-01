FROM golang:1.12.4

# Force the go compiler to use modules
ENV GO111MODULE=on

# Set in-container working directory
WORKDIR /go/src/endpoint

# Copy source files
COPY . .

# Download dependencies (Note: '...' - recursive wildcard)
RUN go get -d -v ./...

# Install the package (Note: '...' - recursive wildcard)
RUN go install -v ./...

# Run the app
CMD ["endpoint"]
