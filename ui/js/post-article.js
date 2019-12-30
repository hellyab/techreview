//logs all values from inputs with name dynamic-input to the consol
//this can be used to send the values to the api after some modification
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
  if (
    document.getElementById("input-type-button-group").style.visibility !=
    "hidden"
  ) {
    showGroup(e);
  } else {
    hideGroup(e);
  }
}

//changes the visibility of the button group and the button
function showGroup(e) {
  e.preventDefault();
  var inputTypeButtonGroup = document.getElementById("input-type-button-group");
  inputTypeButtonGroup.style.visibility = "visible";
  var showToggler = document.getElementById("show-toggler");
  showToggler.style.visibility = "hidden";
}


//changes the visibility of the button group and the button
function hideGroup(e) {
  e.preventDefault();
  var inputTypeButtonGroup = document.getElementById("input-type-button-group");
  inputTypeButtonGroup.style.visibility = "hidden";
  var showToggler = document.getElementById("show-toggler");
  showToggler.style.visibility = "visible";
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

//removes element right before this element
function removePreviousElement(src) {
  var container = document.getElementById("element-container");
  container.removeChild(src);
}

//adds a content editable text component
function addTextComponent(e, src) {
  e.preventDefault();
  var element = document.createElement("p");
  element.setAttribute("class", "dynamic-input");
  element.setAttribute("name", "dynamic-input");
  element.setAttribute("contenteditable", "true");
  element.setAttribute("onkeypress", "conditionalAddText(event, this)");

  var container = document.getElementById("element-container");
  container.appendChild(element);
  element.focus();
  console.log(element.previousSibling.textContent);
  console.log(element.previousSibling.nodeName);

  if (
    element.previousSibling.nodeName != "#test" &&
    checkPreviousElementValue(element.previousSibling) === true
  ) {
    removePreviousElement(element.previousElementSibling);
  }
}

//adds a file input which will post a preview of the image
function addImageComponent(e, src) {
  e.preventDefault(e);
  var element = document.createElement("input");
  element.setAttribute("class", "dynamic-input");
  element.setAttribute("name", "dynamic-input");
  element.setAttribute("type", "file");

  var container = document.getElementById("element-container");
  container.appendChild(element);
  //Add image component code here
  element.onchange = function(e) {
    var input = e.target;

    var reader = new FileReader();

    reader.onload = function() {
      fileContents = reader.result;
      //   console.log(element.value);
      var image = document.createElement("img");
      image.setAttribute("src", fileContents);
      image.setAttribute("class", "col-8 offset-2");
      container.appendChild(image);
      container.removeChild(element);
      addTextComponent(e);
      image.nextSibling.focus();
    };

    reader.readAsDataURL(input.files[0]);
  };
  if (
    element.previousSibling.nodeName != "#test" &&
    checkPreviousElementValue(element.previousSibling) === true
  ) {
    removePreviousElement(element.previousElementSibling);
  }
}

//adds a contenteditable h3
//TODO focus issues
function addSecondTitleComponent(e, src) {
  e.preventDefault();

  var element = document.createElement("h3");
  element.setAttribute("class", "dynamic-input");
  element.setAttribute("name", "dynamic-input");
  element.setAttribute("contenteditable", "true");
  element.setAttribute("onkeypress", "conditionalAddText(event, this)");

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
      //   element = document.createElement("iframe");
      //   element.setAttribute("class", "col-10 yt-vid");
      //   element.setAttribute(
      //     "allow",
      //     "accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
      //   );
      //   element.setAttribute("allowfullscreen", "true");
      //   element.setAttribute("src", textBox.value);
      //   container.appendChild(element);
    }
  };
}

//checks if the pressed key is enter and adds a text component
function conditionalAddText(e, src) {
  //   console.log("ENter ber");
  if (checkPreviousElementValue != false) {
    if (e.key === "Enter") {
      addTextComponent(e, src);
      console.log("ENter pressed");
    }
  }
}

