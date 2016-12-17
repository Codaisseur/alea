class PostgresDatabasesController < ApplicationController
  def create
    pg = PostgresDatabase.create(postgres_params)
    render plain: "DATABASE_URL=#{pg.database_url}"
  end

  private

  def postgres_params
    params.permit(:app)
  end
end
