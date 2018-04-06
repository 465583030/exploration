package task

import (
    "github.com/nienie/exploration/boundcontext/usercontext"
)

//Rule ...
type Rule struct {
    Precondition     *Rule          `json:"precondition"`
    Subject          string         `json:"subject"`
    Noun             string         `json:"noun"`
    Verb             string         `json:"verb"`
    Objects          []*Rule        `json:"objects"`
    Desc             string         `json:"desc"`
    ExpectedVal      []int64        `json:"expected_val"`
    CurrentVal       int64          `json:"current_val"`
}

//Condition ...
type Condition struct {
    Requirement       *Rule
}

//Checker ...
type Checker struct {}

//Allow ...
func (c *Checker)Allow(t *Task, ctx *usercontext.UserContext) bool {
    return true
}

