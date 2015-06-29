#!/usr/bin/env ruby

require 'bundler/setup'

require 'csv'
require 'geocoder'
require 'json'

locations = []

# "Pension Wise Delivery Centre ","Pension Wise Email ","Pension Wise Phone Number  ","Pension Wise Contact Address"
CSV.foreach('cita.csv', headers: true, return_headers: false) do |row|
  region = row.fields.first.strip
  address = row.fields.last.strip.gsub(/\n/, ', ')
  phone = row.fields[2].strip
  lat, lng = Geocoder.coordinates(address)

  locations << {
    title: "#{region} Citizens Advice",
    address: address,
    phone: phone,
    lat: lat,
    lng: lng
  }
  #sleep(1)
end

puts JSON.pretty_generate(locations)
