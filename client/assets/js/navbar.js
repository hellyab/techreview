window.onscroll = function() {scrollFunction()};

function scrollFunction() {
    if (document.body.scrollTop > 200 || document.documentElement.scrollTop > 200) {
        document.getElementById("navbar").style.top = "-70px";

    } else {
        document.getElementById("navbar").style.top = "0";
    }
}