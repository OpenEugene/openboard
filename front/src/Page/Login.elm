module Page.Login exposing (Model, Msg(..), init, toSession, update, view)

import Html.Styled exposing (text)
import Route
import Session exposing (Session)
import Ui


type alias Model =
    { session : Session
    }


type Msg
    = Noop


init : Session -> ( Model, Cmd Msg )
init s =
    ( Model s, Cmd.none )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        Noop ->
            ( model, Cmd.none )


view : Model -> { title : String, content : Html.Styled.Html Msg }
view model =
    { title = "Login"
    , content = loginView model
    }


loginView : Model -> Html.Styled.Html Msg
loginView model =
    Ui.card []
        [ Ui.heading "Login"
        , Ui.linkBtn [ Route.href Route.Home ] [ text "Home" ]
        , Ui.paragraph [] [ text <| Debug.toString model.session ]
        ]


toSession : Model -> Session
toSession { session } =
    session
