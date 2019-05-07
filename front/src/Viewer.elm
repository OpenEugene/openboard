module Viewer exposing (Viewer(..), unbox)


type Viewer
    = Viewer String


unbox : Viewer -> String
unbox (Viewer v) =
    v
