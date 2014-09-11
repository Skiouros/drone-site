mapCanvas = document.getElementById "dragMap"
mapOptions = {
    center: new google.maps.LatLng(44.8281, -98.5463),
    mapTypeControl: false
    streetViewControl: false
    overviewMapControl: false
    zoom: 6,
    mapTypeId: google.maps.MapTypeId.ROADMAP
}
map = new google.maps.Map mapCanvas, mapOptions

marker = null

placeMarker = (latLng) ->
    $("#DronePos").val latLng.toString()
    marker.setMap null if marker?
    marker = new google.maps.Marker
        position: latLng
        map: map
        title: "Drone location"
        draggable: true

    google.maps.event.addListener marker, "drag", (event) ->
        $("#DronePos").val event.latLng.toString()

    google.maps.event.addListener marker, "dragend", (event) ->
        $("#DronePos").val event.latLng.toString()


google.maps.event.addListener map, "click", (e) ->
    placeMarker e.latLng

$("#searchForm2").submit (e) ->
    e.preventDefault()

    address = $("#add_location").val()
    geocoder = new google.maps.Geocoder()
    geocoder.geocode address: address, (results, status) ->
        if (status == google.maps.GeocoderStatus.OK)
            map.setCenter(results[0].geometry.location)
            placeMarker results[0].geometry.location
        else
            console.log "Geocode was not succesful: #{status}"

$("#videoAddForm").submit (e) ->
    e.preventDefault()
    postData = $(this).serializeArray()
    formURL = $(this).attr "action"
    console.log postData

    $.ajax
        url: formURL
        type: "POST"
        data: postData
        dataType: "json"
        success: (data, textStatus) ->
            if data.Errors
                $(".errorMessage").show().html(data.Errors[0])
            else
                $(".errorMessage").hide()
                $("#videoAddForm").html "<label>Video has been submitted</label>"

