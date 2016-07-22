namespace :deis do
  desc "Create initial Deis Postgres database"
  task :create_database do
    connection_config = URI(ENV["DATABASE_URL"])
    database = connection_config.path.match(/\/(.*)/)[1]
    query = "CREATE DATABASE #{database}"

    `PGPASSWORD=#{connection_config.password} \
      psql --host #{connection_config.host} --port #{connection_config.port} -U #{connection_config.user} -c "#{query};" \
      template1`
  end
end
