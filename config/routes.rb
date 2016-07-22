Rails.application.routes.draw do
  resources :postgres_databases, only: [:create]
  root "static#home"
end
