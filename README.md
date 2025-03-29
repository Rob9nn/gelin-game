# Gelin Game

A kind of Olympic games app. Multiple teams confront each other during a day on a set of games.

## Getting started

SERVER
- go run cmd/server/main.go
CLI
- go run cmd/cli/main.go

# Road map

Go app storing data into a DB (postgres)
Most of the use is going to be on mobile. We must be able to load the application in IOS and Android. 
If there is connection problem, data must be saved locally on the device and then sync when ever its possible.

1. Create team 
  a. a team has a set of players
2. Create games
  a. a game has a goal (victory condition ?)
  b. a game has a max time ?
3. Journey -> Day ?
  a. set of games 
  b. set of teams

## Backend first
player(firstname, lastname, phone)
game(name, rule)
