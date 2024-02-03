package entity

type Account struct {
	Id        string `bson:"_id,omitempty" json:"id,omitemptys"`
	UserId    string `bson:"user_id,omitempty" json:"user_id,omitemptys"`
	AccountId string `bson:"account_id,omitempty" json:"account_id,omitemptys"`
}
