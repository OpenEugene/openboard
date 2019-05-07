module Ui exposing (btn, card, heading, linkBtn, paragraph)

import Css exposing (..)
import Html.Styled exposing (..)



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
        , padding (px (gridUnit * 4))
        , boxShadow4 (px 1) (px 1) (px 5) theme.dark
        , margin (px (8 * gridUnit))
        ]


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
            [ backgroundColor theme.hover
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
        , fontSize (6 * gridUnit |> px)
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
        ]


{-| Theme colors
-}
theme =
    { primary = hex "55af6a"
    , secondary = rgb 250 240 230
    , hover = hex "3ebc5b"
    , dark = hex "999"
    }


{-| Grid unit to keep everything lined up and balanced
-}
gridUnit : Float
gridUnit =
    8
