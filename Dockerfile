FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.11.5.linux-amd64.tar.gz
RUN rm /tmp/go1.11.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go get github.com/joho/godotenv && \
  go get golang.org/x/oauth2 && \
  go get golang.org/x/oauth2/google && \
  go get github.com/gorilla/mux && \
  go get github.com/jinzhu/gorm && \
  go get cloud.google.com/go/datastore && \
  go get github.com/satori/go.uuid

RUN apt-get install unzip
RUN wget https://releases.hashicorp.com/terraform/0.12.7/terraform_0.12.7_linux_amd64.zip
RUN unzip terraform_0.12.7_linux_amd64.zip
RUN rm terraform_0.12.7_linux_amd64.zip
RUN mv terraform /usr/local/bin/
RUN terraform --version
RUN terraform init

# RUN apt-get install apt-transport-https ca-certificates curl software-properties-common -y
# RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg |  apt-key add –


# RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu  $(lsb_release -cs)  stable"
# RUN apt-get update
# RUN apt-get install docker-ce

WORKDIR /
COPY . .

EXPOSE 8000
# ENTRYPOINT ["terraform"]
CMD ["go","run","main.go"]
