# Setup

Welcome! Glad you are here.

We use `elm-app` as a convenience in development. It injects things into the source code for Hot Module Reloading — this can be useful for debugging a deeply nested route/part of the application where you want to recompile while maintaining state.

```
npm i -g elm-app
npm start
```

## Generate Elm Decoders/Encoders from *.proto files

We will have a type safe backend, frontend and everything in between.

1. [Install Protobuf Compiler, `protoc`](https://medium.com/@erika_dike/installing-the-protobuf-compiler-on-a-mac-a0d397af46b8)

2. Install `protoc-gen-elm`: 

```
go get github.com/tiziano88/elm-protobuf/protoc-gen-elm
```

3. Run `protoc --elm_out=. *.proto -Ivendor-extra/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ -I.`