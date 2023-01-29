-- 基础信息
local base_info = {
	group_id = 133313161
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
	{ config_id = 161001, gadget_id = 70220103, pos = { x = -405.635, y = -207.497, z = 5435.795 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 161002, gadget_id = 70220103, pos = { x = -341.412, y = -201.319, z = 5421.832 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 161003, gadget_id = 70220103, pos = { x = -345.100, y = -183.031, z = 5452.065 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 161004, gadget_id = 70220103, pos = { x = -323.232, y = -163.297, z = 5427.865 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 161005, gadget_id = 70220103, pos = { x = -375.935, y = -205.269, z = 5429.375 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 161006, gadget_id = 70220103, pos = { x = -390.112, y = -227.827, z = 5440.143 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 161007, gadget_id = 70220103, pos = { x = -324.743, y = -197.052, z = 5461.004 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 161008, gadget_id = 70220103, pos = { x = -311.061, y = -179.093, z = 5480.985 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 32 }
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
		gadgets = { 161001, 161002, 161003, 161004, 161005, 161006, 161007, 161008 },
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