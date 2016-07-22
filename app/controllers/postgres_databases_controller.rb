class PostgresDatabasesController < ApplicationController
  def create
    pg = PostgresDatabase.create
    render text: "DATABASE_URL=#{pg.database_url}"
  end
end
