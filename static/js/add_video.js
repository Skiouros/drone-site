// Generated by CoffeeScript 1.7.1
(function() {
  var map, mapCanvas, mapOptions, marker, placeMarker;

  mapCanvas = document.getElementById("dragMap");

  mapOptions = {
    center: new google.maps.LatLng(44.8281, -98.5463),
    mapTypeControl: false,
    streetViewControl: false,
    overviewMapControl: false,
    zoom: 6,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  };

  map = new google.maps.Map(mapCanvas, mapOptions);

  marker = null;

  placeMarker = function(latLng) {
    $("#DronePos").val(latLng.toString());
    if (marker != null) {
      marker.setMap(null);
    }
    marker = new google.maps.Marker({
      position: latLng,
      map: map,
      title: "Drone location",
      draggable: true
    });
    google.maps.event.addListener(marker, "drag", function(event) {
      return $("#DronePos").val(event.latLng.toString());
    });
    return google.maps.event.addListener(marker, "dragend", function(event) {
      return $("#DronePos").val(event.latLng.toString());
    });
  };

  google.maps.event.addListener(map, "click", function(e) {
    return placeMarker(e.latLng);
  });

  $("#searchForm2").submit(function(e) {
    var address, geocoder;
    e.preventDefault();
    address = $("#add_location").val();
    geocoder = new google.maps.Geocoder();
    return geocoder.geocode({
      address: address
    }, function(results, status) {
      if (status === google.maps.GeocoderStatus.OK) {
        map.setCenter(results[0].geometry.location);
        return placeMarker(results[0].geometry.location);
      } else {
        return console.log("Geocode was not succesful: " + status);
      }
    });
  });

  $("#videoAddForm").submit(function(e) {
    var formURL, postData;
    e.preventDefault();
    postData = $(this).serializeArray();
    formURL = $(this).attr("action");
    console.log(postData);
    return $.ajax({
      url: formURL,
      type: "POST",
      data: postData,
      dataType: "json",
      success: function(data, textStatus) {
        if (data.Errors) {
          return $(".errorMessage").show().html(data.Errors[0]);
        } else {
          $(".errorMessage").hide();
          return $("#videoAddForm").html("<label>Video has been submitted</label>");
        }
      }
    });
  });

}).call(this);
