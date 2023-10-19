package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "time"
    "gopkg.in/mgo.v2/bson"
)

// Test types
type Game struct {
    Winner       string    `bson:"winner"`
    OfficialGame bool      `bson:"official_game"`
    Location     string    `bson:"location"`
    StartTime    time.Time `bson:"start"`
    EndTime      time.Time `bson:"end"`
    Players      []Player  `bson:"players"`
}

type Player struct {
    Name   string    `bson:"name"`
    Decks  [2]string `bson:"decks"`
    Points uint8     `bson:"points"`
    Place  uint8     `bson:"place"`
}

func NewPlayer(name, firstDeck, secondDeck string, points, place uint8) Player {
    return Player{
        Name:   name,
        Decks:  [2]string{firstDeck, secondDeck},
        Points: points,
        Place:  place,
    }
}

func main() {
    Host := []string{
        "localhost:27017",
    }
    const (
        // Username = "YOUR_USERNAME"
        // Password = "YOUR_PASSWORD"
        Database = "test"
        // ReplicaSetName = "c74b5276378ed3bd70cba37a3ac45fea"
        Collection = "test"
    )

    session, err := mgo.DialWithInfo(&mgo.DialInfo{
        Addrs:    Host,
        // Username: Username,
        // Password: Password,
        Database: Database,
        // ReplicaSetName: ReplicaSetName,
    })
    if err != nil {
        panic(err)
    }
    defer session.Close()

    fmt.Printf("Connected to replica set %v!\n", session.LiveServers())

    game := Game{
        Winner:       "Dave",
        OfficialGame: true,
        Location:     "Austin",
        StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
        EndTime:      time.Date(2015, time.February, 12, 05, 54, 0, 0, time.UTC),
        Players: []Player{
            NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
            NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
            NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
            NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
        },
    }

    coll := session.DB(Database).C(Collection)
    if err := coll.Insert(game); err != nil {
        panic(err)
    }
    fmt.Println("Document inserted successfully!")

    // Find the number of games won by Dave
    player := "Dave"
    gamesWon, err := coll.Find(bson.M{"winner": player}).Count()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s has won %d games.\n", player, gamesWon)

    // Remove all official games
    info, err := coll.RemoveAll(bson.M{"official_game": true})
    if err != nil {
        panic(err)
    }

    fmt.Printf("%d official game(s) removed!\n", info.Removed)
}