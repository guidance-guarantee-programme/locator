#!/usr/bin/env ruby

require 'csv'
require 'geocoder'
require 'json'

locations = []

# "REGION","DELIVERY BUREAU","PRIMARY CONTACT","POSTAL ADDRESS","Phone number"
CSV.foreach('cita.csv', headers: true, return_headers: false) do |row|
  address = row.fields[3].gsub(/\n/, ', ').strip
  phone = row.fields.last.strip
  lat, lng = Geocoder.coordinates(address)

  locations << {
    title: 'Citizens Advice',
    address: address,
    phone: phone,
    lat: lat,
    lng: lng
  }
  #sleep(1)
end

puts JSON.pretty_generate(locations)
