module Page.Posts exposing (Kind(..), Model, Msg(..), init, postsView, toSession, update, view)

import Html.Styled exposing (button, form, input, label, text, textarea)
import Html.Styled.Attributes exposing (type_)
import Html.Styled.Events exposing (onInput)
import Route
import Session exposing (Session)
import Ui


type alias Model =
    { session : Session
    , title : String
    , body : String
    , kind : Kind
    , slug : String
    }


type Kind
    = Offer
    | Request
    | Unknown


type Msg
    = SetTitle String
    | SetBody String


init : Session -> Kind -> String -> ( Model, Cmd Msg )
init session kind slug =
    ( Model session "" "" kind slug, Cmd.none )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        SetTitle t ->
            ( { model | title = t }, Cmd.none )

        SetBody b ->
            ( { model | body = b }, Cmd.none )


view : Model -> { title : String, content : Html.Styled.Html Msg }
view model =
    { title = "Posts"
    , content = postsView model
    }


heading : Kind -> Html.Styled.Html Msg
heading kind =
    case kind of
        Offer ->
            Ui.heading "New Offer"

        Request ->
            Ui.heading "New Request"

        Unknown ->
            Ui.heading "Impossible!!"


postsView : Model -> Html.Styled.Html Msg
postsView model =
    Ui.card []
        [ heading model.kind
        , Ui.linkBtn [ Route.href Route.Home ] [ text "Home" ]
        , form []
            [ label []
                [ text "Title"
                , input [ type_ "text", onInput SetTitle ] []
                ]
            , label []
                [ text "Body"
                , textarea [ onInput SetBody ] []
                ]
            , button [] [ text "submit" ]
            ]
        ]


toSession : Model -> Session
toSession { session } =
    session
