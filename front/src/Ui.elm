module Ui exposing (PostingKind(..), btn, card, flexBox, globalStyle, gridUnit, h1Style, heading, kindButton, kindToColor, linkBtn, mainContent, navBar, navBarList, pStyle, paragraph, postingBlurb, postingsList, theme)

import Css exposing (..)
import Css.Global exposing (global, selector)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (..)
import Route


globalStyle =
    global
        [ selector "body"
            [ margin (px 0)
            ]
        ]



-- Reusabled parts of the Ui used by Pages


paragraph =
    styled p [ pStyle ]


btn =
    styled button [ Css.batch [] ]


{-| A reusable card module
-}
card : List (Attribute msg) -> List (Html msg) -> Html msg
card =
    styled div
        [ borderRadius (px 4)
        , padding (px (gridUnit * 1))
        , boxShadow4 (px 1) (px 1) (px 5) theme.dark
        , margin (px (8 * gridUnit))
        ]


{-| navBar
-}
navBar : List (Attribute msg) -> List (Html msg) -> Html msg
navBar =
    styled nav
        [ boxShadow4 (px 1) (px 1) (px 5) theme.dark
        , displayFlex
        ]


navBarList : List (Attribute msg) -> List (Html msg) -> Html msg
navBarList =
    styled ul
        [ listStyleType none
        , margin (px 0)
        , padding (px 0)
        ]


mainContent : List (Attribute msg) -> List (Html msg) -> Html msg
mainContent =
    styled div
        [ maxWidth (px 960)
        , marginLeft auto
        , marginRight auto
        , padding (px 16)
        ]


type PostingKind
    = Request
    | Offer


kindToColor k =
    case k of
        Request ->
            theme.request

        Offer ->
            theme.offer


flexBox =
    styled div
        [ displayFlex
        , marginBottom (px 10)
        , firstChild [ marginRight (px 10) ]
        ]


kindButton : PostingKind -> List (Attribute msg) -> List (Html msg) -> Html msg
kindButton k =
    styled button
        [ padding4 (px 8) (px 16) (px 8) (px 16)
        , color (rgb 250 250 250)
        , hover
            [ backgroundColor theme.dark
            , textDecoration underline
            ]
        , display block
        , backgroundColor (kindToColor k)
        , cursor pointer
        , border (px 0)
        , borderRadius (px 3)
        , fontSize (Css.em 2)
        ]


postingsList =
    styled ul
        [ listStyleType none
        , margin (px 0)
        , padding (px 0)
        ]


postingBlurb { title, body, slug } =
    styled article
        [ borderRadius (px 4)
        , padding (px (gridUnit * 1))
        , boxShadow4 (px 1) (px 1) (px 5) theme.dark
        , marginBottom (px gridUnit)
        ]
        []
        [ a [ Route.PostDetail slug |> Route.href ] [ heading title ], paragraph [] [ text body ] ]


{-| A reusable heading
-}
heading : String -> Html msg
heading title =
    styled h1 [ h1Style ] [] [ text title ]


{-| A reusable link button
-}
linkBtn : List (Attribute msg) -> List (Html msg) -> Html msg
linkBtn =
    styled a
        [ padding4 (px 8) (px 16) (px 8) (px 16)
        , color (rgb 250 250 250)
        , hover
            [ backgroundColor theme.dark
            , textDecoration underline
            ]
        , display block
        , margin (px 10)
        , backgroundColor theme.primary
        , cursor pointer
        , border (px 0)
        , borderRadius (px 3)
        , fontSize (Css.em 1)
        ]


{-| Styles for h1
-}
h1Style : Style
h1Style =
    Css.batch
        [ fontFamilies [ "Palatino Linotype", "Georgia", "serif" ]
        , fontSize (3 * gridUnit |> px)
        , fontWeight bold
        , marginTop (px 0)
        ]


{-| Styles for h1
-}
pStyle : Style
pStyle =
    Css.batch
        [ fontFamilies [ "Helvetica", "Arial", "sans-serif", "serif" ]
        , fontSize (2 * gridUnit |> px)
        , fontWeight normal
        , lineHeight (2 * gridUnit + gridUnit |> px)
        , color theme.dark
        , margin (px 0)
        ]


{-| Theme colors
-}
theme =
    { primary = hex "000"
    , secondary = rgb 250 240 230
    , hover = hex "3ebc5b"
    , dark = hex "999"
    , request = hex "2135d1"
    , offer = hex "50b320"
    }


{-| Grid unit to keep everything lined up and balanced
-}
gridUnit : Float
gridUnit =
    8
