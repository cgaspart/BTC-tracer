package wallet

type DWallet struct {
	Hash160       string `json:"hash160"`
	Address       string `json:"address"`
	NTx           int    `json:"n_tx"`
	NUnredeemed   int    `json:"n_unredeemed"`
	TotalReceived int    `json:"total_received"`
	TotalSent     int    `json:"total_sent"`
	FinalBalance  int    `json:"final_balance"`
	Txs           []struct {
		Hash        string      `json:"hash"`
		Ver         int         `json:"ver"`
		VinSz       int         `json:"vin_sz"`
		VoutSz      int         `json:"vout_sz"`
		Size        int         `json:"size"`
		Weight      int         `json:"weight"`
		Fee         int         `json:"fee"`
		RelayedBy   string      `json:"relayed_by"`
		LockTime    int         `json:"lock_time"`
		TxIndex     int64       `json:"tx_index"`
		DoubleSpend bool        `json:"double_spend"`
		Time        int         `json:"time"`
		BlockIndex  interface{} `json:"block_index"`
		BlockHeight interface{} `json:"block_height"`
		Inputs      []struct {
			Sequence int64  `json:"sequence"`
			Witness  string `json:"witness"`
			Script   string `json:"script"`
			Index    int    `json:"index"`
			PrevOut  struct {
				Spent             bool   `json:"spent"`
				Script            string `json:"script"`
				SpendingOutpoints []struct {
					TxIndex int64 `json:"tx_index"`
					N       int   `json:"n"`
				} `json:"spending_outpoints"`
				TxIndex int64  `json:"tx_index"`
				Value   int    `json:"value"`
				Addr    string `json:"addr"`
				N       int    `json:"n"`
				Type    int    `json:"type"`
			} `json:"prev_out"`
		} `json:"inputs"`
		Out []struct {
			Type              int           `json:"type"`
			Spent             bool          `json:"spent"`
			Value             int           `json:"value"`
			SpendingOutpoints []interface{} `json:"spending_outpoints"`
			N                 int           `json:"n"`
			TxIndex           int64         `json:"tx_index"`
			Script            string        `json:"script"`
			Addr              string        `json:"addr"`
		} `json:"out"`
		Result  int `json:"result"`
		Balance int `json:"balance"`
	} `json:"txs"`
}