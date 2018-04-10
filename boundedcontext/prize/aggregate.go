package prize

//UserPrizeSet ...
type UserPrizeSet struct {
    UserID     int64
    UserPrizes []*UserPrize
}

//GetAcceptablePrizes ...
func (o *UserPrizeSet)GetAcceptablePrizes() ([]*UserPrize, error) {
    return nil, nil
}