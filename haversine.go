package main

import (
    "fmt"
    "math"
)

func toRadians(degree float64) float64 {
    return degree * math.Pi / 180
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
		// Earth radius 
    const R = 6371

    phi1 := toRadians(lat1)
    phi2 := toRadians(lat2)
    deltaPhi := toRadians(lat2 - lat1)
    deltaLambda := toRadians(lon2 - lon1)

    a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
        math.Cos(phi1)*math.Cos(phi2)*
            math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    distance := R * c
		fmt.Println(distance)

    return distance
}

func isWithin5km(currentLat, currentLon, targetLat, targetLon float64) bool {
    distance := haversine(currentLat, currentLon, targetLat, targetLon)
    return distance <= 5
}

func main() {
    currentLat := -6.175110  // Koordinat Jakarta
    currentLon := 106.865036

    points := []struct {
        lat float64
        lon float64
    }{
        {-6.200000, 106.816666},
        {-6.300000, 106.800000},
        {-6.500000, 106.750000},
        {-6.150000, 106.850000},
        {-6.180000, 106.830000},
    }

    for i, point := range points {
        if isWithin5km(currentLat, currentLon, point.lat, point.lon) {
            fmt.Printf("Titik %d berada dalam radius 5 km.\n", i+1)
        } else {
            fmt.Printf("Titik %d berada di luar radius 5 km.\n", i+1)
        }
    }
}
