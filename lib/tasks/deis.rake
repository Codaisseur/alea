namespace :deis do
  desc "Create initial Deis Postgres database"
  task :create_database do
    connection_config = ActiveRecord::Base.connection_config
    query = "CREATE DATABASE #{connection_config[:database]}"

    `PGPASSWORD=#{connection_config[:password]} \
      psql --host #{connection_config[:host]} -U #{connection_config[:user]} -c "#{query};" \
      template1`
  end
end
