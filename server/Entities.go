package server

const (
	AreaEffectCloud = iota
	ArmorStand
	Arrow
	Axolotl
	Bat
	Bee
	Blaze
	Boat
	Cat
	CaveSpider
	Chicken
	Cod
	Cow
	Creeper
	Dolphin
	Donkey
	DragonFireball
	Drowned
	ElderGuardian
	EndCrystal
	EnderDragon
	Enderman
	Endermite
	Evoker
	EvokerFangs
	ExperienceOrb
	EyeofEnder
	FallingBlock
	FireworkRocketEntity
	Fox
	Ghast
	Giant
	GlowItemFrame
	GlowSquid
	Goat
	Guardian
	Hoglin
	Horse
	Husk
	Illusioner
	IronGolem
	Item
	ItemFrame
	Fireball
	LeashKnot
	LightningBolt
	Llama
	LlamaSpit
	MagmaCube
	Marker
	Minecart
	MinecartChest
	MinecraftCommandBlock
	MinecraftFurnace
	MinecartHopper
	MinecartSpawner
	MinecraftTNT
	Mule
	Mooshroom
	Ocelot
	Painting
	Panda
	Parrot
	Phantom
	Pig
	Piglin
	PiglinBrute
	Pillager
	PolarBear
	PrimedTNT
	Pufferfish
	Rabbit
	Salmon
	Sheep
	Shulker
	ShulkerBullet
	Silverfish
	Skeleton
	SkeletonHorse
	Slime
	SmallFireball
	SnowGolem
	Snowball
	SpectralArrow
	Spider
	Squid
	Stray
	Strider
	ThrownEgg
	ThrownEnderPearl
	ThrownExperienceBottle
	ThrownPotion
	ThrownTrident
	TraderLlama
	TropicalFish
	Turtle
	Vex
	Villager
	Vindicator
	WanderingTrader
	Witch
	Wither
	WitherSkull
	Wolf
	Zoglin
	Zombie
	ZombieHorse
	ZombieVillager
	ZombifiedPiglin
	Player
	FishingHook
)

// EntityMetaDataType
const (
	EMDTByte = iota
	EMDTVarInt
	EMDTFloat
	EMDTString
	EMDTChat
	EMDTOptChat
	EMDTSlot
	EMDTBoolean
	EMDTRotation
	EMDTPosition
	EMDTOptPosition
	EMDTDirection // 0:DOWN 1:UP 2:NORTH 3:SOUTH 4:WEST 5:EAST
	EMDTOptUUID
	EMDTBlockID
	EMDTNBT
	EMDTParticle
	EMDTVillagerData
	EMDTOptVarInt
	EMDTPose
)

const (
	DOWN = iota
	UP
	NORTH
	SOUTH
	WEST
	EAST
)
