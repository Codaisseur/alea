class RedisService < ApplicationRecord
  before_validation :set_defaults

  def redis_url
    "redis://#{host}:#{port}/#{db}"
  end

  private

  def set_defaults
    self.name ||= [Faker::Hacker.verb, Faker::Hacker.noun].join(" ")
    self.host ||= redis_host
    self.port ||= redis_port
    create_db_slug
  end

  def redis_host
    connection_config.host
  end

  def redis_port
    connection_config.port
  end

  def connection_config
    URI(ENV["REDIS_URL"] || "redis://localhost:6379/redis_services")
  end

  def create_db_slug
    self.db = name.parameterize

    while RedisService.where(host: host, db: db).count > 0
      self.db = [name, Devise.friendly_token].join("-").parameterize
    end

    self.db = db.gsub(/\-/, '_')
  end
end
