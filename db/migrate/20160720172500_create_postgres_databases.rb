class CreatePostgresDatabases < ActiveRecord::Migration[5.0]
  def change
    create_table :postgres_databases do |t|
      t.string :name
      t.string :db
      t.string :app
      t.string :username
      t.string :password
      t.string :host
      t.integer :port

      t.timestamps
    end

    add_index :postgres_databases, [:db, :host], unique: true
  end
end
