# Based on https://github.com/falcosecurity/falcosidekick/blob/41d530807f1a0294c0276e4cb42af68c8b26a659/Dockerfile
ARG BASE_GOLANG_18_ALPINE
ARG BASE_DISTROLESS

FROM $BASE_GOLANG_18_ALPINE as artifact
ARG GOPROXY
ARG SOURCE_REPO

ENV GOPROXY=${GOPROXY} \
    SOURCE_REPO=${SOURCE_REPO} \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src
RUN apk add --no-cache make git bash && \
    git clone --branch 2.26.0 --depth 1 ${SOURCE_REPO}/falcosecurity/falcosidekick.git . && \
    make falcosidekick && \
    chown -R 64535:64535 /src/falcosidekick && \
    chmod 0755 /src/falcosidekick

FROM $BASE_DISTROLESS
COPY --from=artifact /src/falcosidekick /falcosidekick
ENTRYPOINT [ "/falcosidekick" ]
