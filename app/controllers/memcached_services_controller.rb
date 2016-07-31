class MemcachedServicesController < ApplicationController
  def create
    memcached = MemcachedService.create
    render plain: memcached.servers_and_namespace
  end
end
