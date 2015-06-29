#!/usr/bin/env ruby

require 'bundler/setup'

require 'csv'
require 'geocoder'
require 'json'

locations = []

# "Antrim","Farranshane House","1 Ballygore Road","Antrim","BT41 2RN","028 9442 8176"
CSV.foreach('nicab.csv', headers: true, return_headers: false) do |row|
  title = row.fields.first.strip
  address = row.fields[1..4].reject { |value| value.nil? || value.strip == '' }.map(&:strip).join(', ')
  phone = row.fields.last.strip
  lat, lng = Geocoder.coordinates(address)

  locations << {
    title: "#{title} Citizens Advice Bureau",
    address: address,
    phone: phone,
    lat: lat,
    lng: lng
  }
  #sleep(1)
end

puts JSON.pretty_generate(locations)
