import Html exposing (...)
import Svg exposing (...)
import Svg.Attributes exposing (...)

-- hello
main = 
    Html.text "Hello, world!"

-- list
list = 
    div []
        [ h1 [] [ text "Hello, world!" ]
        , ul []
            [ li [] [ text "foo" ]
            , li [] [ text "bar" ]
            , li [] [ text "baz" ]
            , li [] [ text "quux" ]
            , li [] [ text "quuux" ]
            , li [] [ text "quuuux" ]
            ]
        ]

-- shape

shape = 
    Svg
        [ viewbox "0 0 100 100"
        , width "100"
        , height "100"
        ]
        [ g []
            [ rect []
                [ x "0"
                , y "0"
                , width "100"
                , height "100"
                , fill "red"
                ]
            , rect []
                [ x "50"
                , y "50"
                , width "50"
                , height "50"
                , fill "green"
                ]
            , circle []
                [ cx "50"
                , cy "50"
                , r "50"
                , fill "blue"
                ]
            ]
            , line []
                [ x1 "0"
                , y1 "0"
                , x2 "100"
                , y2 "100"
                , stroke "black"
                ]
            , polygon []
                [ points "0,0 50,0 100,100 0,100"
                , fill "black"
                ]
            , text_ []
                [ x "50"
                , y "50"
                , fill "white"
                ]
                [ text "Hello, world!" ]
            ]
        ]
