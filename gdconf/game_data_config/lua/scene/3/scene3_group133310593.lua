-- 基础信息
local base_info = {
	group_id = 133310593
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 593001, monster_id = 28020108, pos = { x = -2499.864, y = 264.186, z = 4085.405 }, rot = { x = 0.000, y = 71.789, z = 0.000 }, level = 32, drop_tag = "走兽", area_id = 26 },
	{ config_id = 593002, monster_id = 28020108, pos = { x = -2498.429, y = 258.459, z = 4116.825 }, rot = { x = 0.000, y = 91.369, z = 0.000 }, level = 32, drop_tag = "走兽", area_id = 26 },
	{ config_id = 593003, monster_id = 28020108, pos = { x = -2518.969, y = 267.135, z = 4084.022 }, rot = { x = 0.000, y = 268.450, z = 0.000 }, level = 32, drop_tag = "走兽", area_id = 26 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { 593001, 593002, 593003 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================