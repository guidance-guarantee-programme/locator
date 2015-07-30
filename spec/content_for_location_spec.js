describe('contentForLocation()', function() {
  describe('with no separate booking location', function() {
    var locationWithoutBookingCentre = {
      "title": "Antrim",
      "address": "Farranshane House\n1 Ballygore Road\nAntrim\nBT41 2RN",
      "phone": "028 9442 8176",
      "hours": "Monday to Thursday, 10am to 12pm and 2pm to 4pm\nFriday, 10am to 12pm and 1:30pm to 2:30pm"
    };

    var locationWithDuplicateBookingCentre = {
      "title": "Antrim",
      "address": "Farranshane House\n1 Ballygore Road\nAntrim\nBT41 2RN",
      "booking_centre": "Antrim",
      "phone": "028 9442 8176",
      "hours": "Monday to Thursday, 10am to 12pm and 2pm to 4pm\nFriday, 10am to 12pm and 1:30pm to 2:30pm"
    };

    var html = [
      '<b>',
        'Antrim Citizens Advice',
      '</b>',
      '<p>',
        'Farranshane House<br />1 Ballygore Road<br />Antrim<br />BT41 2RN',
      '</p>',
      '<p>',
        'Monday to Thursday, 10am to 12pm and 2pm to 4pm<br />Friday, 10am to 12pm and 1:30pm to 2:30pm',
      '</p>',
      '028 9442 8176'
    ].join("\n");

    describe('and booking_centre is not in the data', function() {
      it('renders the html', function() {
        expect(contentForLocation(locationWithoutBookingCentre)).toEqual(html);
      });
    });

    describe('and booking_centre is in the data', function() {
      it('renders the html', function() {
        expect(contentForLocation(locationWithDuplicateBookingCentre)).toEqual(html);
      });
    });
  });

  describe('with a separate booking location', function() {
    var location = {
      "title": "Antrim",
      "address": "Farranshane House\n1 Ballygore Road\nAntrim\nBT41 2RN",
      "booking_centre": "Belfast Citizens Advice",
      "phone": "028 9442 8176",
      "hours": "Monday to Thursday, 10am to 12pm and 2pm to 4pm\nFriday, 10am to 12pm and 1:30pm to 2:30pm"
    };

    var html = [
      '<b>',
        'Antrim Citizens Advice',
      '</b>',
      '<p>',
        'Farranshane House<br />1 Ballygore Road<br />Antrim<br />BT41 2RN',
      '</p>',
      '<p><b>Booking Centre Details</b></p>', 
      '<p>',
        'Belfast Citizens Advice',
      '</p>',
      '<p>',
        'Monday to Thursday, 10am to 12pm and 2pm to 4pm<br />Friday, 10am to 12pm and 1:30pm to 2:30pm',
      '</p>',
      '028 9442 8176'
    ].join("\n");

    it('renders the html', function() {
      expect(contentForLocation(location)).toEqual(html);
    });
  });
});
