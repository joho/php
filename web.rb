require 'sinatra'
require 'sinatra/reloader' if development?

require 'compass'
require 'sass'

# Asset settings
set :public_folder, Proc.new { File.join(root, "public") }
configure do
  Compass.add_project_configuration(File.join(Sinatra::Application.root, 'config', 'compass.rb'))
end
# End asset settings

class ExternalLink < Struct.new(:title, :url)
  def self.recent_writings
    [
      new("Interview for The Fetch", "http://blog.thefetch.com/2012/10/29/interview-melbourne-local-john-barton/"),
      new("The Minimum Viable Rails Stack", "http://goodfil.ms/blog/posts/2012/09/15/minimum-viable-rails-stack/"),
      new("Will App.net be the connective tissue founders & developers can rely on?", "http://thenextweb.com/apps/2012/08/18/why-10000-people-care-app-net/"),
      new("Scaling Rails @ Melbourne Ruby Meetup", "http://jrb.tumblr.com/post/30570014929/scaling-rails-at-melbourne-roro")
    ]
  end
end

get '/stylesheets/:name.css' do
  content_type 'text/css', :charset => 'utf-8'
  scss(:"stylesheets/#{params[:name]}", Compass.sass_engine_options)
end

get '/' do
  @recent_writings = ExternalLink.recent_writings

  @interesting_projects = []
  @interesting_projects << ["This site (Sinatra + SCSS)", "https://github.com/joho/whoisjohnbarton.com"]
  @interesting_projects << ["Spoilerless Tour De France", "https://github.com/joho/letour"]
  @interesting_projects << ["Loose (and fast) Postgres Row Counts", "https://github.com/goodfilms/postgres_loose_table_counts"]
  @interesting_projects << ["Netflix API wrapper for Ruby", "https://github.com/goodfilms/netflix-ruby"]
  @interesting_projects << ["HTTP Status codes 7xx (Developer Fouls)", "https://github.com/joho/7XX-rfc"]
  erb :index
end
