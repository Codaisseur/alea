FROM ruby:alpine

# Add Bash

# Install Mongo (we only need the client)
RUN echo http://dl-4.alpinelinux.org/alpine/edge/testing >> /etc/apk/repositories && \
  apk add --no-cache mongodb

RUN apk --update add --virtual build-dependencies build-base ruby-dev openssl-dev \
  libxml2-dev libxslt-dev postgresql-dev postgresql-client libc-dev linux-headers \
  nodejs tzdata bash && \
  rm -rf /var/cache/apk/*

ADD Gemfile /app/
ADD Gemfile.lock /app/

RUN gem install bundler && \
    cd /app ; bundle install --without development test

ADD . /app
RUN chown -R nobody:nogroup /app
USER nobody

ENV PORT 5000
ENV RAILS_ENV production
ENV RAILS_LOG_TO_STDOUT true
ENV RAILS_SERVE_STATIC_FILES true

EXPOSE 5000

WORKDIR /app

CMD ["bundle", "exec", "puma", "-C", "config/puma.rb"]
