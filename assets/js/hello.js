$(document).ready(function() {
    $.ajax({
        url: "/login"
    }).then(function(data) {
       $('.title').append(data.title);
    });
});
