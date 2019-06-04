FROM golang:1

RUN apt-get update \
  && apt-get -y install --no-install-recommends apt-utils 2>&1

RUN go get -u -v \
  github.com/mdempsky/gocode \
  github.com/uudashr/gopkgs/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/acroca/go-symbols \
  golang.org/x/tools/cmd/guru \
  golang.org/x/tools/cmd/gorename \
  github.com/rogpeppe/godef \
  github.com/zmb3/gogetdoc \
  github.com/sqs/goreturns \
  golang.org/x/tools/cmd/goimports \
  golang.org/x/lint/golint \
  github.com/alecthomas/gometalinter \
  honnef.co/go/tools/... \
  github.com/golangci/golangci-lint/cmd/golangci-lint \
  github.com/mgechev/revive \
  github.com/derekparker/delve/cmd/dlv 2>&1

# gocode-gomodpre
RUN go get -x -d github.com/stamblerre/gocode 2>&1 \
  && go build -o gocode-gomod github.com/stamblerre/gocode \
  && mv gocode-gomod $GOPATH/bin/

# Install git, process tools, lsb-release (common in install instructions for CLIs)
RUN apt-get update && apt-get -y install git procps lsb-release wget

ENV SHELL /bin/bash

RUN echo 'deb http://download.opensuse.org/repositories/shells:/fish:/release:/3/Debian_9.0/ /' > /etc/apt/sources.list.d/shells:fish:release:3.list
RUN wget -nv https://download.opensuse.org/repositories/shells:fish:release:3/Debian_9.0/Release.key -O Release.key
RUN apt-key add - < Release.key
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash - && \
  apt-get install -y nodejs
RUN git clone https://github.com/smiller171/dotfiles.git ~/dotfiles
RUN cd ~/dotfiles; ./setup.sh
RUN cd; fish -c "echo test"
ENV SHELL /usr/bin/fish
