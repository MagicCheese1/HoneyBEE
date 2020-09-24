package blocks

var BlockID map[int]string

//Initialise BlockID map
func InitGlobalID() map[int]string {
	BlockID = map[int]string{
		0:   "Air",
		1:   "Stone",
		2:   "Grass",
		3:   "Dirt",
		4:   "Cobblestone",
		5:   "WoodenPlanks",
		6:   "Sapling",
		7:   "Bedrock",
		8:   "Water",
		9:   "WaterStill",
		10:  "Lava",
		11:  "LavaStill",
		12:  "Sand",
		13:  "Gravel",
		14:  "GoldOre",
		15:  "IronOre",
		16:  "CoalOre",
		17:  "Log",
		18:  "Leaves",
		19:  "Sponge",
		20:  "Glass",
		21:  "LapisLazuliOre",
		22:  "LapisLazuliBlock",
		23:  "Dispenser",
		24:  "Sandstone",
		25:  "NoteBlock",
		26:  "Bed",
		27:  "PoweredRail",
		28:  "DetectorRail",
		29:  "StickyPiston",
		30:  "Cobweb",
		31:  "TallGrass", //Includes Fern
		32:  "DeadBush",
		33:  "Piston",
		34:  "PistonHead",
		35:  "Wool", //Includes Coloured Varients
		36:  "MovedByPiston",
		37:  "Dandelion",
		38:  "Flower", //Includes: Oxeye Daisy; Pink, White, Orange, red tulip; Azure; Allium; Blue Orchid and poppy
		39:  "BrownMushroom",
		40:  "RedMushroom",
		41:  "GoldBlock",
		42:  "IronBlock",
		43:  "DoubleStoneSlab", //This includes petrified oakslab
		44:  "StoneSlab",       //This includes petrified oakslab
		45:  "Bricks",
		46:  "TNT",
		47:  "Bookshelf",
		48:  "MossStone",
		49:  "Obsidian",
		50:  "Torch",
		51:  "Fire",
		52:  "MonsterSpawner",
		53:  "OakStairs",
		54:  "Chest",
		55:  "RedstoneWire",
		56:  "DiamondOre",
		57:  "DiamondBlock",
		58:  "CraftingTable",
		59:  "Wheat",
		60:  "Farmland",
		61:  "Furnace",
		62:  "BurningFurnace",
		63:  "SignGround",
		64:  "WoodenDoor",
		65:  "Ladders",
		66:  "Rails",
		67:  "CobblestoneStairs",
		68:  "SignWall",
		69:  "Lever",
		70:  "StonePressurePlate",
		71:  "IronDoor",
		72:  "WoodenPressurePlate",
		73:  "RedstoneOre",
		74:  "RedstoneOreGlowing",
		75:  "RedstoneTorch",
		76:  "RedstoneTorchActive",
		77:  "StoneButton",
		78:  "SnowLayer",
		79:  "Ice",
		80:  "SnowBlock",
		81:  "Cactus",
		82:  "ClayBlock",
		83:  "SugarCane",
		84:  "Jukebox",
		85:  "Fence",
		86:  "Pumpkin",
		87:  "Netherrack",
		88:  "Soulsand",
		89:  "Glowstone",
		90:  "NetherPortal",
		91:  "JackOLantern",
		92:  "CakeBlock",
		93:  "RedstoneRepeater",
		94:  "RedstoneRepeaterActive",
		95:  "StainedGlass",
		96:  "Trapdoor",
		97:  "MonsterEgg",
		98:  "StoneBricks",
		99:  "BrownMushroomBlock",
		100: "RedMushroomBlock",
		101: "IronBars",
		102: "GlassPane",
		103: "Melon",
		104: "PumpkinStem",
		105: "MelonStem",
		106: "Vines",
		107: "FenceGate",
		108: "BrickStairs",
		109: "StoneBrickStairs",
		110: "Mycelium",
		111: "Lilypad",
		112: "NetherBrick",
		113: "NetherBrickFence",
		114: "NetherBrickStairs",
		115: "NetherWart",
		116: "EnchantmentTable",
		117: "BrewingStand",
		118: "Cauldron",
		119: "EndPortal",
		120: "EndPortalFrame",
		121: "EndStone",
		122: "DragonEgg",
		123: "RedstoneLamp",
		124: "RedstoneLampActive",
		125: "WoodenDoubleSlab",
		126: "WoodenSlab",
		127: "CocoaPod",
		128: "SandstoneStairs",
		129: "EmeraldOre",
		130: "EnderChest",
		131: "TripwireHook",
		132: "Tripwire",
		133: "EmeraldBlock",
		134: "SpruceStairs",
		135: "BirchStairs",
		136: "JungleStairs",
		137: "CommandBlock",
		138: "Beacon",
		139: "CobblestoneWall",
		140: "FlowerPot",
		141: "Carrots",
		142: "Potatoes",
		143: "WoodenButton",
		144: "Head",
		145: "Anvil",
		146: "TrappedChest",
		147: "GoldPressurePlate",
		148: "IronPressurePlate",
		149: "RedstoneComparator",
		150: "RedstoneComparatorActive",
		151: "DaylightSensor",
		152: "RedstoneBlock",
		153: "NetherQuartzOre",
		154: "Hopper",
		155: "QuartzBlock",
		156: "QuartzStairs",
		157: "ActivatorRail",
		158: "Dropper",
		159: "StainedClay",
		160: "StainedGlassPane",
		161: "Leaves2",
		162: "Log2",
		163: "AcaciaStairs",
		164: "DarkStairs",
		165: "SlimeBlock",
		166: "Barrier",
		167: "IronTrapDoor",
		168: "Prismarine",
		169: "SeaLantern",
		170: "HayBlock",
		171: "Carpet",
		172: "HardenedClay",
		173: "CoalBlock",
		174: "PackedIce",
		175: "DoublePlant",
		176: "BannerGround",
		177: "BannerWall",
		178: "InvertedDaylightSensor",
		179: "RedSandStone",
		180: "RedSandStoneStairs",
		181: "DoubleStoneSlab2",
		182: "StoneSlab2",
		183: "SpruceFenceGate",
		184: "BirchFenceGate",
		185: "JungleFenceGate",
		186: "DarkFenceGate",
		187: "AcaciaFenceGate",
		188: "SpruceFence",
		189: "BirchFence",
		190: "JungleFence",
		191: "DarkFence",
		192: "AcaciaFence",
		193: "SpruceDoor",
		194: "BirchDoor",
		195: "JungleDoor",
		196: "AcaciaDoor",
		197: "DarkDoor",
		198: "EndRod",
		199: "ChorusPlant",
		200: "ChorusFlower",
		201: "PurPurBlock",
		202: "PurPurPillar",
		203: "PurPurStairs",
		204: "PurPurDoubleSlab",
		205: "PurPurSlab",
		206: "EndBricks",
		207: "BeetRoot",
		208: "GrassPath",
		209: "EndGateWay",
		210: "RepeatingCommandBlock",
		211: "ChainCommandBlock",
		212: "FrostedIce",
		213: "MagmaBlock",
		214: "NetherWartBlock",
		215: "RedNetherBrick",
		216: "BoneBlock",
		217: "StructureVoid",
		218: "Observer",
		219: "WhiteShulkerBox",
		220: "OrangeShulkerBox",
		221: "MagentaShulkerBox",
		222: "LightBlueShulkerBox",
		223: "YellowShulkerBox",
		224: "LimeShulkerBox",
		225: "PinkShulkerBox",
		226: "GrayShulkerBox",
		227: "LightGrayShulkerBox",
		228: "CyanShulkerBox",
		229: "PurpleShulkerBox",
		230: "BlueShulkerBox",
		231: "BrownShulkerBox",
		232: "GreenShulkerBox",
		233: "RedShulkerBox",
		234: "BlackShulkerBox",
		235: "WhiteGlazedTerracotta",
		236: "OrangeGlazedTerracotta",
		237: "MagentaGlazedTerracotta",
		238: "LightBlueGlazedTerracotta",
		239: "YellowGlazedTerracotta",
		240: "LimeGlazedTerracotta",
		241: "PinkGlazedTerracotta",
		242: "GrayGlazedTerracotta",
		243: "LightGrayGlazedTerracotta",
		244: "CyanGlazedTerracotta",
		245: "PurpleGlazedTerracotta",
		246: "BlueGlazedTerracotta",
		247: "BrownGlazedTerracotta",
		248: "GreenGlazedTerracotta",
		249: "RedGlazedTerracotta",
		250: "BlackGlazedTerracotta",
		251: "Concrete",
		252: "ConcretePowder",
		253: "Undefined",
		254: "Undefined",
		255: "StructureBlock",
		256: "Undefined",
	}
	return BlockID
}
