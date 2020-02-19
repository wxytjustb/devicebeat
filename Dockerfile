FROM "golang:1.13.8-stretch"

# 初始化环境
RUN go get -u github.com/magefile/mage

WORKDIR /go/src/github.com/wxytjustb/devicebeat

COPY . /go/src/github.com/wxytjustb/devicebeat/

RUN cp -r ./vendor/github.com/elastic /go/src/github.com/

RUN mage build

RUN chmod +x devicebeat;cp devicebeat /usr/local/bin/

RUN cp devicebeat.yml /etc/

CMD ["devicebeat", "-c", "/etc/devicebeat.yml"]
