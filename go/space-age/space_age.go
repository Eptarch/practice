package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    var scale float64
    switch planet {
        case "Mercury": scale = 0.2408467
        case "Venus": scale = 0.61519726
        case "Earth": scale = 1.0
        case "Mars": scale = 1.8808158
        case "Jupiter": scale = 11.862615
        case "Saturn": scale = 29.447498 
        case "Uranus": scale = 84.016846 
        case "Neptune": scale = 164.79132
    }
    return seconds / scale / 60 / 60 / 24 / 365.25 
}
