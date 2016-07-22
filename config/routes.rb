Rails.application.routes.draw do
  resources :postgres_databases, only: [:create]
end
