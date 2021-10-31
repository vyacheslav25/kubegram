FROM golang:1.16
WORKDIR /build
COPY . /build
RUN go get ./bot
RUN GOOS=linux go build


FROM ubuntu:20.04
RUN apt update -y && apt install curl -y
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && chmod +x ./kubectl && mv ./kubectl /usr/local/bin/kubectl
WORKDIR /app
COPY --from=0 /build/kubegram /app/kubegram
CMD ["./kubegram"]


