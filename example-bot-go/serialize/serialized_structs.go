package serialize

type Entity struct {
	My    bool   `json:"my"`
	Id    uint64 `json:"id"`
	Owner int32  `json:"owner"`
}

type GameState struct {
	Entities []Entity `json:"entities"`
}
