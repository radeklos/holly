FROM alpine:3.4

ENV TZ=Europe/Prague
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

## Environment Variables
ENV GOPATH /holly
ENV GOBIN /usr/local/bin
ENV PATH $PATH:$GOPATH/bin

# OS Dependencies
RUN apk update && apk add go \
        tzdata \
        go-tools \
        git \
        ca-certificates \
        make \
        bash \
        wget \
    && rm -rf /var/cache/apk/*

# Set working Directory
WORKDIR /holly

# GPM (Go Package Manager)
RUN wget https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm \
        && chmod +x gpm \
        && mv gpm /usr/local/bin

# Set our final working dir to be where the source code lives
WORKDIR /holly/src/github.com/radeklos/holly

# Copy source code into the deepmind src directory so Go can build the package
COPY ./ /holly/src/github.com/radeklos/holly
RUN gpm install

# Install the go package
RUN go install main.go

# Set the default entrypoint to be deepmind
ENTRYPOINT main
