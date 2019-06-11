module Route exposing (Route(..), fromUrl, href)

import Html.Styled exposing (Attribute)
import Html.Styled.Attributes as Attr
import Url exposing (Url)
import Url.Parser as Parser exposing ((</>), Parser, oneOf, s, string)



-- ROUTING


type Route
    = Home
    | Login
    | NewOffer
    | NewRequest
    | PostDetail String
    | EditPost String


parser : Parser (Route -> a) a
parser =
    oneOf
        [ Parser.map Home Parser.top
        , Parser.map Login (s "login")
        , Parser.map NewOffer (s "offer" </> s "new")
        , Parser.map NewRequest (s "request" </> s "new")
        , Parser.map PostDetail (s "posts" </> string)
        , Parser.map EditPost (s "posts" </> string </> s "edit")
        ]



-- PUBLIC HELPERS


href : Route -> Attribute msg
href targetRoute =
    Attr.href (routeToString targetRoute)


fromUrl : Url -> Maybe Route
fromUrl url =
    -- The RealWorld spec treats the fragment like a path.
    -- This makes it *literally* the path, so we can proceed
    -- with parsing as if it had been a normal path all along.
    { url | path = Maybe.withDefault "" url.fragment, fragment = Nothing }
        |> Parser.parse parser



-- INTERNAL


routeToString : Route -> String
routeToString page =
    let
        pieces =
            case page of
                Home ->
                    []

                Login ->
                    [ "login" ]

                NewRequest ->
                    [ "request", "new" ]

                NewOffer ->
                    [ "offer", "new" ]

                PostDetail string ->
                    [ "posts", string ]

                EditPost string ->
                    [ "posts", string, "edit" ]
    in
    "#/" ++ String.join "/" pieces
