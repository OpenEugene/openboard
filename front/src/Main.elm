module Main exposing (Model, Msg(..), init, main, update, view)

import Browser
import Css exposing (..)
import Html
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Html.Styled.Events exposing (onClick)
import Proto.User
import Ui


test : Proto.User.RoleResp
test =
    { id = 0
    , name = "hello"
    }



---- MODEL ----


type alias Model =
    {}


init : ( Model, Cmd Msg )
init =
    ( {}, Cmd.none )



---- UPDATE ----


type Msg
    = NoOp


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    ( model, Cmd.none )



---- VIEW ----


view : Model -> Html Msg
view model =
    Ui.card []
        [ Ui.heading [] [ text "Openboard!" ]
        , Ui.linkBtn [ href "https://github.com/rtfeldman/elm-css" ] [ text "Learn more about elm-css" ]
        , Ui.linkBtn [ href "https://github.com/champagneabuelo/openboard" ] [ text "the repo" ]
        ]



---- PROGRAM ----


main : Program () Model Msg
main =
    Browser.element
        { view = view >> toUnstyled
        , init = \_ -> init
        , update = update
        , subscriptions = always Sub.none
        }
