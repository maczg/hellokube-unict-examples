package main

import "encoding/json"

type Vote struct {
	VoterID string `json:"voter_id"`
	Vote    string `json:"vote"`
}

func parseVote(vote []byte) (*Vote, error) {
	v := &Vote{}
	err := json.Unmarshal(vote, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
