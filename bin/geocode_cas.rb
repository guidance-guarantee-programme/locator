#!/usr/bin/env ruby

require 'csv'
require 'geocoder'
require 'json'

locations = []

# "Aberdeen Citizens Advice Bureau",,"41 Union Street","Aberdeen","AB11 5BN","01224 569 750","Not available","298,721",910,"a","a"
CSV.foreach('cas.csv', headers: true, return_headers: false) do |row|
  title = row.fields.first.strip
  address = row.fields[1..4].reject { |value| value.nil? || value.strip == '' }.map(&:strip).join(', ')
  phone = row.fields[5].strip
  lat, lng = Geocoder.coordinates(address)

  locations << {
    title: title,
    address: address,
    phone: phone,
    lat: lat,
    lng: lng
  }
  #sleep(1)
end

puts JSON.pretty_generate(locations)
