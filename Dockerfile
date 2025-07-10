# Install CA cert
FROM alpine AS certs
RUN apk add --no-cache ca-certificates

FROM scratch

# Copy certs from Alpine
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the Pre-built binary file
COPY /build/main /app/main
COPY /config/config.yml /config/config.yml

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
ENTRYPOINT ["/app/main"]
