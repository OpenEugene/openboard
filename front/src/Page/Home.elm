module Page.Home exposing (Model, Msg(..), init, toSession, update, view)

import Html.Styled exposing (text)
import Html.Styled.Events exposing (onClick)
import Route
import Session exposing (Session)
import Ui


type alias Model =
    { session : Session
    , greeting : String
    }


type Msg
    = InternalHomeMsg


init : Session -> ( Model, Cmd Msg )
init s =
    ( Model s "Change me", Cmd.none )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        InternalHomeMsg ->
            ( { model | greeting = "Hello from Page.Home" }, Cmd.none )


view : Model -> { title : String, content : Html.Styled.Html Msg }
view model =
    { title = "Home"
    , content = homeView model
    }


homeView : Model -> Html.Styled.Html Msg
homeView model =
    Ui.card []
        [ Ui.heading "Home"
        , Ui.linkBtn [ Route.href Route.Login ] [ text "Login" ]
        , Ui.paragraph [] [ text model.greeting ]
        , Ui.btn [ onClick InternalHomeMsg ] [ text "Click me to update the home page" ]
        ]


toSession : Model -> Session
toSession { session } =
    session
