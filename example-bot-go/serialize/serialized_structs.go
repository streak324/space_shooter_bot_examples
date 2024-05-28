package serialize

type Vec2 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Entity struct {
	My    bool   `json:"my"`
	Id    uint64 `json:"id"`
	Owner int32  `json:"owner"`
}

type GameState struct {
	Entities []Entity `json:"entities"`
}
