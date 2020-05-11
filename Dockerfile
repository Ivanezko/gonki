FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server
FROM ubuntu
#RUN apt-get update && apt-get install -y chromium-browser
COPY --from=builder /build/server /
#ADD .env /
EXPOSE 2022
RUN chmod +x /server
ENTRYPOINT ["/server"]