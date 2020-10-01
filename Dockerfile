FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache tzdata

ENV NODE_ENV=production
ENV TZ="Asia/Shanghai"

# redis
ENV MySqlUrl="xxx"
ENV MySqlPort="xxx"
ENV MySqlPass="xxx"

ENV RedisUrl="xxx"
ENV RedisPort="xxx"
ENV RedisPass="xxx"

CMD ["./server"]
