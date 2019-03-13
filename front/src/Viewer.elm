module Viewer exposing (Viewer, avatar, cred, minPasswordChars)

{-| The logged-in user currently viewing this page. It stores enough data to
be able to render the menu bar (username and avatar), along with Cred so it's
impossible to have a Viewer if you aren't logged in.
-}

import Avatar exposing (Avatar)



-- TYPES


type Viewer
    = Viewer Avatar String



-- INFO


cred : Viewer -> String
cred (Viewer _ val) =
    val


avatar : Viewer -> Avatar
avatar (Viewer val _) =
    val


{-| Passwords must be at least this many characters long!
-}
minPasswordChars : Int
minPasswordChars =
    6
