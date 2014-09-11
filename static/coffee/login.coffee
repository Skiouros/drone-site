$("#register").magnificPopup
    items:
        src: "/static/users/register.html"
    type: "ajax"

$("#loginForm").submit (e) ->
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
                $("#loginForm .error").show().html(data.Errors[0])
            else
                window.location.href = "/"

