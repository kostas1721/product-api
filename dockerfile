# syntax=docker/dockerfile:1

################################################################################
# Create a stage for building the application.
ARG GO_VERSION=1.23
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS build

# Install dependencies for Go and CGO with SQLite
RUN apk add --no-cache gcc g++ make libc-dev sqlite-dev

# Set up the working directory
WORKDIR /src

# Download dependencies as a separate step to take advantage of Docker's caching.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# This is the architecture you’re building for, which is passed in by the builder.
ARG TARGETARCH

# Build the application with CGO_ENABLED=1 to support go-sqlite3
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=1 GOARCH=$TARGETARCH go build -o /bin/server .

################################################################################
# Create a new stage for running the application that contains the minimal
# runtime dependencies for the application.
FROM alpine:latest AS final

# Install any runtime dependencies that are needed to run your application.
RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    sqlite-libs && \
    update-ca-certificates

# Copy the executable from the "build" stage.
COPY --from=build /bin/server /bin/

# Change to root user to change permissions
USER root
RUN chmod +x /bin/server  # Ensure the binary is executable

# Create a non-privileged user that the app will run under.
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Expose the port that the application listens on.
EXPOSE 1721

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]
