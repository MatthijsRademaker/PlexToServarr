# Expected behavior

- watchlist gets requested and is [1,2]
  - [1,2] is in the requested list
  - [1] is in the watchlist for movies
  - [2] is in the watchlist for series
- watchlist gets requested and is [1,2]
  - [1,2] is ALREADY in the requested list
  - Nothing happens
- watchlist gets requested and is [1]
  - [1,2] is in the requested list
  - [1] is in the watchlist for movies
  - [2] is in the watchlist for series
  - remove [2] from the requested list and from the watchlist for series
