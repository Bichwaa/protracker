import * as Neu from './neutralino.js'

export default {
    install: (app, options) => {
      // Plugin code goes here;
      console.log("Neu plugin =====>", Neu)
      app.config.globalProperties.$neu =  Neu;
    }
  }



  