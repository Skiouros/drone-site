root = exports ? this

$(document).ready () ->
    setupMap()
    setupSideBar()
    $.extend true, $.magnificPopup.defaults,
        removalDelay: 300
        mainClass: "my-mfp-zoom-in",
        type: "ajax"
        closeOnContentClick: false
        fixedContentPos: true
    setupModals()

setupMap = ->
    mapCanvas = document.getElementById('map-canvas')
    mapOptions = {
        center: new google.maps.LatLng(44.8281, -98.5463),
        zoom: 6,
        mapTypeId: google.maps.MapTypeId.ROADMAP
    }
    root.map = new google.maps.Map(mapCanvas, mapOptions)

    root.codeAddress = () ->
        address = $("#address").val()
        console.log address
        root.geocoder = new google.maps.Geocoder()
        geocoder.geocode address: address, (results, status) ->
            if (status == google.maps.GeocoderStatus.OK)
                map.setCenter(results[0].geometry.location)
            else
                console.log "Geocode was not succesful: #{status}"

    $(".search").submit (e) ->
        e.preventDefault()
        codeAddress()

setupSideBar = ->
    # Sidebar tab handling code
    tabs = [ "latest", "editors", "popular", "favourites" ]
    switchTab = () ->
        $(".drone-list.#{tab}").hide() for tab in tabs
        $(".drone-list.#{this.id}").show()

        $("\##{tab}").attr "class", "" for tab in tabs
        $(this).attr "class", "active"

    $("\##{tab}").click(switchTab) for tab in tabs

    $('#menu-trigger').click (e) ->
        e.preventDefault()
        $("body").toggleClass "menu-active"


setupModals = ->
    $("#contact").magnificPopup
        items:
            src: "/static/pages/contact_us.html"
        type: "ajax"

    $("#about").magnificPopup
        items:
            src: "/static/pages/about_us.html"
        type: "ajax"

    $("#login").magnificPopup
        items:
            src: "/static/users/login.html"
        type: "ajax"

    $("#add_drone").magnificPopup
        items:
            src: "/static/drones/add.html"
        type: "ajax"
