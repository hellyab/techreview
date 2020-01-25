//logs all values from inputs with name dynamic-input to the consol
//this can be used to send the values to the api after some modification
document.addEventListener("DOMContentLoaded", event => {
  // console.log("loaded");
  $("#input-type-button-group").toggle();
});

// NOTE Removed the add video component
// <button onclick="addVideoComponent(event, this)"
//     class="btn btn-sm btn-sm-in-group bg-white rounded-circle border-dark m-1"><i
//         class="material-icons">
//         play_circle_outline
//     </i></button>

function logValues() {
  var dynamicElements = document.getElementsByName("dynamic-input");
  var i;
  var atricle = {};
  for (i = 0; i < dynamicElements.length; i++) {
    if (dynamicElements[i].nodeName == "IMG") {
      //TODO we need to upload the file and replace the src path with a new path
      atricle[`img${i}`] = dynamicElements[i].src;
    } else if (dynamicElements[i].nodeName == "H1") {
      atricle[`title${i}`] = dynamicElements[i].textContent;
    } else if (dynamicElements[i].nodeName == "H3") {
      atricle[`subtitle${i}`] = dynamicElements[i].textContent;
    } else if (dynamicElements[i].nodeName == "P") {
      atricle[`text${i}`] = dynamicElements[i].textContent;
    }
  }
  console.log(atricle);
}

//toggles the visibility of the button group based on the current visibility status
function toggleButtonGroup(e) {
  $("#input-type-button-group").toggle();
}

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

function moveInputSelector(e, src) {
  var pos = $(src).position();
  // console.log(pos.left, pos.top);
  document.getElementById("input-type").style.position = "absolute";
  // document.getElementById("input-type").style.zIndex = 10;
  document.getElementById("input-type").style.top = pos.top - 5 + "px";
  document.getElementById("input-type").style.left = -50 + "px";
}

//removes element right before this element
function removePreviousElement(src) {
  var container = document.getElementById("element-container");
  container.removeChild(src);
}

//adds a content editable text component
function addTextComponent(e, src) {
  // $("#input-type-button-group").hide();

  e.preventDefault();
  // toggleButtonGroup(e);
  var element = document.createElement("p");
  element.setAttribute("class", "dynamic-input");
  element.setAttribute("name", "dynamic-input");
  element.setAttribute("contenteditable", "true");
  element.setAttribute("onfocus", "moveInputSelector(event, this)");
  element.setAttribute("onkeydown", "conditionalTextBox(event, this)");

  var container = document.getElementById("element-container");
  //   src.nextSibling = element;
  $(element).insertAfter(src);
  // var elementInCharge = document.activeElement;

  //   container.appendChild(element);

  // $("#input-type").style.left = pos.left-75
  element.focus();
}

//adds a file input which will post a preview of the image
function addImageComponent(e, src) {
  // $("#input-type-button-group").hide();

  e.preventDefault(e);
  // toggleButtonGroup(e);
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
      image.setAttribute("class", "col-8 offset-2 mt-2 mb-2");
      image.setAttribute("name", "dynamic-input");
      var imageAnchor = document.createElement("a");
      imageAnchor.setAttribute("class", "imageAnchor");
      imageAnchor.setAttribute("href", "#");
      imageAnchor.appendChild(image);
      imageAnchor.setAttribute("onfocus", "moveInputSelector(event, this)");
      container.appendChild(imageAnchor);
      container.removeChild(element);
      // console.log(image);
      document.activeElement = imageAnchor;
      addTextComponent(e, imageAnchor);

      // console.log(imageAnchor.nextSibling);
      // $(imageAnchor).nextSibling.focus();
    };

    reader.readAsDataURL(input.files[0]);
  };
}

//adds a contenteditable h3
//TODO focus issues
function addSecondTitleComponent(e, src) {
  // $("#input-type-button-group").hide();

  e.preventDefault();
  // toggleButtonGroup(e);
  var element = document.createElement("h3");
  element.setAttribute("class", "dynamic-input");
  element.setAttribute("name", "dynamic-input");
  element.setAttribute("contenteditable", "true");
  element.setAttribute("onkeydown", "conditionalTextBox(event, this)");
  element.setAttribute("onfocus", "moveInputSelector(event, this)");
  var container = document.getElementById("element-container");
  container.appendChild(element);
  if (
    element.previousSibling.nodeName != "#test" &&
    checkPreviousElementValue(element.previousSibling) === true
  ) {
    removePreviousElement(element.previousElementSibling);
  }
  element.focus();
}

//checks if the pressed key is enter and adds a text component
function conditionalTextBox(e, src) {
  if (e.key === "Enter") {
    src.textContent += "\xa0";
    addTextComponent(e, src);
  } else if (
    e.key == "Backspace" &&
    (src.textContent == "" || src.nodeName == "IMG") &&
    src != src.parentNode.firstElementChild
  ) {
    if (
      src.previousSibling.nodeName != "H1" &&
      src.previousSibling.nodeName != "A"
    ) {
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
    } else if (src.previousSibling.nodeName == "IMG") {
      src.previousSibling.focus();
    }
    document.getElementById("element-container").removeChild(src);
  } else if (e.key == "ArrowUp" && src.nodeName != "H1") {
    // console.log("up");
    src.previousSibling.focus();
  } else if (e.key == "ArrowDown" && src != src.parentNode.lastElementChild) {
    // console.log("down");
    src.nextSibling.focus();
  }
}
