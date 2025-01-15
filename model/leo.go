package model

type LeoBlock struct {
	BlockHash    string `json:"block_hash"`
	PreviousHash string `json:"previous_hash"`
	Header       struct {
		PreviousStateRoot string `json:"previous_state_root"`
		TransactionsRoot  string `json:"transactions_root"`
		FinalizeRoot      string `json:"finalize_root"`
		RatificationsRoot string `json:"ratifications_root"`
		SolutionsRoot     string `json:"solutions_root"`
		SubdagRoot        string `json:"subdag_root"`
		Metadata          struct {
			Network               int `json:"network"`
			Round                 int `json:"round"`
			Height                int `json:"height"`
			CoinbaseTarget        int `json:"coinbase_target"`
			ProofTarget           int `json:"proof_target"`
			LastCoinbaseTarget    int `json:"last_coinbase_target"`
			LastCoinbaseTimestamp int `json:"last_coinbase_timestamp"`
			Timestamp             int `json:"timestamp"`
		} `json:"metadata"`
	} `json:"header"`
	Authority struct {
		Type      string `json:"type"`
		Signature string `json:"signature"`
	} `json:"authority"`
	Ratifications []struct {
		Type      string `json:"type"`
		Committee struct {
			ID            string `json:"id"`
			StartingRound int    `json:"starting_round"`
			TotalStake    int64  `json:"total_stake"`
		} `json:"committee"`
	} `json:"ratifications"`
	Solutions struct {
		Version int `json:"version"`
	} `json:"solutions"`
	AbortedSolutionIds []interface{} `json:"aborted_solution_ids"`
	Transactions       []struct {
		Status      string `json:"status"`
		Type        string `json:"type"`
		Index       int    `json:"index"`
		Transaction struct {
			Type      string `json:"type"`
			ID        string `json:"id"`
			Execution struct {
				Transitions []struct {
					ID       string `json:"id"`
					Program  string `json:"program"`
					Function string `json:"function"`
					Inputs   []struct {
						Type  string `json:"type"`
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"inputs"`
					Outputs []struct {
						Type     string `json:"type"`
						ID       string `json:"id"`
						Checksum string `json:"checksum,omitempty"`
						Value    string `json:"value"`
					} `json:"outputs"`
					Tpk string `json:"tpk"`
					Tcm string `json:"tcm"`
					Scm string `json:"scm"`
				} `json:"transitions"`
				GlobalStateRoot string `json:"global_state_root"`
				Proof           string `json:"proof"`
			} `json:"execution"`
			Fee struct {
				Transition struct {
					ID       string `json:"id"`
					Program  string `json:"program"`
					Function string `json:"function"`
					Inputs   []struct {
						Type  string `json:"type"`
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"inputs"`
					Outputs []struct {
						Type  string `json:"type"`
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"outputs"`
					Tpk string `json:"tpk"`
					Tcm string `json:"tcm"`
					Scm string `json:"scm"`
				} `json:"transition"`
				GlobalStateRoot string `json:"global_state_root"`
				Proof           string `json:"proof"`
			} `json:"fee"`
		} `json:"transaction"`
		Finalize []struct {
			Type      string `json:"type"`
			MappingID string `json:"mapping_id"`
			KeyID     string `json:"key_id"`
			ValueID   string `json:"value_id"`
		} `json:"finalize"`
	} `json:"transactions"`
	AbortedTransactionIds []interface{} `json:"aborted_transaction_ids"`
}

type ReqSaveDeposit struct {
	AleoAddress string  `json:"aleo_address"`
	AleoAmount  float64 `json:"aleo_amount"`
	BscAddress  string  `json:"bsc_address"`   //
	Email       string  `json:"email"`         //
	Stages      int     `json:"stages"`        //
	DayPerStage int     `json:"day_per_stage"` //
	LoanType    int     `json:"loan_type"`     // 1 loan , 2 pos
	Type        int     `json:"type"`          // 0 create, 1 add
	LoanId      int     `json:"loan_id"`       //
}

type SaveLoan struct {
	Role      int    `json:"role"`
	Count     int    `json:"count"`
	Address   string `json:"address"`
	Guarantor string `json:"guarantor"`
	Miner     string `json:"miner"`
}
