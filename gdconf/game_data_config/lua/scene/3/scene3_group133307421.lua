-- 基础信息
local base_info = {
	group_id = 133307421
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 421001, gadget_id = 70500000, pos = { x = -1229.982, y = 70.070, z = 4985.546 }, rot = { x = 0.000, y = 9.750, z = 0.000 }, level = 32, point_type = 1001, area_id = 32 },
	{ config_id = 421002, gadget_id = 70500000, pos = { x = -1225.610, y = 70.687, z = 4984.353 }, rot = { x = 0.000, y = 359.246, z = 0.000 }, level = 32, point_type = 1001, area_id = 32 },
	{ config_id = 421003, gadget_id = 70500000, pos = { x = -1228.259, y = 69.557, z = 4983.522 }, rot = { x = 0.000, y = 339.783, z = 0.000 }, level = 32, point_type = 1001, area_id = 32 },
	{ config_id = 421004, gadget_id = 70500000, pos = { x = -1226.797, y = 70.247, z = 4984.146 }, rot = { x = 0.000, y = 313.262, z = 0.000 }, level = 32, point_type = 1001, area_id = 32 }
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
		monsters = { },
		gadgets = { 421001, 421002, 421003, 421004 },
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