class PostgresDatabase < ApplicationRecord
  before_validation :set_defaults
  after_create :setup_credentials, :setup_database

  def database_url
    "postgres://#{username}:#{password}@#{host}:#{port}/#{db}"
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
    connection_config[:host] || "localhost"
  end

  def database_port
    connection_config[:port] || 5432
  end

  def connection_config
    @connection_config ||= ActiveRecord::Base.connection_config
  end

  def create_db_slug
    self.db = name.parameterize

    while PostgresDatabase.where(host: host, db: db).count > 0
      self.db = [name, Devise.friendly_token].join("-").parameterize
    end

    self.db = db.gsub(/\-/, '_')
  end

  def create_username
    base_name = username || Faker::Internet.user_name(nil, %w())
    self.username = base_name

    tries = 0
    while PostgresDatabase.where(host: host, username: username).count > 0
      tries += 1
      self.username = base_name + "#{tries}"
    end
  end

  private

  def setup_credentials
    run_query "CREATE USER #{username} WITH PASSWORD '#{password}'"
  end

  def setup_database
    run_query "CREATE DATABASE #{db}"
    run_query "GRANT ALL PRIVILEGES ON DATABASE #{db} TO #{username}"
  end

  def run_query(query)
    Rails.logger.info connection_config

    `PGPASSWORD=#{connection_config[:password]} \
      psql --host #{connection_config[:host]} \
        --port #{connection_config[:port]} \
        -U #{connection_config[:username]} \
        -d template1 \
        -c "#{query};"`
  end
end
