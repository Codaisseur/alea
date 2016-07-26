Rails.application.routes.draw do
  resources :postgres_databases, only: [:create]
  resources :redis_services, only: [:create]
  resources :mongodb_services, only: [:create]
  root "static#home"
end
