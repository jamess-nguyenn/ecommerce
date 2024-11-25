FROM golang:1.23

# Install necessary tools
RUN apt-get update && apt-get install -y \
    git \
    build-essential \
    && apt-get clean

# Install Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set working directory
WORKDIR /var/www/go/ecommerce

# Copy the wrapper script
COPY ./golang-entrypoint.sh /golang-entrypoint.sh
RUN chmod +x /golang-entrypoint.sh

# Set entrypoint
ENTRYPOINT ["/golang-entrypoint.sh"]
