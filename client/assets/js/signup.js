$('#registration').submit(function(e) {
    e.preventDefault();
    var first_name = $('#firstname').val();
    var last_name = $('#lastname').val();
    var email = $('#email').val();
    var password = $('#password').val();
    var conpassword = $('#conpassword').val();
 
    $(".error").remove();
 
    if (first_name.length < 1) {
      document.getElementById('one').innerHTML="This field is required";

    }
    if (last_name.length < 1) {
        document.getElementById('two').innerHTML="This field is required";
    }
    if (email.length < 1) {
        document.getElementById('three').innerHTML="This field is required";
    } else {
      var regEx = /^[A-Z0-9][A-Z0-9._%+-]{0,63}@(?:[A-Z0-9-]{1,63}\.){1,125}[A-Z]{2,63}$/;
      var validEmail = regEx.test(email);
      if (!validEmail) {
        document.getElementById('three').innerHTML="Please enter a valid email.";
      }
    }
    if (password.length < 8) {
        document.getElementById('four').innerHTML="Password must be at least 8 characters long.";
    }

    if (conpassword!=password) {
        document.getElementById('four').innerHTML="The passwords didn't match. Please try again.";
    }
  });
 
