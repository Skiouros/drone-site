$("#passwordForm").submit (e) ->
    e.preventDefault()

    postData = $(this).serializeArray()
    formURL = $(this).attr "action"

    $.ajax
        url: formURL
        type: "POST"
        data: postData
        dataType: "json"
        success: (data, textStatus) ->
            console.log data
            if data.Errors
                $("#passwordForm .error").show().html(data.Errors[0])
            else
                $.magnificPopup.close()

