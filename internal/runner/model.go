package runner

type summary struct {
	Node struct {
		NodeName string `json:"nodeName"`
	} `json:"node"`
	Pods []struct {
		PodRef struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"podRef"`
		EphemeralStorage struct {
			UsedBytes uint64 `json:"usedBytes"`
		} `json:"ephemeral-storage"`
	} `json:"pods"`
}
