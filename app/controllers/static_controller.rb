class StaticController < ApplicationController
  def home
    check_pg = PostgresDatabase.count
    render plain: "Ohai!"
  end
end
