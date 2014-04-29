require 'active_record'
require 'sinatra'
require 'json'
require 'puma'
require 'awesome_print'

ActiveRecord::Base.logger = Logger.new(File.open('database.log', 'w'))

ActiveRecord::Base.establish_connection(
  :adapter  => 'sqlite3',
  :database => 'people.db'
)

ActiveRecord::Schema.define do
  if !ActiveRecord::Base.connection.tables.include? 'people'
    create_table :people do |table|
      table.column :name,  :string
      table.column :email, :string
      table.column :nationbuilder_id, :integer
      table.column :updates, :integer  # count # of updates
    end
  end
end

class Person < ActiveRecord::Base
  has_many :tracks

  def self.synchronize(name, email, id)
    person = Person.find_or_create_by(nationbuilder_id: id)
    person.name = name
    person.email = email
    person.updates = person.updates.to_i + 1

    ap person

    person.save
  end
end



post '/person' do
  person = JSON.parse(request.body.read)
  full_name = person['first_name'] + ' ' + person['last_name']
  Person.synchronize(full_name, person['email'], person['id'])
end

get '/' do
  'Webhook consumer is running...'
end
