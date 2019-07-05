$(window).scroll(function() {
    var elem = document.getElementById("sten");
    if ($(window).scrollTop() > 10) {
        elem.addClass('floatingNav');
    } else {
        elem.removeClass('floatingNav');
    }
});