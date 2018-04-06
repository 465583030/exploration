package execution

import "github.com/nienie/exploration/model/dao"

//Progress ...
type Progress struct {
    ProcessID   int64
    TaskID      int64
    UserID      int64
    ProgressDao *dao.ProgressDao
    RuleVal     *RuleVal
}

//IsCompleted ...
func (p *Progress)IsCompleted() bool {
    return true
}

//Save ...
func (p *Progress)Save() error {
    return nil
}