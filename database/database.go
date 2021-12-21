package database

import "errors"

type Vote struct {
	Person string //ID
	Vote   bool   //False -> against , True -> for
}

type Poll struct {
	Id                int
	T                 string //Type
	Author            string //Nickname
	AllowedRole       string //ID
	Voted             []Vote
	TotalVotesFor     int
	TotalVotesAgainst int
	NeededVotes       int
}

var activePolls = map[int]*Poll{}

func User_Vote(user_id string, poll_id int, vote int, votes int) error {
	poll, exists := activePolls[poll_id]
	if !exists {
		return errors.New("Poll does not exist")
	}
	poll.Voted = append(poll.Voted, Vote{user_id, vote})
	if vote {

	}
	return nil
}

func Check_Finished() []*Poll {
	finished_polls := make([]*Poll, 0)
	for id, poll := range activePolls {
		if poll.TotalVotes >= poll.NeededVotes {
			finished_polls = append(finished_polls, poll)
		}
	}
	return finished_polls
}
