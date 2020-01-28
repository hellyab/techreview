window.onscroll = function() {scrollFunction()};

function scrollFunction() {
    if (document.body.scrollTop > 200 || document.documentElement.scrollTop > 200) {
        document.getElementById("theWholeNavbar").style.top = "-50px";

    } else {
        document.getElementById("theWholeNavbar").style.top = "0";
    }
}