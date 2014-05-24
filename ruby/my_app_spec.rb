require 'minitest/autorun'
require 'minitest/spec'
require 'rack/test'

require_relative 'my_app'

ENV['RACK_ENV'] = 'test'

include Rack::Test::Methods

def app
  Sinatra::Application
end

describe Person do
  after do
    Person.destroy_all
  end

  after do
    Person.destroy_all
  end

  it 'can synchronize data' do
    Person.count.must_equal 0

    Person.synchronize('Joe Smith', 'js@example.com', 4567)
    Person.count.must_equal 1
    Person.first.reload.updates.must_equal 1

    Person.synchronize('Joe Smith', 'js@example.com', 4567)
    Person.count.must_equal 1
    Person.first.reload.updates.must_equal 2

    Person.synchronize('Joe Smith', 'js@example.com', 4567)
    Person.count.must_equal 1
    Person.first.reload.updates.must_equal 3

    Person.synchronize('Tom Smith', 'ts@example.com', 1111)
    Person.count.must_equal 2
    Person.last.reload.updates.must_equal 1
  end

  describe 'the webhook server' do
    it 'can list the records for people' do
      get '/list_people'
      last_response.body.must_include 'People'
    end

    it 'can accept a post with a person record' do
      json = File.read('../person_created.json')
      post '/update_person', json, "CONTENT_TYPE" => "application/json"
      last_response.ok?.must_equal true
      Person.count.must_equal 1
    end
  end

end
