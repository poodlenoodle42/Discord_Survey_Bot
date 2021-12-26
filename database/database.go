package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Vote struct {
	Person string //Username
	Vote   bool   //False -> against , True -> for
}

type Poll struct {
	Id                int
	T                 *PollType //Type
	Author            string    //Nickname
	AllowedRole       string    //ID
	Voted             []Vote
	TotalVotesFor     int
	TotalVotesAgainst int
	NeededVotes       int
}

type PollType struct {
	Name  string
	Votes map[string]int //Mappes Role ID to Votes
}

var PollTypes map[string]*PollType
var activePolls map[int]*Poll //ID to polls
var logfile os.File

func get_votes(user *discordgo.Member, poll *Poll) int {
	votes := 0
	has_needed_role := false
	for _, r := range user.Roles {
		vote, ex := poll.T.Votes[r]
		if ex {
			votes += vote
		}
		if r == poll.AllowedRole {
			has_needed_role = true
		}
	}
	if has_needed_role {
		return votes
	} else {
		return 0
	}
}

func User_Vote(user *discordgo.Member, poll_id int, vote bool) error {
	poll, exists := activePolls[poll_id]
	if !exists {
		return errors.New("Poll does not exist")
	}

	for _, v := range poll.Voted {
		if v.Person == user.User.Username {
			return errors.New("Already voted")
		}
	}

	votes := get_votes(user, poll)
	if votes == 0 {
		return errors.New("Not allowed to vote in the poll")
	}
	poll.Voted = append(poll.Voted, Vote{user.User.Username, vote})
	if vote {
		poll.TotalVotesFor += votes
	} else {
		poll.TotalVotesAgainst += votes
	}
	return nil
}

func Check_Finished() []*Poll {
	finished_polls := make([]*Poll, 0)
	for id, poll := range activePolls {
		if poll.TotalVotesFor >= poll.NeededVotes || poll.TotalVotesAgainst >= poll.NeededVotes {
			finished_polls = append(finished_polls, poll)
		}
		delete(activePolls, id)
	}
	return finished_polls
}

func New_Poll(username string, message string, role string) (*Poll, error) {
	var poll_type string
	var needed_votes int

	scaned, err := fmt.Scanf("$new %s %d", &poll_type, &needed_votes)
	if scaned != 2 {
		return nil, err
	}
	t, ex := PollTypes[poll_type]
	if ex {
		return nil, errors.New("Poll type does not exist")
	}
	p := new(Poll)
	p.AllowedRole = role
	p.Author = username
	p.Id = rand.Int()
	p.NeededVotes = needed_votes
	p.TotalVotesAgainst = 0
	p.TotalVotesFor = 0
	p.Voted = make([]Vote, 0)
	p.T = t
	activePolls[p.Id] = p
	return p, nil
}

func Log(poll *Poll) {
	b, _ := json.Marshal(poll)
	logfile.Write(b)
}
