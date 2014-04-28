require 'sinatra'
require 'json'
require 'puma'

get '/' do
  'Webhook consumer is running...'
end

post '/person' do
  person = JSON.parse(request.body.read)
  person["email"]
end
