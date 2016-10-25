namespace :deis do
  desc "Create initial Deis Postgres database"
  task :create_database do
    connection_config = URI(ENV["DATABASE_URL"])
    database = connection_config.path.match(/\/(.*)/)[1]

    db_select_query = "SELECT 1 AS result FROM pg_database WHERE datname='#{database}'"
    db_does_not_exist = `PGPASSWORD=#{connection_config.password} \
      psql --host #{connection_config.host} --port #{connection_config.port} \
      -U #{connection_config.user} \
      -c "#{db_select_query};" \
      template1 | grep '1 row' | wc -l`.to_i == 0

    if db_does_not_exist
      create_query = "CREATE DATABASE #{database}"

      `PGPASSWORD=#{connection_config.password} \
        psql --host #{connection_config.host} \
        --port #{connection_config.port} \
        -U #{connection_config.user} \
        -c "#{create_query};" \
        template1`
    end
  end
end
