module Page.Home exposing (Model, Msg(..), init, toSession, update, view)

import Html.Styled exposing (..)
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
    Ui.mainContent []
        [ Ui.flexBox []
            [ Ui.linkBtn [ Route.href Route.NewRequest ] [ text "Request" ]
            , Ui.linkBtn [ Route.href Route.NewOffer ] [ text "Offer" ]
            ]
        , Ui.postingsList []
            (List.map Ui.postingBlurb
                [ { slug = "1", title = "Golang Needed", body = "A bunch of text blablablaosdihf osidhf sdoifh sdfs df...." }
                , { slug = "2", title = "Another Posting Example", body = "lsjkdfno sidhsdf sdf sdfhj osd fsdf sdf sdf sdfmore some posting stuff" }
                ]
            )
        ]


toSession : Model -> Session
toSession { session } =
    session
