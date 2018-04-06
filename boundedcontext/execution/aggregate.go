package execution

//Progress ...
type Progress struct {
    ProcessID   int64
    TaskID      int64
    UserID      int64
    ProgressDao *ProgressDao
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