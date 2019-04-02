module Main exposing (Model, Msg(..), init, main, update, view)

import Browser exposing (Document, UrlRequest(..))
import Browser.Navigation as Nav
import Html
import Html.Styled
import Page
import Page.Home
import Page.Login
import Page.Posts
import Route exposing (Route(..))
import Session
import Url exposing (Url)



---- MODEL ----


type Model
    = Login Page.Login.Model
    | Home Page.Home.Model
    | Posts Page.Posts.Model



--TODO: decode incoming JWT from localStorage and initialize Session accordingly


init : () -> Url -> Nav.Key -> ( Model, Cmd Msg )
init _ url navKey =
    changeRouteTo (Route.fromUrl url) (Home (Page.Home.Model (Session.Guest navKey) ""))



---- UPDATE ----


type Msg
    = ChangedUrl Url
    | ClickedLink UrlRequest
    | GotHomeMsg Page.Home.Msg
    | GotLoginMsg Page.Login.Msg
    | GotPostsMsg Page.Posts.Msg


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case ( msg, model ) of
        ( ChangedUrl url, _ ) ->
            changeRouteTo (Route.fromUrl url) model

        ( ClickedLink urlRequest, _ ) ->
            case urlRequest of
                Internal url ->
                    ( model, Nav.pushUrl (Session.navKey (toSession model)) (Url.toString url) )

                External url ->
                    ( model, Nav.load url )

        ( GotHomeMsg subMsg, Home home ) ->
            Page.Home.update subMsg home
                |> updateWith Home GotHomeMsg model

        ( GotLoginMsg subMsg, Login home ) ->
            Page.Login.update subMsg home
                |> updateWith Login GotLoginMsg model

        ( GotPostsMsg subMsg, Posts home ) ->
            Page.Posts.update subMsg home
                |> updateWith Posts GotPostsMsg model

        ( _, _ ) ->
            -- Disregard messages that arrived for the wrong page.
            ( model, Cmd.none )


updateWith : (subModel -> Model) -> (subMsg -> Msg) -> Model -> ( subModel, Cmd subMsg ) -> ( Model, Cmd Msg )
updateWith toModel toMsg _ ( subModel, subCmd ) =
    ( toModel subModel
    , Cmd.map toMsg subCmd
    )


changeRouteTo : Maybe Route -> Model -> ( Model, Cmd Msg )
changeRouteTo maybeRoute model =
    case maybeRoute of
        Just Route.Home ->
            Page.Home.init (toSession model)
                |> updateWith Home GotHomeMsg model

        Just Route.Login ->
            Page.Login.init (toSession model)
                |> updateWith Login GotLoginMsg model

        Just Route.NewPost ->
            Page.Posts.init (toSession model) ""
                |> updateWith Posts GotPostsMsg model

        Just (Route.PostDetail slug) ->
            Page.Posts.init (toSession model) slug
                |> updateWith Posts GotPostsMsg model

        Just (Route.EditPost slug) ->
            Page.Posts.init (toSession model) slug
                |> updateWith Posts GotPostsMsg model

        Just Route.AllPosts ->
            Page.Posts.init (toSession model) ""
                |> updateWith Posts GotPostsMsg model

        Nothing ->
            ( model, Cmd.none )


toSession : Model -> Session.Session
toSession model =
    case model of
        Login loginModel ->
            Page.Login.toSession loginModel

        Home homeModel ->
            Page.Home.toSession homeModel

        Posts postsModel ->
            Page.Posts.toSession postsModel



---- VIEW ----


view : Model -> Document Msg
view model =
    let
        viewPage : Page.Page -> (msg -> Msg) -> { title : String, content : Html.Styled.Html msg } -> Document Msg
        viewPage page toMsg config =
            let
                { title, body } =
                    Page.view (Session.viewer (toSession model)) page config
            in
            { title = title
            , body = List.map (Html.map toMsg) body
            }
    in
    case model of
        Home home ->
            viewPage Page.Home GotHomeMsg (Page.Home.view home)

        Login login ->
            viewPage Page.Other GotLoginMsg (Page.Login.view login)

        Posts posts ->
            viewPage Page.Other GotPostsMsg (Page.Posts.view posts)



---- PROGRAM ----


main : Program () Model Msg
main =
    Browser.application
        { init = init
        , onUrlChange = ChangedUrl
        , onUrlRequest = ClickedLink
        , subscriptions = always Sub.none
        , update = update
        , view = view
        }
