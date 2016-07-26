class CreateMongodbServices < ActiveRecord::Migration[5.0]
  def change
    create_table :mongodb_services do |t|
      t.string :name
      t.string :db
      t.string :app
      t.string :host
      t.integer :port
      t.string :username
      t.string :password

      t.timestamps
    end
  end
end
