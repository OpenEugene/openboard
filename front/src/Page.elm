module Page exposing (Page(..), view)

import Browser
import Html.Styled exposing (..)
import Viewer exposing (Viewer)
import Route
import Ui

type Page
    = Home
    | Other



-- Wraps the Pages


view : Maybe Viewer -> Page -> { title : String, content : Html.Styled.Html msg } -> Browser.Document msg
view maybeViewer page { title, content } =
    { title = title ++ " - Openboard"
    , body = Ui.globalStyle :: navBar page maybeViewer :: content :: [] |> List.map Html.Styled.toUnstyled
    }


navBar page maybeViewer =
    Ui.navBar [] [
        Ui.navBarList [] [
            li [] [
                Ui.linkBtn [ Route.href Route.Home ] [ text "OpenBoard" ]
            ]
        ]
    ]
    



viewerAsString : Maybe Viewer -> String
viewerAsString maybeV =
    case maybeV of
        Just v ->
            Viewer.unbox v

        Nothing ->
            "no viewer"
