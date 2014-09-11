$("#registerForm").submit (e) ->
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
                $("#registerForm .error").html(data.Errors[0]).show()
            else
                $("#registerForm").html "U has account"

