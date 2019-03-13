module Profile exposing (Profile, avatar, bio)

{-| A user's profile - potentially your own!
Contrast with Cred, which is the currently signed-in user.
-}

import Avatar exposing (Avatar)



-- TYPES


type Profile
    = Profile Internals


type alias Internals =
    { bio : Maybe String
    , avatar : Avatar
    }



-- INFO


bio : Profile -> Maybe String
bio (Profile info) =
    info.bio


avatar : Profile -> Avatar
avatar (Profile info) =
    info.avatar
