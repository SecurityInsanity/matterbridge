FROM alpine:edge AS builder

COPY . /go/src/github.com/SecurityInsanity/matterbridge
RUN apk update && apk add go git gcc musl-dev \
        && cd /go/src/github.com/SecurityInsanity/matterbridge \
        && export GOPATH=/go \
        && go get \
        && go build -x -ldflags "-X main.githash=$(git log --pretty=format:'%h' -n 1)" -o /bin/matterbridge

FROM alpine:edge
RUN apk --no-cache add ca-certificates mailcap
COPY --from=builder /bin/matterbridge /bin/matterbridge
RUN mkdir /etc/matterbridge \
  && touch /etc/matterbridge/matterbridge.toml \
  && ln -sf /matterbridge.toml /etc/matterbridge/matterbridge.toml
ENTRYPOINT ["/bin/matterbridge", "-conf", "/etc/matterbridge/matterbridge.toml"]
