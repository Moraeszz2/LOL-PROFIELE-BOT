package lol

type profileStruct struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	Name          string `json: "name"`
	SummonerLevel int    `json: "summonerLevel"`
}
