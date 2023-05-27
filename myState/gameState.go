package myState

var CurrentGS GameState

type GameState int

const (
	FadeScreen GameState = iota
	StartScreen
	GoToScreen

	StageSelect
	TownScreen
	EquipmentScreen
	JobSelect
	SaveScreen

	PlayingScreen
	BattleEnemyScreen
	SkillScreen
	EndScreen
	TestState

	WeaponShop
	ArmorShop
	AccessoryShop
	BlackSmith
)

var CurrentBelong BelongState

type BelongState int

const (
	WeaponBelong BelongState = iota
	ArmorBelong
	AccessoryBelong
	MaterialsBelong
)
