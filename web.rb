require 'sinatra'
require 'sinatra/reloader' if development?

require 'compass'
require 'sass'

$: << File.dirname(__FILE__) + "/lib"
require 'external_link'
require 'quote'

# Asset settings
set :public_folder, Proc.new { File.join(root, "public") }
configure do
  Compass.add_project_configuration(File.join(Sinatra::Application.root, 'config', 'compass.rb'))
end
# End asset settings

get '/stylesheets/:name.css' do
  content_type 'text/css', :charset => 'utf-8'
  scss(:"stylesheets/#{params[:name]}", Compass.sass_engine_options)
end

get '/' do
  @recent_writings = ExternalLink.recent_writings
  @interesting_projects = ExternalLink.interesting_projects
  @social_networks = ExternalLink.social_networks

  @quote = Quote.random

  erb :index
end
