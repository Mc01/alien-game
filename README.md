# Aliens

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
Random alien name: Daniel
Random city: ThousandOaks
City that are left: [Orange Salem Visalia Antioch Oxnard Bakersfield Downey Aubrey ChulaVista]
Random alien name: Misty
Random city: Downey
City that are left: [Orange Salem Visalia Antioch Oxnard Bakersfield Aubrey ChulaVista]
Random alien name: Tim
Random city: ChulaVista
City that are left: [Orange Salem Visalia Antioch Oxnard Bakersfield Aubrey]
Random alien name: Freddie
Random city: Orange
City that are left: [Salem Visalia Antioch Oxnard Bakersfield Aubrey]
Aliens: map[Daniel:%!s(*main.Alien=&{Daniel 0xc0000c0400}) Freddie:%!s(*main.Alien=&{Freddie 0xc0000c0380}) Misty:%!s(*main.Alien=&{Misty 0xc0000c0340}) Tim:%!s(*main.Alien=&{Tim 0xc0000c0480})]
Moving alien Daniel from city ThousandOaks to city Bakersfield
Alien Daniel was moved to city Bakersfield
Moving alien Misty from city Downey to city Salem
Alien Misty was moved to city Salem
Moving alien Tim from city ChulaVista to city Oxnard
Alien Tim was moved to city Oxnard
Moving alien Freddie from city Orange to city Oxnard
City Oxnard was destroyed by alien Freddie and alien Tim
Moving alien Daniel from city Bakersfield to city Downey
Alien Daniel was moved to city Downey
Moving alien Misty from city Salem to city Downey
City Downey was destroyed by alien Misty and alien Daniel
All aliens were destroyed
```