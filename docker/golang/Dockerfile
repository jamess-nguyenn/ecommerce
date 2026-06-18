# syntax=docker/dockerfile:1.4

FROM --platform=linux/amd64 golang:1.25

# Install necessary tools and dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    build-essential \
    libssl-dev \
    zlib1g-dev \
    gcc \
    make \
    wget \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Clone and build librdkafka (compatible with amd64)
RUN git clone https://github.com/edenhill/librdkafka.git /tmp/librdkafka && \
    cd /tmp/librdkafka && \
    ./configure --prefix=/usr && \
    make -j$(nproc) && \
    make install && \
    cd / && rm -rf /tmp/librdkafka

# Install Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set working directory
WORKDIR /var/www/go/ecommerce

# Copy entrypoint script
COPY ./golang-entrypoint.sh /golang-entrypoint.sh
RUN chmod +x /golang-entrypoint.sh

# Entrypoint
ENTRYPOINT ["/golang-entrypoint.sh"]
