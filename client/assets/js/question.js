document.addEventListener("DOMContentLoaded", event => {
  event.preventDefault();
  $("#question-card").toggle();
});

function toggleQuestion() {
  $("#question-card").toggle();
  $("#ask-btn").toggle();
}

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
