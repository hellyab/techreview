var resData = 0;
var editor = new EditorJS({
  tools: {
    header: Header,
    linkTool: {
      class: LinkTool
    },
    raw: RawTool,
    list: {
      class: List,
      inlineToolbar: true
    },
    image: {
      class: ImageTool,
      config: {
        uploader: {
          /**
           * Upload file to the server and return an uploaded image handler
           * @param {File} file - file selected from the device or pasted by drag-n-drop
           * @return {Promise.<{success, file: {url}}>}
           */
          uploadByFile(file) {
            // your own uploading logic here

            var fd = new FormData();

            fd.append("file", file);

            return $.ajax({
              url: "upload",
              type: "post",
              data: fd,
              dataType: "json",
              contentType: false,
              processData: false
            })
              .done(function(data) {
                console.log(data);
                return new Promise((resolve, reject) => {
                  resolve(data);
                  reject(data);
                });
              })
          },
          /**
           * Send URL-string to the server. Backend should load image by this URL and return an uploaded image handler
           * @param {string} url - pasted image URL
           * @return {Promise.<{success, file: {url}}>}
           */
          uploadByUrl(url) {
            // your ajax request for uploading
            return MyAjax.upload(file).then(() => {
              return {
                success: 1,
                file: {
                  url:
                    "https://codex.so/upload/redactor_images/o_e48549d1855c7fc1807308dd14990126.jpg"
                  // any other image handler you want to store, such as width, height, color, extension, etc
                }
              };
            });
          }
        }
      }
    }
  },

  // Other configuration properties

  /**
   * onReady callback
   */
  onReady: () => {
    console.log("Editor.js is ready to work!");
  }
});

function save(){
  editor.save().then((outputData) => {

      fetch("http://localhost:8181/articles", {
    method: "POST",
    mode: "no-cors",
    headers: new Headers(),
    body: JSON.stringify({ Author: "llavignee1", Content: outputData, Topics:["techy techy", "Techa techa"], AverageRating: 0, NumberOfRatings: 0 })
  })
    .then(console.log("Success"))
    .catch(err => console.log(err))
  }).catch((error) => {
    console.log('Saving failed: ', error)
  });
}
