# Aliens

Note: To run programs you should have internet connection.
Domain `https://randomuser.me` is used for random alien and city names.

## Random map generation

To create new map for game run:
```
./run_generate.sh <limit_of_cities>
```

Example:
```
./run_generate.sh 50
```

Limit of cities will set limit for random generation of new cities.
Algorithm will try to randomly pick spot for new city, 
but will not guarantee creation if the same spot is picked twice.

## Starting the program

To start new game run:
```
./run_app.sh <number_of_aliens>
```

Example:
```
./run_app.sh 10
```

Number of aliens will determine how many aliens will play in game.

## Assumptions

- Number of cities should be greater than 0
- Number of aliens should be greater than 0
- Number of aliens has to be equal of less than number of cities
- Line should contain at least city name and one connection

## Example outputs

Map creation
```
Will try to create cities until reach limit: 10
0 - Success. Added new city!
1 - Success. Added new city!
2 - Success. Added new city!
3 - Success. Added new city!
4 - Success. Added new city!
5 - Success. Added new city!
6 - Success. Added new city!
7 - Fail. Randomly picked chosen position.
8 - Success. Added new city!
9 - Success. Added new city!
Map generated!
```

Running the game
```
Number of aliens: 4
Random alien name: Luke
Random city: Salem
City that are left: [Orange ChulaVista Oxnard Downey ThousandOaks Visalia Antioch Bakersfield Aubrey]
Random alien name: Nicholas
Random city: Antioch
City that are left: [Orange ChulaVista Oxnard Downey ThousandOaks Visalia Bakersfield Aubrey]
Random alien name: Christine
Random city: Aubrey
City that are left: [Orange ChulaVista Oxnard Downey ThousandOaks Visalia Bakersfield]
Random alien name: Chester
Random city: Orange
City that are left: [ChulaVista Oxnard Downey ThousandOaks Visalia Bakersfield]

Starting the game with: 4 aliens

Round: 1
Moving alien Christine from city Aubrey to city Antioch
City Antioch was destroyed by alien Christine and alien Nicholas
Moving alien Chester from city Orange to city Oxnard
Alien Chester was moved to city Oxnard
Moving alien Luke from city Salem to city Downey
Alien Luke was moved to city Downey
Round: 2
Moving alien Chester from city Oxnard to city ChulaVista
Alien Chester was moved to city ChulaVista
Moving alien Luke from city Downey to city Salem
Alien Luke was moved to city Salem
Round: 3
Moving alien Chester from city ChulaVista to city Oxnard
Alien Chester was moved to city Oxnard
Moving alien Luke from city Salem to city Downey
Alien Luke was moved to city Downey
Round: 4
Moving alien Luke from city Downey to city Bakersfield
Alien Luke was moved to city Bakersfield
Moving alien Chester from city Oxnard to city Orange
Alien Chester was moved to city Orange
Round: 5
Moving alien Chester from city Orange to city Oxnard
Alien Chester was moved to city Oxnard
Moving alien Luke from city Bakersfield to city Oxnard
City Oxnard was destroyed by alien Luke and alien Chester
All aliens were destroyed
```