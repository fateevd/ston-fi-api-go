package stonApi

type Asset struct {
	ContractAddress    string   `json:"contract_address"`
	Symbol             string   `json:"symbol"`
	DisplayName        string   `json:"display_name"`
	Priority           int      `json:"priority"`
	ImageUrl           string   `json:"image_url"`
	Decimals           int      `json:"decimals"`
	Kind               string   `json:"kind"` // "Ton" | "Wton" | "Jetton"
	Deprecated         bool     `json:"deprecated"`
	Community          bool     `json:"community"`
	Blacklisted        bool     `json:"blacklisted"`
	DefaultSymbol      bool     `json:"default_symbol"`
	Taxable            bool     `json:"taxable"`
	Tags               []string `json:"tags"`
	ThirdPartyUsdPrice string   `json:"third_party_usd_price"`
	ThirdPartyPriceUsd string   `json:"third_party_price_usd"`
	DexUsdPrice        string   `json:"dex_usd_price"`
	DexPriceUsd        string   `json:"dex_price_usd"`
}

type Reward struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type RewardDetail struct {
	Address          string `json:"address"`
	Status           string `json:"status"`
	RemainingRewards string `json:"remaining_rewards"`
	RewardRate24h    string `json:"reward_rate_24h"`
}

type NftInfo struct {
	Address             string   `json:"address"`
	CreateTimestamp     string   `json:"create_timestamp"`
	MinUnstakeTimestamp string   `json:"min_unstake_timestamp"`
	NonclaimedRewards   string   `json:"nonclaimed_rewards"`
	Rewards             []Reward `json:"rewards"`
	StakedTokens        string   `json:"staked_tokens"`
	Status              string   `json:"status"`
}

type Farm struct {
	Apy                *string        `json:"apy,omitempty"`
	MinStakeDurationS  string         `json:"min_stake_duration_s"`
	MinterAddress      string         `json:"minter_address"`
	NftInfos           []NftInfo      `json:"nft_infos"`
	PoolAddress        string         `json:"pool_address"`
	RewardTokenAddress string         `json:"reward_token_address"`
	Status             string         `json:"status"`
	Rewards            []RewardDetail `json:"rewards"`
}

type SwapPairs [2]string

type Pool struct {
	Address                    string `json:"address"`
	RouterAddress              string `json:"router_address"`
	Reserve0                   string `json:"reserve0"`
	Reserve1                   string `json:"reserve1"`
	Token0Address              string `json:"token0_address"`
	Token1Address              string `json:"token1_address"`
	LpTotalSupply              string `json:"lp_total_supply"`
	LpTotalSupplyUsd           string `json:"lp_total_supply_usd"`
	LpFee                      string `json:"lp_fee"`
	ProtocolFee                string `json:"protocol_fee"`
	RefFee                     string `json:"ref_fee"`
	ProtocolFeeAddress         string `json:"protocol_fee_address"`
	CollectedToken0ProtocolFee string `json:"collected_token0_protocol_fee"`
	CollectedToken1ProtocolFee string `json:"collected_token1_protocol_fee"`
	LpPriceUsd                 string `json:"lp_price_usd"`
	Apy1D                      string `json:"apy_1d"`
	Apy7D                      string `json:"apy_7d"`
	Apy30D                     string `json:"apy_30d"`
	Deprecated                 bool   `json:"deprecated"`
}
