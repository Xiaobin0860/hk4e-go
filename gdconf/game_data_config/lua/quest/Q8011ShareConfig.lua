-- 任务配置数据开始-----------------------------

main_id = 8011sub_ids = {	801101,	801102,	801103,	801104,	801105,	801106,	801107,	801108,	801109,	801110,	801111,}-- 任务配置数据结束----------------------------------- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>-- 父任务执行项数据开始-----------------------------finish_action = {	CLIENT = { },	SERVER = { },}fail_action = {	CLIENT = { },	SERVER = { },}cancel_action = {	CLIENT = { },	SERVER = { },}-- 父任务执行项数据结束------------------------------- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>-- Actor模块数据开始---------------------------------- 空-- Actor模块数据结束---------------------------------- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>-- 文本模块数据开始----------------------------------- 空-- 文本模块数据结束----------------------------------- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>-- 路点模块数据开始----------------------------------- 空-- 路点模块数据结束----------------------------------- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>-- 断线重连生成内容 开始------------------------------ 和questdata配的存档点对应rewind_data = {	["801101"] = { },	["801102"] = { },	["801103"] = { },	["801104"] = { },	["801110"] = { },	["801111"] = { },}-- 断线重连生成内容 结束------------------------------ >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>-- 校验数据 开始------------------------------------ 和任务lua中生成NPC/Monster/Gadget/Item等对应quest_data = {	["801101"] = { },	["801102"] = { },	["801103"] = 	{		npcs = 		{			{				id = 13072,				alias = "Npc13072",				script = "Actor/Npc/TempNPC",				pos = "Q801014_N13072",				scene_id = 3,				room_id = 0,				data_index = 1,			},		},	},	["801104"] = 	{		npcs = 		{			{				id = 13072,				alias = "Npc13072",				script = "Actor/Npc/TempNPC",				pos = "Q801002_N13072",				scene_id = 3,				room_id = 0,				data_index = 1,			},		},	},	["801105"] = 	{		npcs = 		{			{				id = 13072,				alias = "Npc13072",				script = "Actor/Npc/TempNPC",				pos = "Q801105_N13072",				scene_id = 3,				room_id = 0,				data_index = 1,			},			{				id = 13078,				alias = "Npc13078",				script = "Actor/Npc/TempNPC",				pos = "Q801106_N13078",				scene_id = 3,				room_id = 0,				data_index = 2,			},		},	},	["801106"] = 	{		npcs = 		{			{				id = 13078,				alias = "Npc13078",				script = "Actor/Npc/TempNPC",				pos = "Q801106_N13078",				scene_id = 3,				room_id = 0,				data_index = 1,			},		},	},	["801107"] = 	{		npcs = 		{			{				id = 13078,				alias = "Npc13078",				script = "Actor/Npc/TempNPC",				pos = "Q801106_N13078",				scene_id = 3,				room_id = 0,				data_index = 1,			},		},	},	["801108"] = 	{		npcs = 		{			{				id = 13078,				alias = "Npc13078",				script = "Actor/Npc/TempNPC",				pos = "Q801106_N13078",				scene_id = 3,				room_id = 0,				data_index = 1,			},		},	},	["801109"] = 	{		npcs = 		{			{				id = 13078,				alias = "Npc13078",				script = "Actor/Npc/TempNPC",				pos = "Q801106_N13078",				scene_id = 3,				room_id = 0,				data_index = 1,			},			{				id = 13079,				alias = "Npc13079",				script = "Actor/Npc/TempNPC",				pos = "Q801109_N13079",				scene_id = 3,				room_id = 0,				data_index = 2,			},		},	},	["801110"] = 	{		npcs = 		{			{				id = 13072,				alias = "Npc13072",				script = "Actor/Npc/TempNPC",				pos = "Q801002_N13072",				scene_id = 3,				room_id = 0,				data_index = 1,			},		},	},	["801111"] = 	{		npcs = 		{			{				id = 1022,				alias = "Npc1022",				script = "Actor/Npc/TempNPC",				pos = "Q800904_N1022",				scene_id = 3,				room_id = 0,				data_index = 1,			},		},	},}-- 校验数据 结束------------------------------------ >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>