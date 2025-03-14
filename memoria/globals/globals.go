package globals

type MemoriaConfig struct {
	IpMemory   string   `json:"ip_memory"`
	PortMemory int      `json:"port_memory"`
	IpKernel   string   `json:"ip_kernel"`
	PortKernel int      `json:"port_kernel"`
	Port       int      `json:"port"`
	LogLevel   string   `json:"log_level"`
	Resources  []string `json:"resources"`
}

var ConfigMemoria *MemoriaConfig
