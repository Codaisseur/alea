class StaticController < ApplicationController
  def home
    render plain: "Ohai!"
  end
end
