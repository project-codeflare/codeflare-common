package support

var (
	AMD    = Accelerator{Type: "gpu", ResourceLabel: "amd.com/gpu"}
	CPU    = Accelerator{Type: "cpu"}
	NVIDIA = Accelerator{Type: "gpu", ResourceLabel: "nvidia.com/gpu", PrometheusGpuUtilizationLabel: "DCGM_FI_DEV_GPU_UTIL"}
)

type Accelerator struct {
	Type                          string
	ResourceLabel                 string
	PrometheusGpuUtilizationLabel string
}

// Method to check if the accelerator is a GPU
func (a Accelerator) IsGpu() bool {
	return a != CPU
}
