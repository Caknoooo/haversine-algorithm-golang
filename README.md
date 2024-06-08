# Simple Haversine Algorithm

*The haversine algorithm is an example of modeling calculation in the form of trigonometry, which is applied to a round shape. This algorithm discusses the shapes of sides and angles in spherical triangles. In 1805, a scientist created a Haversine table to determine a distance between points*

The Haversine algorithm is used to calculate the shortest distance between two points on the earth's surface, expressed in latitude and longitude coordinates. This algorithm is useful in navigation and geolocation to calculate "great-circle" distances measured along the surface of a sphere, such as the Earth.

![image](https://github.com/Caknoooo/haversine-algorithm-golang/assets/92671053/21cca654-d4ae-446e-8ff8-7ea8bc0bc415)

## Haversine Formula 
The Haversine formula is perhaps the first equation to consider when understanding how to calculate distances on a sphere. The word "Haversine" comes from the function:

```
haversine(θ) = sin²(θ/2)
```

The following equation where φ is latitude, λ is longitude, R is earth’s radius (mean radius = 6,371km) is how we translate the above formula to include latitude and longitude coordinates. Note that angles need to be in radians to pass to trig functions:

```
a = sin²(φB - φA/2) + cos φA * cos φB * sin²(λB - λA/2)
c = 2 * atan2( √a, √(1−a) )
d = R ⋅ c
```