class CreateMemcachedServices < ActiveRecord::Migration[5.0]
  def change
    create_table :memcached_services do |t|
      t.string :name
      t.string :namespace
      t.string :app
      t.string :hosts

      t.timestamps
    end
  end
end
