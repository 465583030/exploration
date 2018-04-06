package completion

import "github.com/nienie/exploration/model/dao"

//Prize ...
type Prize struct {
    PrizeDao *dao.PrizeDao
}

//TaskPrize ...
type TaskPrize struct {
    TaskID      int64
    Prizes      []*Prize
    TaskPrizeDao *dao.TaskPrizeDao
}