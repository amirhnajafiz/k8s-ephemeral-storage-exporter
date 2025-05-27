package runner

// summary represents the structure of the JSON response from the kubelet summary API.
type summary struct {
	Node struct {
		NodeName string `json:"nodeName"`
	} `json:"node"`
	Pods []struct {
		PodRef struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
			UID       string `json:"uid"`
		} `json:"podRef"`
		Containers []struct {
			Name   string `json:"name"`
			Memory struct {
				UsageBytes uint64 `json:"usageBytes"`
			} `json:"memory"`
			Rootfs struct {
				UsedBytes uint64 `json:"usedBytes"`
			} `json:"rootfs"`
			Logs struct {
				UsedBytes uint64 `json:"usedBytes"`
			} `json:"logs"`
		} `json:"containers"`
		EphemeralStorage struct {
			UsedBytes uint64 `json:"usedBytes"`
		} `json:"ephemeral-storage"`
	} `json:"pods"`
}
