class PostgresDatabasesController < ApplicationController
  def create
    pg = PostgresDatabase.create
    render plain: "DATABASE_URL=#{pg.database_url}"
  end
end
