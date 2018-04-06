package execution

import "time"

//RuleVal ...
type RuleVal struct {
    Subject          string         `json:"subject"`
    Noun             string         `json:"noun"`
    Verb             string         `json:"verb"`
    Objects          []*RuleVal     `json:"objects"`
    Desc             string         `json:"desc"`
    ExpectedVal      []int64        `json:"expected_val"`
    CurrentVal       int64          `json:"current_val"`
    ModifyTime       time.Time      `json:"modify_time"`
}