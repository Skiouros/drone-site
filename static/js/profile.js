// Generated by CoffeeScript 1.7.1
(function() {
  $("#passwordForm").submit(function(e) {
    var formURL, postData;
    e.preventDefault();
    postData = $(this).serializeArray();
    formURL = $(this).attr("action");
    return $.ajax({
      url: formURL,
      type: "POST",
      data: postData,
      dataType: "json",
      success: function(data, textStatus) {
        console.log(data);
        if (data.Errors) {
          return $("#passwordForm .error").show().html(data.Errors[0]);
        } else {
          return $.magnificPopup.close();
        }
      }
    });
  });

}).call(this);
