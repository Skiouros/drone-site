// Generated by CoffeeScript 1.7.1
(function() {
  $("#registerForm").submit(function(e) {
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
          return $("#registerForm .error").html(data.Errors[0]).show();
        } else {
          return $("#registerForm").html("U has account");
        }
      }
    });
  });

}).call(this);
