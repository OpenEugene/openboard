module Main exposing (Model, Msg(..), init, linkBtn, main, theme, update, view)

import Browser
import Css exposing (..)
import Html.Attributes exposing (src)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Html.Styled.Events exposing (onClick)



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
    div []
        [ img [ src "/logo.svg" ] []
        , h1 [] [ text "Your Elm App is working! And hot reloading!" ]
        , linkBtn [ href "https://package.elm-lang.org/packages/rtfeldman/elm-css/latest" ] [ text "Elm CSS sounds fun" ]
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



-- Elm CSS --


{-| A reusable button which has some styles pre-applied to it.
-}
linkBtn : List (Attribute msg) -> List (Html msg) -> Html msg
linkBtn =
    styled a
        [ padding4 (px 8) (px 16) (px 8) (px 16)
        , color (rgb 250 250 250)
        , hover
            [ backgroundColor theme.hover
            , textDecoration underline
            ]
        , backgroundColor theme.primary
        , cursor pointer
        , border (px 0)
        , borderRadius (px 3)
        , fontSize (Css.em 1)
        ]


{-| A plain old record holding a couple of theme colors.
-}
theme =
    { primary = hex "55af6a"
    , secondary = rgb 250 240 230
    , hover = hex "3ebc5b"
    }
