FROM golang:1.23-bullseye as base

# RUN adduser \
#   --disabled-password \
#   --gecos "" \
#   --home "/nonexistent" \
#   --shell "/sbin/nologin" \
#   --no-create-home \
#   --uid 65532 \
#   small-user

WORKDIR /app
RUN mkdir -p /app/config
RUN mkdir -p /app/log

COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go mod verify

# Check for security vulnerabilities
ARG RUN_SECURITY_CHECK=false
RUN if [  "$RUN_SECURITY_CHECK" = "true" ]; then \
    go install github.com/securego/gosec/v2/cmd/gosec@latest && \
    go install golang.org/x/vuln/cmd/govulncheck@latest && \
    gosec ./.. && \
    govulncheck -scan=package ./... ; \
    fi

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM gcr.io/distroless/static-debian11

COPY --from=base /main .
COPY --from=base /app/config /config
COPY --from=base /app/log /log

# USER small-user:small-user

CMD ["./main"]
