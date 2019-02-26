require('./index.scss')

const app = require('./Main.elm').Elm.Main.init({
    node: document.getElementById('app'),
    flags: null,
})
