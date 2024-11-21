# syntax=docker/dockerfile:1

# Create a build stage for the app
ARG GO_VERSION=1.23.2
FROM golang:${GO_VERSION} AS build

# Add certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

# Add users here, addgroup & adduser are not available in scratch.
RUN groupadd -r appgroup && useradd -r -g appgroup -m jre

WORKDIR /src

# Cache downloaded Go modules
# Bind go.mod and go.sum from host with copying to container
# Download Go modules as a seperate step to leverage caching
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# Go SQLite wrapper mattn/go-sqlite3 requires CGo
# This causes issues for the final Alpine image
# since various C dependencies are not included.
# Simplest solution is to create a 'static build'
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o /bin/server .

#####################################################################
FROM scratch AS final

# Copy app files from host to container
COPY ./database/db.db ./database/db.db

# Copy the executable from the "build" stage.
COPY --from=build /bin/server /bin/

# Copy certificates
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy users
COPY --from=build /etc/passwd /etc/passwd

USER jre

# Expose the port that the application listens on.
EXPOSE 8080

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]

# #####################################################################
# FROM alpine:latest AS final
#
# # Copy the executable from the "build" stage.
# COPY --from=build /bin/server /bin/
#
# # Copy app files from host to container
# COPY ./database/db.db ./database/db.db

# # Copy app users
# COPY --from=build /etc/passwd /etc/passwd
#
# # Switch users
# USER jre
#
# # Expose the port that the application listens on.
# EXPOSE 8080
#
# # What the container should run when it is started.
# ENTRYPOINT [ "/bin/server" ]
