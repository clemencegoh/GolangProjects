package main

import "sync"

type ResumeDetails struct{
	Name string `json:"Name"`
	CurrentJobTitle string `json:"CurrentJobTitle"`
	CurrentJobDescription string `json:"CurrentJobDescription"`
	CurrentJobCompany string `json:"CurrentJobCompany"`
	ResumeID int `json:"ResumeId"`
}

// returns a json version
func (rd *ResumeDetails) toString() string{
	return toJson(rd)
}

func (rd *ResumeDetails) compareID(resumeID int) bool{
	if rd.ResumeID==resumeID{
		return true
	}
	return false
}

// guarded by mu
type IDCounter struct{
	mu sync.Mutex
	ID int
}

func (i *IDCounter) GetAndIncrement() int{
	// synchronize
	i.mu.Lock()
	x := i.ID
	i.ID += 1
	i.mu.Unlock()
	return x
}