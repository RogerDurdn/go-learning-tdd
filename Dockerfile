# Docker image for golang development
FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN ["apt-get", "update"]
RUN apt-get -y install git
RUN apt-get -y install vim
RUN ["apt-get", "install", "-y", "zsh"]
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true

CMD ["zsh"]