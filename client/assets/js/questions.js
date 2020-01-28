
$(document).ready(function() {
    var userID = $("#savior").val()

    console.log(userID)
    // executes when HTML-Document is loaded and DOM is ready
    console.log("document is ready");
    var questionIDs = $(".question-id")
    console.log(questionIDs[0].value)
    for (let i=0; i<questionIDs.length; i++){
        fetch("http://localhost:8181/questions/follow/status", {
            method: "POST",
            headers: new Headers(),
            body: JSON.stringify({ "UserID": userID, "QuestionID": questionIDs[i].value })
        })
            .then((response) => {
                return response.json().then((data) => {
                    if (data){
                        console.log("did this")
                        $(`#follow-${questionIDs[i].value}`).addClass("btn-success")
                        $(`#follow-${questionIDs[i].value}`).removeClass("btn-outline-success")
                    }else{
                        console.log("did this too")
                        $(`#follow-${questionIDs[i].value}`).removeClass("btn-success")
                        $(`#follow-${questionIDs[i].value}`).addClass("btn-outline-success")                    }
                })})
            .catch(err => console.log(err));
    }})

// function follow(src) {
//     console.log(src.id.substr(7))
//     fetch("http://localhost:8181/questions/follow", {
//         method: "POST",
//         mode: "no-cors",
//         headers: new Headers(),
//         body: JSON.stringify({"UserID": userID, "QuestionID": src.id.substr(7)})
//     }).then(res=>{
//         console.log(res)
//     }).catch(err => console.log(err))
// }