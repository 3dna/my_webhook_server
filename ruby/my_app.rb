require 'active_record'
require 'awesome_print'
require 'json'
require 'puma'
require 'sinatra'

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

    person.save
  end
end

post '/update_person' do
  person = JSON.parse(request.body.read)
  full_name = person['first_name'] + ' ' + person['last_name']
  Person.synchronize(full_name, person['email'], person['id'])
end

get '/list_people' do
  rows = ""
  Person.find_each do |p|
    rows << <<-ROWS
    <tr>
      <td>#{p.id}</td>
      <td>#{p.name}</td>
      <td>#{p.email}</td>
      <td>#{p.updates}</td>
    </tr>
    ROWS
  end

  <<-TABLE
  <h4>People</h4>
  <table style="width:300px">
  <tr>
    <td>ID</td>
    <td>Name</td>
    <td>Email</td>
    <td>Updates</td>
  </tr>
  #{rows}
  </table>
  TABLE
end
