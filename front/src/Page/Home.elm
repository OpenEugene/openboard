module Page.Home exposing (Model, Msg, init, subscriptions, toSession, update, view)

{-| The homepage. You can get here via either the / or /#/ routes.
-}

import Html exposing (..)
import Session exposing (Session)
import Time





-- MODEL


type alias Model =
    { session : Session
    , timeZone : Time.Zone
    }



init : Session -> ( Model, Cmd Msg )
init session =
    ( { session = session
      , timeZone = Time.utc
      }
    , Cmd.none
    )



-- VIEW


view : Model -> { title : String, content : Html Msg }
view model =
    { title = "Openboard"
    , content = text "home"
    }


-- UPDATE


type Msg
    = Noop
    | GotSession Session


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
            ( model
            , Cmd.none
            )



subscriptions : Model -> Sub Msg
subscriptions model =
    Session.changes GotSession (Session.navKey model.session)



-- EXPORT


toSession : Model -> Session
toSession model =
    model.session