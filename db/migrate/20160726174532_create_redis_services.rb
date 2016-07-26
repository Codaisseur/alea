class CreateRedisServices < ActiveRecord::Migration[5.0]
  def change
    create_table :redis_services do |t|
      t.string :name
      t.string :db
      t.string :app
      t.string :host
      t.integer :port

      t.timestamps
    end
  end
end
