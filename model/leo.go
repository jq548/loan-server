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
