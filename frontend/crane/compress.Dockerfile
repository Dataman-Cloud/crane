FROM digitallyseamless/nodejs-bower-grunt:5

MAINTAINER upccup <yyao@dataman-inc.com>

WORKDIR /data

ADD ./compress.sh /data/compress.sh

RUN chmod +x compress.sh && npm install -g gulp && npm install

ENTRYPOINT ["./compress.sh"]

