require 'sinatra'
require 'compass'
require 'sass'

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
  File.read(File.join('public', 'index.html'))
end
