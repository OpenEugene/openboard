module Session exposing (Session(..), navKey, viewer)

import Browser.Navigation as Nav
import Viewer exposing (Viewer)


type Session
    = LoggedIn Nav.Key Viewer
    | Guest Nav.Key



-- INFO


viewer : Session -> Maybe Viewer
viewer session =
    case session of
        LoggedIn _ val ->
            Just val

        Guest _ ->
            Nothing


navKey : Session -> Nav.Key
navKey session =
    case session of
        LoggedIn key _ ->
            key

        Guest key ->
            key
