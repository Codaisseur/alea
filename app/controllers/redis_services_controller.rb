class RedisServicesController < ApplicationController
  def create
    redis = RedisService.create
    render plain: "REDIS_URL=#{redis.redis_url}"
  end
end
