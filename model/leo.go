package model

type LeoTransaction struct {
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
					Type  string `json:"type"`
					ID    string `json:"id"`
					Value string `json:"value"`
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
}

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
			Network               int    `json:"network"`
			Round                 int    `json:"round"`
			Height                int    `json:"height"`
			CumulativeWeight      string `json:"cumulative_weight"`
			CumulativeProofTarget string `json:"cumulative_proof_target"`
			CoinbaseTarget        int    `json:"coinbase_target"`
			ProofTarget           int    `json:"proof_target"`
			LastCoinbaseTarget    int    `json:"last_coinbase_target"`
			LastCoinbaseTimestamp int    `json:"last_coinbase_timestamp"`
			Timestamp             int    `json:"timestamp"`
		} `json:"metadata"`
	} `json:"header"`
	Authority struct {
		Type   string `json:"type"`
		Subdag struct {
			Subdag struct {
				Num8518066 []struct {
					BatchHeader struct {
						BatchID                string        `json:"batch_id"`
						Author                 string        `json:"author"`
						Round                  int           `json:"round"`
						Timestamp              int           `json:"timestamp"`
						CommitteeID            string        `json:"committee_id"`
						TransmissionIds        []interface{} `json:"transmission_ids"`
						PreviousCertificateIds []string      `json:"previous_certificate_ids"`
						Signature              string        `json:"signature"`
					} `json:"batch_header"`
					Signatures []string `json:"signatures"`
				} `json:"8518066"`
				Num8518067 []struct {
					BatchHeader struct {
						BatchID                string        `json:"batch_id"`
						Author                 string        `json:"author"`
						Round                  int           `json:"round"`
						Timestamp              int           `json:"timestamp"`
						CommitteeID            string        `json:"committee_id"`
						TransmissionIds        []interface{} `json:"transmission_ids"`
						PreviousCertificateIds []string      `json:"previous_certificate_ids"`
						Signature              string        `json:"signature"`
					} `json:"batch_header"`
					Signatures []string `json:"signatures"`
				} `json:"8518067"`
				Num8518068 []struct {
					BatchHeader struct {
						BatchID                string        `json:"batch_id"`
						Author                 string        `json:"author"`
						Round                  int           `json:"round"`
						Timestamp              int           `json:"timestamp"`
						CommitteeID            string        `json:"committee_id"`
						TransmissionIds        []interface{} `json:"transmission_ids"`
						PreviousCertificateIds []string      `json:"previous_certificate_ids"`
						Signature              string        `json:"signature"`
					} `json:"batch_header"`
					Signatures []string `json:"signatures"`
				} `json:"8518068"`
			} `json:"subdag"`
		} `json:"subdag"`
	} `json:"authority"`
	Ratifications []struct {
		Type   string `json:"type"`
		Amount int    `json:"amount"`
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
						Type  string `json:"type"`
						ID    string `json:"id"`
						Value string `json:"value"`
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
	AleoAddress string `json:"aleo_address"`
	AleoAmount  int64  `json:"aleo_amount"`
	BscAddress  string `json:"bsc_address"`
	Email       string `json:"email"`
	Stages      int    `json:"stages"`
	DayPerStage int    `json:"day_per_stage"`
	Type        int    `json:"type"`
}

type SaveLoan struct {
	Role      int    `json:"role"`
	Count     int    `json:"count"`
	Address   string `json:"address"`
	Guarantor string `json:"guarantor"`
	Miner     string `json:"miner"`
}
