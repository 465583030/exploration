package prize

import (
    "github.com/nienie/exploration/boundedcontext/usercontext"
    "github.com/nienie/exploration/boundedcontext/task"
)

//PrizeService ...
type PrizeService struct {}

//Award 发奖
func (s *PrizeService)Award(ctx *usercontext.UserContext, task *task.Task) error {
    return nil
}

//Accept 领奖
func (s *PrizeService)Accept(ctx *usercontext.UserContext, task *task.Task) error {
    return nil
}

//GetAcceptablePrizes 获取为领取的奖品
func (s *PrizeService)GetAcceptablePrizes(ctx *usercontext.UserContext) ([]*UserPrize, error) {
    return nil, nil
}