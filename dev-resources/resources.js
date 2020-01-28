// // Bootstrap CSS
// <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous"></link>
//
// //These should be imported at the end of the body right before the body closing tag. IN THIS ORDER.
// //jQuery slim
// //Popper.js
// //Bootstrap JS
// <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
// <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
// <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>

//
// // video link for embedding youtube vids TESTING
// //https://www.youtube.com/embed/3AKPaq0IaDk?list=PLOHoVaTp8R7f9GV4HW3joHrz4Oq16QBph
//
// //link for the google material icons stylesheet. You will need this to include icons in the page.
// <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet"></link>
//
// //link for the icons list in the material icons set
// //https://material.io/resources/icons/?icon=add&style=outline
// //NOTE We are using outlined icons
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// //adds a textbox which takes url and displays the video in iframe when enter is hit
// //TODO make it accept any video link, not just youtube embed links
// //TODO fix the textbox
// //TODO Update everything for a removed video component.
//
// // function addVideoComponent(e, src) {
// //   e.preventDefault();
// //   $("#input-type-button-group").hide();
// //   var textBox = document.createElement("input");
// //   textBox.setAttribute("type", "url");
// //   textBox.setAttribute("class", "form-control dynamic-input");
// //   textBox.setAttribute("name", "dynamic-input");
// //   textBox.setAttribute("required", "true");
//
// //   var container = document.getElementById("element-container");
// //   container.appendChild(textBox);
// //   if (
// //     textBox.previousSibling.nodeName != "#test" &&
// //     checkPreviousElementValue(textBox.previousSibling) === true
// //   ) {
// //     removePreviousElement(textBox.previousElementSibling);
// //   }
// //   textBox.onkeypress = function(e) {
// //     if (e.key == "Enter") {
// //       container.removeChild(textBox);
// //       container.innerHTML += `<div style="position:relative;padding-top:56.25%;">
// //         <iframe src=${textBox.value}
// //         frameborder="0"
// //         allowfullscreen
// //         style="position:absolute;top:0;left:0;width:100%;height:100%;">
// //         </iframe>
// //         </div>`;
// //     }
// //   };
// // }
