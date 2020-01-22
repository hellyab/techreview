

var editor = new EditorJS({
    tools: { 
        header: Header, 
        linkTool: {
            class: LinkTool,
          },
          raw: RawTool,
          list: {
            class: List,
            inlineToolbar: true,
          },
      
    },

    // Other configuration properties
 
    /**
     * onReady callback
     */
    onReady: () => {
       console.log('Editor.js is ready to work!')
    }
 });