module Main exposing (Model, Msg(..), init, main, update, view)

import Browser
import Css exposing (..)

import Html.Styled exposing (toUnstyled, Html)
import Proto.User
import Ui
import Route exposing (Route(..))
import Json.Decode
import Page.Home as Home
import Session exposing (Session)
import Viewer exposing (Viewer)
import Url exposing (Url)
import Browser.Navigation as Nav
import Browser exposing (Document)

test : Proto.User.RoleResp
test =
    { id = 0
    , name = "hello"
    }



---- MODEL ----


type Model
    = Redirect Session
    | NotFound Session
    | Home Home.Model
    | Login



init : Json.Decode.Value -> Url -> Nav.Key -> ( Model, Cmd Msg )
init maybeViewer url navKey =
    let
--        TODO: decode the incoming js value
        maybeViewer_ = Nothing
    in
    changeRouteTo (Route.fromUrl url)
        (Redirect (Session.fromViewer navKey maybeViewer_))


changeRouteTo : Maybe Route -> Model -> ( Model, Cmd Msg )
changeRouteTo maybeRoute model =
    let
        session =
            toSession model
    in
    case maybeRoute of
        Nothing ->
            ( NotFound session, Cmd.none )

        Just Route.Root ->
            ( model, Route.replaceUrl (Session.navKey session) Route.Home )

        Just Route.Home ->
            Home.init session
                |> updateWith Home GotHomeMsg model

        Just Route.Login ->
            Home.init session
                |> updateWith Home GotHomeMsg model




toSession : Model -> Session
toSession page =
    case page of
        Redirect session ->
            session

        NotFound session ->
            session

        Home home ->
            Home.toSession home

        Login ->
            Debug.todo "handle login case"



---- UPDATE ----

type Msg
    = Ignored
    | ChangedRoute (Maybe Route)
    | ChangedUrl Url
    | ClickedLink Browser.UrlRequest
    | GotHomeMsg Home.Msg



update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case ( msg, model ) of
        ( Ignored, _ ) ->
            ( model, Cmd.none )

        ( ClickedLink urlRequest, _ ) ->
            case urlRequest of
                Browser.Internal url ->
                    case url.fragment of
                        Nothing ->
                            -- If we got a link that didn't include a fragment,
                            -- it's from one of those (href "") attributes that
                            -- we have to include to make the RealWorld CSS work.
                            --
                            -- In an application doing path routing instead of
                            -- fragment-based routing, this entire
                            -- `case url.fragment of` expression this comment
                            -- is inside would be unnecessary.
                            ( model, Cmd.none )

                        Just _ ->
                            ( model
                            , Nav.pushUrl (Session.navKey (toSession model)) (Url.toString url)
                            )

                Browser.External href ->
                    ( model
                    , Nav.load href
                    )

        ( ChangedUrl url, _ ) ->
            changeRouteTo (Route.fromUrl url) model

        ( ChangedRoute route, _ ) ->
            changeRouteTo route model


        ( GotHomeMsg subMsg, Home home ) ->
            Home.update subMsg home
                |> updateWith Home GotHomeMsg model


        ( _, _ ) ->
            -- Disregard messages that arrived for the wrong page.
            ( model, Cmd.none )

updateWith : (subModel -> Model) -> (subMsg -> Msg) -> Model -> ( subModel, Cmd subMsg ) -> ( Model, Cmd Msg )
updateWith toModel toMsg model ( subModel, subCmd ) =
    ( toModel subModel
    , Cmd.map toMsg subCmd
    )

---- VIEW ----


view : Model -> Browser.Document Msg
view model =
    {
    title = "Openboard",
    body = List.map toUnstyled [Ui.card []
                      [ Ui.heading  "openboard!"
                      ]]
    }




---- PROGRAM ----


main : Program Json.Decode.Value Model Msg
main =
    Browser.application
        { init = init
        , onUrlChange = ChangedUrl
        , onUrlRequest = ClickedLink
        , subscriptions = always Sub.none
        , update = update
        , view = view
        }