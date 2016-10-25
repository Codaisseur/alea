FROM ruby:alpine

# Add Bash
RUN echo "ipv6" >> /etc/modules

# Install Mongo (we only need the client)
RUN echo 'http://nl.alpinelinux.org/alpine/edge/main' >> /etc/apk/repositories && \
  echo 'http://nl.alpinelinux.org/alpine/edge/community' >> /etc/apk/repositories && \
  echo 'http://nl.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories && \
  apk update && \
  apk upgrade && \
  apk add --no-cache mongodb

RUN apk --update add --virtual build-dependencies build-base ruby-dev openssl-dev \
  libxml2-dev libxslt-dev libc-dev linux-headers curl \
  nodejs tzdata bash && \
  rm -rf /var/cache/apk/*

RUN echo "@edge http://nl.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
  apk update && apk add "postgresql-dev@edge<9.6" "postgresql-client@edge<9.6"

ADD Gemfile /app/
ADD Gemfile.lock /app/

RUN gem install bundler && \
    cd /app ; bundle install --without development test

ADD . /app
RUN chown -R nobody:nogroup /app
USER nobody

EXPOSE 5000
ENV PORT 5000
ENV RAILS_ENV production
ENV RAILS_LOG_TO_STDOUT true
ENV RAILS_SERVE_STATIC_FILES true


WORKDIR /app

CMD ["./run.sh"]
