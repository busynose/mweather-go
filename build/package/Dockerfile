FROM makeblock/alpine

ARG tmp_api_version=default_value
ENV API_VERSION=$tmp_api_version

ADD mweather-go /app/

WORKDIR /app

EXPOSE 8080

CMD ["./mweather-go"]