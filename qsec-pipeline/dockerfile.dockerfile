# -------- build stage --------
FROM golang:1.23-alpine AS build
RUN apk add --no-cache gcc musl-dev git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG BIN
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o /out/${BIN} ./cmd/${BIN}

# -------- runtime stage --------
FROM gcr.io/distroless/static:nonroot
ARG BIN
COPY --from=build /out/${BIN} /usr/local/bin/${BIN}
USER 65532:65532
ENTRYPOINT ["/usr/local/bin/qsec-controller"]