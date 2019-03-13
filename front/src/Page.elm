module Page exposing (Page(..), view)

import Browser
import Html.Styled
import Viewer exposing (Viewer)


type Page
    = Home
    | Other



-- Wraps the Pages


view : Maybe Viewer -> Page -> { title : String, content : Html.Styled.Html msg } -> Browser.Document msg
view maybeViewer page { title, content } =
    { title = title ++ " - Openboard"
    , body = nav page maybeViewer :: content :: [] |> List.map Html.Styled.toUnstyled
    }


nav page maybeViewer =
    Html.Styled.text ("Top Level Stuff: " ++ Debug.toString page ++ " and " ++ viewerAsString maybeViewer)


viewerAsString : Maybe Viewer -> String
viewerAsString maybeV =
    case maybeV of
        Just v ->
            Viewer.unbox v

        Nothing ->
            "no viewer"
