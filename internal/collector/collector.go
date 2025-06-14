package collector

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/k8sese/internal/metrics"
	"github.com/amirhnajafiz/k8sese/pkg/fetch"
	"github.com/amirhnajafiz/k8sese/pkg/types"

	"go.uber.org/zap"
)

// Collector is responsible for collecting storage usage metrics from the kubelet summary endpoint
// and updating the provided metrics instance with the collected data.
type Collector struct {
	Logr     *zap.Logger
	Metrics  *metrics.Metrics
	Interval time.Duration
}

// Start initiates the process of fetching storage usage metrics from the kubelet summary endpoint
// and updates the provided metrics instance with the data.
func (c *Collector) Start(endpoint string) error {
	// build the HTTP request to the kubelet summary endpoint
	req, err := buildHTTPRequest(endpoint)
	if err != nil {
		return fmt.Errorf("failed to build HTTP request: %w", err)
	}

	c.Logr.Info(
		"starting kubelet summary collector",
		zap.String("endpoint", req.URL.String()),
		zap.Duration("interval", c.Interval),
	)

	for {
		// wait for the specified interval before fetching metrics
		time.Sleep(c.Interval)
		c.Logr.Debug("fetching kubelet summary for storage usage metrics")

		// perform the HTTP GET request
		resp, err := fetch.GET(req)
		if err != nil {
			c.Logr.Error("failed to fetch kubelet summary", zap.Error(err))
			continue
		}

		// decode the JSON response into a summary structure
		var summary types.Summary
		if err := fetch.JSON(resp, &summary); err != nil {
			c.Logr.Error("failed to decode kubelet summary JSON", zap.Error(err))
			continue
		}

		// process the summary data and update the metrics
		for _, pod := range summary.Pods {
			c.setPodStorageUsage(pod, summary.Node.NodeName)
			c.setContainerStorageUsage(pod, summary.Node.NodeName)
		}

		c.Logr.Debug("successfully updated storage usage metrics", zap.String("node", summary.Node.NodeName))
	}
}

// setPodStorageUsage sets the ephemeral storage usage for a pod in the provided metrics instance.
func (c *Collector) setPodStorageUsage(pod types.PodSummary, nodeName string) {
	c.Metrics.SetStorageUsageBytes(
		pod.PodRef.Name,
		pod.PodRef.Namespace,
		nodeName,
		float64(pod.EphemeralStorage.UsedBytes),
	)
}

// setContainerStorageUsage sets the storage usage for each container in a pod in the provided metrics instance.
func (c *Collector) setContainerStorageUsage(pod types.PodSummary, nodeName string) {
	for _, container := range pod.Containers {
		c.Metrics.SetContainerStorageUsageBytes(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			"memory",
			float64(container.Memory.UsageBytes),
		)
		c.Metrics.SetContainerStorageUsageBytes(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			"logs",
			float64(container.Logs.UsedBytes),
		)
		c.Metrics.SetContainerStorageUsageBytes(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			"rootfs",
			float64(container.Rootfs.UsedBytes),
		)
	}
}
