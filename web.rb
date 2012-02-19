require 'sinatra'

set :public_folder, Proc.new { File.join(root, "public") }

get '/' do
  File.read(File.join('public', 'index.html'))
end

