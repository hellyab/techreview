// document.addEventListener("DOMContentLoaded", event => {
//   event.preventDefault();
//   $("#question-card").toggle();
// });
//
// function toggleQuestion() {
//   $("#question-card").toggle();
//   $("#ask-btn").toggle();
// }
//
// $(document).ready(function() {
//   // executes when HTML-Document is loaded and DOM is ready
//   console.log("document is ready");
//
//
//   $( ".card" ).hover(
//       function() {
//         $(this).addClass('shadow-lg').css('cursor', 'pointer');
//       }, function() {
//         $(this).removeClass('shadow-lg');
//       }
//   );
//
// // document ready
// });


function submitQuestion() {
  var inquiry = document.getElementById("inquiry");

  fetch("http://localhost:8181/question", {
    method: "POST",
    mode: "no-cors",
    headers: new Headers(),
    body: JSON.stringify({ Inquiry: inquiry.value, InquirerID: "mvaney95" })
  })
    .then(toggleQuestion())
    .catch(err => console.log(err));
  document.location.reload();
}

function deleteQuestion(e, src) {
  e.preventDefault();
  fetch(`http://localhost:8181/questions/${src.getAttribute("name")}`, {
    method: "DELETE",
    headers: new Headers()
  })
    .then(res => {res.json
        $(src).closest('.question-card').remove()})
    .catch(err => console.log(err));

}
