FROM golang:1.20-bullseye
LABEL maintainer="Ashwin Singh <ashwinsinghsingh666672@gmail.com>"
LABEL description="This tool acts as a wrapper for Nuclei and generates a formatted JSON output that can be loaded and tables."
RUN apt-get update && \
    apt-get install -y git

RUN go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
WORKDIR /usr/src/app
# Copy the entrypoint script
COPY . .
RUN chmod +x main.go
RUN go build main.go

# Set the entrypoint script as the default command
ENTRYPOINT ["./main"]