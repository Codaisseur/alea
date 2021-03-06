require 'mongo'

class MongodbService < ApplicationRecord
  before_validation :set_defaults
  after_create :setup_credentials

  def mongodb_url
    "mongodb://#{username}:#{password}@#{host}:#{port}/#{db}"
  end

  def set_defaults
    self.name ||= [Faker::Hacker.verb, Faker::Hacker.noun].join(" ")
    self.password ||= Devise.friendly_token
    self.host ||= database_host
    self.port ||= database_port
    create_db_slug
    create_username
  end

  def database_host
    connection_config.host || "localhost"
  end

  def database_port
    connection_config.port || 27017
  end

  def create_db_slug
    self.db = name.parameterize

    while MongodbService.where(host: host, db: db).count > 0
      self.db = [name, Devise.friendly_token].join("-").parameterize
    end

    self.db = db.gsub(/\-/, '_')
  end

  def create_username
    base_name = username || Faker::Internet.user_name(nil, %w())
    self.username = base_name

    tries = 0
    while MongodbService.where(host: host, username: username).count > 0
      tries += 1
      self.username = base_name + "#{tries}"
    end
  end

  private

  def connection_config
    URI(admin_mongo_url)
  end

  def admin_mongo_url
    ENV["MONGODB_URL"] || "mongodb://localhost:27017/admin"
  end

  def setup_credentials
    `mongo #{admin_mongo_url} << EOF
use #{db}
db.createUser({user: '#{username}', pwd: '#{password}', roles:[{role:'dbOwner',db:'#{db}'}]})
EOF`
  end
end
