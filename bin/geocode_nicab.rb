#!/usr/bin/env ruby

require 'csv'
require 'geocoder'
require 'json'

locations = []

# "Antrim CAB","Farranshane House","1 Ballygore Road","Antrim","BT41 2RN"
CSV.foreach('nicab.csv', headers: true, return_headers: false) do |row|
  title = row.fields.first.strip
  address = row.fields[1..4].reject { |value| value.nil? || value.strip == '' }.map(&:strip).join(', ')
  lat, lng = Geocoder.coordinates(address)

  locations << {
    title: title,
    address: address,
    phone: '',
    lat: lat,
    lng: lng
  }
end

puts JSON.pretty_generate(locations)
