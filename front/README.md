# Setup

Welcome! Glad you are here.

### Installation

First step would be to [install Elm](https://guide.elm-lang.org/install.html).

Then you should have the `elm` command available. To compile the application run:

`elm make src/Main.elm --output=public/app.js`

Then serve up the `public` directory. If you have `go` installed:

`go run tools/front.go`

### Tests

Check out `tests/Tests.elm`.

```
npm i -g elm-test
elm-test
```

### Live reloading

```
npm i -g elm-live
elm-live src/Main.elm -u --open --dir=public -- --output=public/app.js
```

### Contributing

Read the official [Elm Guide](https://guide.elm-lang.org/). Check out issues tagged with `good first issue` and open a draft PR so the community can gain visibility into the work and offer constructive feedback.

### Protobuf

TODO