# proxx_go_test
GO test application, store and compute game field, proxx game


Vendor Take Home Interview Questions

Proxx

Rules and playable game. Review the rules and familiarize yourself with the game. You don’t
need to implement the flag functionality.
There are three parts to the exercise. For each part, please include a working coded solution
along with an explanation for choosing a certain approach.
Part 1:
Choose a data structure(s) to represent the game state. You need to keep track of the following:
- NxN board
- Location of black holes
- Counts of # of adjacent black holes
- Whether a cell is open

Part 2:
Populate your data structure with K black holes placed in random locations. Note that should
place exactly K black holes and their location should be random with a uniform distribution.

Part 3
For each cell without a black hole, compute and store the number of adjacent black holes. Note
that diagonals also count. E.g.

0 2 H
1 3 H
H 2 1
