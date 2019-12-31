//logs all values from inputs with name dynamic-input to the consol
//this can be used to send the values to the api after some modification
document.addEventListener("DOMContentLoaded", event => {
  console.log("loaded");
  $("#input-type-button-group").toggle();
  document.onclick = function() {
    var inputGroup = `<div id="input-type">
      <button id="show-toggler" class="btn btn-sm-toggle btn-sm bg-white rounded-circle border-dark"
          onclick="toggleButtonGroup(event)">
          <i class="material-icons">
              add
          </i>
      </button>

      <div class="btn-group" id="input-type-button-group">
          <button onclick="addImageComponent(event,this)"
              class="btn btn-sm btn-sm-in-group bg-white rounded-circle border-dark m-1"><i
                  class="material-icons">
                  add_a_photo
              </i></button>
          <button onclick="addSecondTitleComponent(event, this)"
              class="btn btn-sm btn-sm-in-group bg-white rounded-circle border-dark m-1"><i
                  class="material-icons">
                  title
              </i></button>
          <button onclick="addVideoComponent(event, this)"
              class="btn btn-sm btn-sm-in-group bg-white rounded-circle border-dark m-1"><i
                  class="material-icons">
                  play_circle_outline
              </i></button>
      </div>
  </div>`;
    if ($("#input-type")) {
      $("#input-type").remove();
    }
    if (
      document.activeElement.nodeName == "P" ||
      document.activeElement.nodeName == "H1" ||
      document.activeElement.nodeName == "H3"
    ) {
      $(inputGroup).insertAfter(document.activeElement);
    }
  };
});

function logValues() {
  var dynamicElements = document.getElementsByName("dynamic-input");
  var i;
  var atricle = {};
  for (i = 0; i < dynamicElements.length; i++) {
    atricle[`value${i}`] = dynamicElements[i].textContent;
  }
  console.log(atricle);
}

//toggles the visibility of the button group based on the current visibility status
function toggleButtonGroup(e) {
  $("#input-type-button-group").toggle();
  //   if (
  //     document.getElementById("input-type-button-group").style.visibility !=
  //     "hidden"
  //   ) {
  //     console.log("showing group condition");
  //     showGroup(e);
  //   } else {
  //     console.log("hiding group condition");
  //     hideGroup(e);
  //   }
}

//changes the visibility of the button group and the button
function showGroup(e) {
  console.log("showing group");
  e.preventDefault();
  var inputTypeButtonGroup = document.getElementById("input-type-button-group");
  inputTypeButtonGroup.style.visibility = "visible";
  var showToggler = document.getElementById("show-toggler");
  showToggler.style.visibility = "hidden";
}

//changes the visibility of the button group and the button
// function hideGroup(e) {
//   console.log("hiding group");

//   e.preventDefault(e);
//   var inputTypeButtonGroup = document.getElementById("input-type-button-group");
//   inputTypeButtonGroup.style.visibility = "hidden";
//   var showToggler = document.getElementById("show-toggler");
//   showToggler.style.visibility = "visible";
// }

//checks if the value of the node right before this node is empty
function checkPreviousElementValue(src) {
  if (
    src.nodeName != "INPUT" &&
    src.nodeName != "IMG" &&
    src.nodeName != "IFRAME"
  ) {
    if (src.textContent === "") {
      return true;
    } else {
      return false;
    }
  } else if (src.nodeName === "INPUT") {
    if ((src.value = null)) {
      return false;
    } else {
      return true;
    }
  }
}

//removes element right before this element
function removePreviousElement(src) {
  var container = document.getElementById("element-container");
  container.removeChild(src);
}

//adds a content editable text component
function addTextComponent(e, src) {
  $("#input-type-button-group").hide();

  e.preventDefault();
  toggleButtonGroup(e);
  var element = document.createElement("p");
  element.setAttribute("class", "dynamic-input");
  element.setAttribute("name", "dynamic-input");
  element.setAttribute("contenteditable", "true");
  element.setAttribute("onkeydown", "conditionalTextBox(event, this)");

  var container = document.getElementById("element-container");
  //   src.nextSibling = element;
  $(element).insertAfter(src);

  //   container.appendChild(element);
  element.focus();
}

//adds a file input which will post a preview of the image
function addImageComponent(e, src) {
  $("#input-type-button-group").hide();

  e.preventDefault(e);
  toggleButtonGroup(e);
  var container = document.getElementById("element-container");

  container.innerHTML += `<input id='uploadFile' type='file' hidden/>`;
  var element = document.getElementById("uploadFile");
  element.click();
  //TODO Add image component code here
  element.onchange = function(e) {
    var input = e.target;

    var reader = new FileReader();

    reader.onload = function() {
      fileContents = reader.result;
      var image = document.createElement("img");
      image.setAttribute("src", fileContents);
      image.setAttribute("class", "col-8 offset-2");
      container.appendChild(image);
      container.removeChild(element);
      console.log(image)
      document.activeElement = image;
      addTextComponent(e, image);
      console.log(image.nextSibling)
      $(image).nextSibling.focus();
    };

    reader.readAsDataURL(input.files[0]);
  };
}

//adds a contenteditable h3
//TODO focus issues
function addSecondTitleComponent(e, src) {
  $("#input-type-button-group").hide();

  e.preventDefault();
  toggleButtonGroup(e);
  var element = document.createElement("h3");
  element.setAttribute("class", "dynamic-input");
  element.setAttribute("name", "dynamic-input");
  element.setAttribute("contenteditable", "true");
  element.setAttribute("onkeydown", "conditionalTextBox(event, this)");

  var container = document.getElementById("element-container");
  container.appendChild(element);
  if (
    element.previousSibling.nodeName != "#test" &&
    checkPreviousElementValue(element.previousSibling) === true
  ) {
    removePreviousElement(element.previousElementSibling);
  }
}

//adds a textbox which takes url and displays the video in iframe when enter is hit
//TODO make it accept any video link, not just youtube embed links
//TODO fix the textbox
function addVideoComponent(e, src) {
  e.preventDefault();
  $("#input-type-button-group").hide();
  var textBox = document.createElement("input");
  textBox.setAttribute("type", "url");
  textBox.setAttribute("class", "form-control dynamic-input");
  textBox.setAttribute("name", "dynamic-input");
  textBox.setAttribute("required", "true");

  var container = document.getElementById("element-container");
  container.appendChild(textBox);
  if (
    textBox.previousSibling.nodeName != "#test" &&
    checkPreviousElementValue(textBox.previousSibling) === true
  ) {
    removePreviousElement(textBox.previousElementSibling);
  }
  textBox.onkeypress = function(e) {
    if (e.key == "Enter") {
      container.removeChild(textBox);
      container.innerHTML += `<div style="position:relative;padding-top:56.25%;">
        <iframe src=${textBox.value} 
        frameborder="0" 
        allowfullscreen
        style="position:absolute;top:0;left:0;width:100%;height:100%;">
        </iframe>
        </div>`;
    }
  };
}

//checks if the pressed key is enter and adds a text component
function conditionalTextBox(e, src) {
  if (e.key === "Enter") {
    src.textContent += "\xa0";
    addTextComponent(e, src);
  } else if (e.key == "Backspace" && src.textContent == "") {
    if (src.previousSibling.nodeName != "h") {
      //TODO understand what this code is really doing. It was a mix and match kinda build.
      var range = document.createRange();
      var sel = window.getSelection();
      range.setStart(
        src.previousSibling.childNodes[0],
        src.previousSibling.textContent.length
      );
      range.collapse(true);
      sel.removeAllRanges();
      sel.addRange(range);
    }
    document.getElementById("element-container").removeChild(src);
  }
}
