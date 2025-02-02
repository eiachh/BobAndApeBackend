package types

type MobName string

const (
	SpawnMobCommandName CommandName = "SpawnMobCommand"
)

var (
	BossDefaultX int = 1000
	BossDefaultY int = 500
)

var (
	HarambeSpawn = SpawnMobCommand{
		MobName:   "harambe",
		MobHealth: 100000,
		PosX:      BossDefaultX,
		PosY:      BossDefaultY,
	}
)

type SpawnMobCommand struct {
	MobName    string `json:"mobname"`
	MobHealth  int    `json:"health"`
	PosX       int    `json:"posx"`
	PosY       int    `json:"posy"`
	Invincible bool   `json:"invincible"`
}
