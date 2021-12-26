package database

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

type Vote struct {
	Person string //ID
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

var PollTypes = map[string]*PollType{}
var activePolls = map[int]*Poll{}

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
	votes := get_votes(user, poll)
	if votes == 0 {
		return errors.New("Not allowed to vote in the poll")
	}
	poll.Voted = append(poll.Voted, Vote{user.User.ID, vote})
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

func New_Poll() {

}
