FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /app

# Copy the code into the container
COPY . /app

RUN go mod download
RUN go build -o main ./cmd

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["./main"]

#docker build . -t golang:cloneIndeed
#docker run -p 8080:8080 golang:cloneIndeed