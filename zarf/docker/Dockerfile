# Build the Go Binary.
FROM golang:1.18 as build_stage_sales-api
ENV CGO_ENABLED 0
ARG BUILD_REF
# Copy the source code into the container.
# COPY . /service

# Build the admin binary.
WORKDIR /service
COPY . .
RUN go build -ldflags "-X main.build=${BUILD_REF}"

FROM alpine:3.14
ARG BUILD_REF
ARG BUILD_DATE
COPY --from=build_stage_sales-api /service/service /service/service
WORKDIR /service
CMD ["./service"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="sales-api" \
      org.opencontainers.image.authors="Filip Boye-Kofi <filip.boyekofi@gmail.com>" \
      org.opencontainers.image.source="https://github.com/spuxy/service/app/sales-api" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Spuxy"