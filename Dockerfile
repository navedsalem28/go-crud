FROM golang:1.20-alpine


# Set destination for COPY
WORKDIR /app

# Print Directory Path
RUN pwd && ls

# Copy the source code.
COPY . .

# Download Go modules
RUN go mod download

# Build
RUN go build -o go-crud

# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080


# Run executable file
CMD [ "./go-crud" ]
