class MongodbServicesController < ApplicationController
  def create
    mongo = MongodbService.create
    render plain: "MONGODB_URL=#{mongo.mongodb_url}"
  end
end
