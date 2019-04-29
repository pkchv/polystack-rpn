FROM golang:1.12.4-alpine3.9

# Set in-container working directory
WORKDIR /go/src/polystack-endpoint

# Copy source files
COPY . .

# Download dependencies (Note: '...' - recursive wildcard)
RUN go get -d -v ./...

# Install the package (Note: '...' - recursive wildcard)
RUN go install -v ./...

# Run the app
CMD ["polystack-endpoint"]
