class MemcachedService < ApplicationRecord
  before_validation :set_defaults

  def servers_and_namespace
    "MEMCACHED_SERVERS=#{hosts} MEMCACHED_NAMESPACE=#{namespace}"
  end

  private

  def set_defaults
    self.name ||= [Faker::Hacker.verb, Faker::Hacker.noun].join(" ")
    self.hosts ||= memcached_hosts.join(",")
    create_namespace_slug
  end

  def memcached_hosts
    (ENV["MEMCACHED_SERVERS"] || "localhost").split(",")
  end

  def create_namespace_slug
    self.namespace = name.parameterize

    while MemcachedService.where(hosts: hosts, namespace: namespace).count > 0
      self.namespace = [name, Devise.friendly_token].join("-").parameterize
    end

    self.namespace = namespace.gsub(/\-/, '_')
  end
end
