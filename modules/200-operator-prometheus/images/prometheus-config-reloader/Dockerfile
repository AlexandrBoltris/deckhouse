ARG BASE_DISTROLESS
ARG BASE_GOLANG_18_ALPINE

FROM $BASE_GOLANG_18_ALPINE as artifact

ARG GOPROXY
ARG SOURCE_REPO

ENV GOPROXY=${GOPROXY} \
    SOURCE_REPO=${SOURCE_REPO} \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git
RUN git clone --depth 1 --branch v0.62.0 ${SOURCE_REPO}/prometheus-operator/prometheus-operator.git /prometheus-operator
WORKDIR /prometheus-operator/

RUN go build -ldflags="-s -w" -o prometheus-config-reloader cmd/prometheus-config-reloader/main.go && \
    chown -R 64535:64535 /prometheus-operator/ && \
    chmod 0700 /prometheus-operator/prometheus-config-reloader

FROM $BASE_DISTROLESS

COPY --from=artifact /prometheus-operator/prometheus-config-reloader /bin/
ENTRYPOINT ["/bin/prometheus-config-reloader"]
