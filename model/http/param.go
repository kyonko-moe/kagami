package http

// ArgRegister is.
type ArgRegister struct {
	Name string `form:"name" binding:"required"`
}

type RespRegister struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ArgLocate struct {
	Name string `form:"name" binding:"required"`
}

type RespLocate struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	IPv4 string `form:"ipv4"`
}

type ArgConnect struct {
	IPv4 string `form:"ipv4" binding:"required"`
}

type RespConnect struct {
	IPv4 string `json:"ipv4"`
	Port int    `json:"port"`
}
