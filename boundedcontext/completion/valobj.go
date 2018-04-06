package completion

//Prize ...
type Prize struct {
    PrizeDao *PrizeDao
}

//TaskPrize ...
type TaskPrize struct {
    TaskID      int64
    Prizes      []*Prize
    TaskPrizeDao *TaskPrizeDao
}

//UserPrizeSet ...
type UserPrizeSet struct {
    UserID     int64
    UserPrizes []*UserPrize
}