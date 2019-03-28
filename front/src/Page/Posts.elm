module Page.Posts exposing (Model, Msg(..), init, postsView, update, view, toSession)

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
    { title = "Posts"
    , content = postsView model
    }


postsView : Model -> Html.Styled.Html Msg
postsView model =
    Ui.card []
        [ Ui.heading "New Post"
        , Ui.linkBtn [ Route.href Route.Home ] [ text "Home" ]
        , Ui.paragraph [] [ text "This is where new posts would go?" ]
        ]


toSession : Model -> Session
toSession { session } =
    session
