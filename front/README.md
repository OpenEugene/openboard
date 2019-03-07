# Setup

Welcome! Glad you are here.

### Installation

First step would be to [install Elm](https://guide.elm-lang.org/install.html).

Then you should have the `elm` command available. To compile the application run:

`elm make src/Main.elm --output=public/app.js --debug`

<sub>*Please be sure to add `**/elm-stuff/` to your global git ignore file.</sub>

### Get started

You can serve up the frontend on it's own, or use spin the whole stack up together (see top level README)

```
npm i -g elm-live
elm-live src/Main.elm -u --open --dir=public -- --output=public/app.js --debug
```

### Tests

Check out `tests/Tests.elm`.

```
npm i -g elm-test
elm-test
```

### Contributing

Read the official [Elm Guide](https://guide.elm-lang.org/). Check out issues tagged with `good first issue` and open a draft PR so the community can gain visibility into the work and offer constructive feedback.

### Protobuf

Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data. Learn More(https://developers.google.com/protocol-buffers/).

To generate type aliases, encoders and decoders run `pb2elm`. This will read all the `.proto` files and generate corresponding `.elm` files. This command should be run frequently and updates to the `Proto/` modules should have there own commit. This gives us a type safe border between the client and the server. If something changes on the server — a field is added or removed, a data model is added or removed — our Elm app should fail to compile. No more `400 Bad Request` at runtime!

`pb2elm src/Proto/ ../msgs/proto/*.proto`

