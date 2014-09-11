root = exports ? this
markers = []
markerCluster = null

$(document).ready () ->
    root.map = setupMap('map-canvas')
    $(".search").submit (e) ->
        e.preventDefault()
        codeAddress("#address", root.map)


    setupSideBar()
    $.extend true, $.magnificPopup.defaults,
        removalDelay: 300
        mainClass: "my-mfp-zoom-in",
        type: "ajax"
        closeOnContentClick: false
        fixedContentPos: true
    setupModals()
    getVideos()

setupMap = (id) ->
    mapCanvas = document.getElementById(id)
    mapOptions = {
        center: new google.maps.LatLng(44.8281, -98.5463),
        zoom: 6,
        mapTypeId: google.maps.MapTypeId.ROADMAP
    }

    return new google.maps.Map(mapCanvas, mapOptions)

codeAddress = (id, map) ->
        address = $(id).val()
        root.geocoder = new google.maps.Geocoder()
        geocoder.geocode address: address, (results, status) ->
            if (status == google.maps.GeocoderStatus.OK)
                map.setCenter(results[0].geometry.location)
            else
                console.log "Geocode was not succesful: #{status}"

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

    $("#add_drone").magnificPopup
        items:
            src: "/static/drones/add.html"
        type: "ajax"

    $.get "/users", (data) ->
        src = "/users/profile"
        if data.Errors
            src = "/static/users/login.html"

        $("#login").magnificPopup
            items:
                src: src
            type: "ajax"

root.getLatLong = (str) ->
    latlngStr = str.replace("(", "").split(",",2)
    lat = parseFloat(latlngStr[0])
    lng = parseFloat(latlngStr[1])
    return new google.maps.LatLng(lat, lng)

placeMarker = (video) ->
    marker = new google.maps.Marker
        content: video
        position: root.getLatLong video.Position
        map: root.map

    google.maps.event.addListener marker, "click", (marker, i) ->
        openVideo this.content

    markers.push marker

getVideos = ->
    $.get "/videos", (data) ->
        for video in data
            placeMarker video

        markerCluster = new MarkerClusterer root.map, markers

openVideo = (video) ->
    console.log "Playing Video #{video.Url}"
    window.history.pushState("", "Play Vdieo", "/play/" + video.Id)
    $.magnificPopup.open
        removalDelay: 0
        mainClass: 'my-mfp-zoom-in',
        items: video.Url
        type: "inline"
        callbacks:
            close: ->
                window.history.pushState "", "Home", "/"
        inline:
            markup: """
                <div class="mfp-iframe-scaler">
                    <button title="Close (Esc)" type="button" class="mfp-close">×</button>
                    <iframe src="http://www.youtube.com/embed/#{video.Url}?autoplay=1&amp;rel=0" class="mfp-iframe" frameborder="0" allowfullscreen=""></iframe>
                    <div class="mfp-bottom-bar footer">
                        <ul class="no-style menu social-share float-left social-links-block cf hide-mobile ss-initialized" data-share-url="http://www.travelbydrone.com/play/79">
                            <li><a target="_blank" class="share-facebook facebook" href="#">Share via Facebook</a>
                            </li>
                            <li><a target="_blank" class="share-twitter twitter" href="#">Share via Twitter</a>
                            </li>
                            <li><a target="_blank" class="share-google google" href="#">Share via Google+</a>
                            </li>
                            <li><a class="report action text-center trans" href="/reports/add/79">Report!</a>
                            </li>
                            <li class="mia disabled"><a class="addToFavourite" href="#">Add To Favourites</a>
                            </li>
                            <li class="mia disabled">
                                <div class="rating" data-average="null" data-id="79" style="height: 20px; width: 115px; overflow: hidden; z-index: 1; position: relative;">
                                    <div class="jRatingColor"></div>
                                    <div class="jRatingAverage" style="width: 0px; top: -20px;"></div>
                                    <div class="jStar" style="width: 115px; height: 20px; top: -40px; background: url(http://travelbydrone.com/img/stars.png) repeat-x;"></div>
                                </div>
                            </li>
                        </ul>
                    </div>
                </div>
            """
