package game

import (
	"math"

	"hk4e/common/config"
	"hk4e/common/constant"
	"hk4e/common/utils"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/alg"
	"hk4e/pkg/logger"
	"hk4e/pkg/reflection"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

var cmdProtoMap *cmd.CmdProtoMap = nil

func DoForward[IET model.InvokeEntryType](player *model.Player, req pb.Message, copyFieldList []string, forwardField string, invokeHandler *model.InvokeHandler[IET]) {
	if cmdProtoMap == nil {
		cmdProtoMap = cmd.NewCmdProtoMap()
	}
	cmdId := cmdProtoMap.GetCmdIdByProtoObj(req)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	if invokeHandler.AllLen() == 0 && invokeHandler.AllExceptCurLen() == 0 && invokeHandler.HostLen() == 0 {
		ntf := cmdProtoMap.GetProtoObjByCmdId(cmdId)
		for _, fieldName := range copyFieldList {
			reflection.CopyStructField(ntf, req, fieldName)
		}
		for _, v := range world.GetAllPlayer() {
			GAME_MANAGER.SendMsg(cmdId, v.PlayerID, player.ClientSeq, ntf)
		}
	}
	if invokeHandler.AllLen() > 0 {
		ntf := cmdProtoMap.GetProtoObjByCmdId(cmdId)
		for _, fieldName := range copyFieldList {
			reflection.CopyStructField(ntf, req, fieldName)
		}
		reflection.SetStructFieldValue(ntf, forwardField, invokeHandler.EntryListForwardAll)
		GAME_MANAGER.SendToWorldA(world, cmdId, player.ClientSeq, ntf)
	}
	if invokeHandler.AllExceptCurLen() > 0 {
		ntf := cmdProtoMap.GetProtoObjByCmdId(cmdId)
		for _, fieldName := range copyFieldList {
			reflection.CopyStructField(ntf, req, fieldName)
		}
		reflection.SetStructFieldValue(ntf, forwardField, invokeHandler.EntryListForwardAllExceptCur)
		GAME_MANAGER.SendToWorldAEC(world, cmdId, player.ClientSeq, ntf, player.PlayerID)
	}
	if invokeHandler.HostLen() > 0 {
		ntf := cmdProtoMap.GetProtoObjByCmdId(cmdId)
		for _, fieldName := range copyFieldList {
			reflection.CopyStructField(ntf, req, fieldName)
		}
		reflection.SetStructFieldValue(ntf, forwardField, invokeHandler.EntryListForwardHost)
		GAME_MANAGER.SendToWorldH(world, cmdId, player.ClientSeq, ntf)
	}
}

func (g *GameManager) UnionCmdNotify(player *model.Player, payloadMsg pb.Message) {
	// logger.Debug("user send union cmd, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.UnionCmdNotify)
	_ = req
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	DoForward[proto.CombatInvokeEntry](player, &proto.CombatInvocationsNotify{}, []string{}, "InvokeList", player.CombatInvokeHandler)
	DoForward[proto.AbilityInvokeEntry](player, &proto.AbilityInvocationsNotify{}, []string{}, "Invokes", player.AbilityInvokeHandler)
	player.CombatInvokeHandler.Clear()
	player.AbilityInvokeHandler.Clear()
}

func (g *GameManager) MassiveEntityElementOpBatchNotify(player *model.Player, payloadMsg pb.Message) {
	// logger.Debug("user meeo sync, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.MassiveEntityElementOpBatchNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	ntf := req
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v", player.SceneId)
		return
	}
	ntf.OpIdx = scene.GetMeeoIndex()
	scene.SetMeeoIndex(scene.GetMeeoIndex() + 1)
	g.SendToWorldA(world, cmd.MassiveEntityElementOpBatchNotify, player.ClientSeq, ntf)
}

func (g *GameManager) CombatInvocationsNotify(player *model.Player, payloadMsg pb.Message) {
	// logger.Debug("user combat invocations, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.CombatInvocationsNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v", player.SceneId)
		return
	}
	for _, entry := range req.InvokeList {
		switch entry.ArgumentType {
		case proto.CombatTypeArgument_COMBAT_EVT_BEING_HIT:
			hitInfo := new(proto.EvtBeingHitInfo)
			if config.GetConfig().Hk4e.ClientProtoProxyEnable {
				clientProtoObj := g.GetClientProtoObjByName("EvtBeingHitInfo")
				if clientProtoObj == nil {
					logger.Error("get client proto obj is nil")
					continue
				}
				ok := utils.UnmarshalProtoObj(hitInfo, clientProtoObj, entry.CombatData)
				if !ok {
					continue
				}
			} else {
				err := pb.Unmarshal(entry.CombatData, hitInfo)
				if err != nil {
					logger.Error("parse EvtBeingHitInfo error: %v", err)
					continue
				}
			}
			attackResult := hitInfo.AttackResult
			if attackResult == nil {
				logger.Error("attackResult is nil")
				continue
			}
			logger.Debug("run attack handler, attackResult: %v", attackResult)
			target := scene.GetEntity(attackResult.DefenseId)
			if target == nil {
				logger.Error("could not found target, defense id: %v", attackResult.DefenseId)
				continue
			}
			attackResult.Damage *= 10
			damage := attackResult.Damage
			attackerId := attackResult.AttackerId
			_ = attackerId
			currHp := float32(0)
			fightProp := target.GetFightProp()
			if fightProp != nil {
				currHp = fightProp[constant.FIGHT_PROP_CUR_HP]
				currHp -= damage
				if currHp < 0 {
					currHp = 0
				}
				fightProp[constant.FIGHT_PROP_CUR_HP] = currHp
			}
			entityFightPropUpdateNotify := &proto.EntityFightPropUpdateNotify{
				FightPropMap: fightProp,
				EntityId:     target.GetId(),
			}
			g.SendToWorldA(world, cmd.EntityFightPropUpdateNotify, player.ClientSeq, entityFightPropUpdateNotify)
			if currHp == 0 && target.GetEntityType() != constant.ENTITY_TYPE_AVATAR {
				scene.SetEntityLifeState(target, constant.LIFE_STATE_DEAD, proto.PlayerDieType_PLAYER_DIE_GM)
			}
			combatData, err := pb.Marshal(hitInfo)
			if err != nil {
				logger.Error("create combat invocations entity hit info error: %v", err)
			}
			entry.CombatData = combatData
			player.CombatInvokeHandler.AddEntry(entry.ForwardType, entry)
		case proto.CombatTypeArgument_ENTITY_MOVE:
			entityMoveInfo := new(proto.EntityMoveInfo)
			if config.GetConfig().Hk4e.ClientProtoProxyEnable {
				clientProtoObj := g.GetClientProtoObjByName("EntityMoveInfo")
				if clientProtoObj == nil {
					logger.Error("get client proto obj is nil")
					continue
				}
				ok := utils.UnmarshalProtoObj(entityMoveInfo, clientProtoObj, entry.CombatData)
				if !ok {
					continue
				}
			} else {
				err := pb.Unmarshal(entry.CombatData, entityMoveInfo)
				if err != nil {
					logger.Error("parse EntityMoveInfo error: %v", err)
					continue
				}
			}
			motionInfo := entityMoveInfo.MotionInfo
			if motionInfo.Pos == nil || motionInfo.Rot == nil {
				continue
			}
			sceneEntity := scene.GetEntity(entityMoveInfo.EntityId)
			if sceneEntity == nil {
				continue
			}
			if sceneEntity.GetEntityType() == constant.ENTITY_TYPE_AVATAR {
				// 玩家实体在移动
				g.AoiPlayerMove(player, player.Pos, &model.Vector{
					X: float64(motionInfo.Pos.X),
					Y: float64(motionInfo.Pos.Y),
					Z: float64(motionInfo.Pos.Z),
				}, sceneEntity.GetId())
				// 更新玩家的位置信息
				player.Pos.X = float64(motionInfo.Pos.X)
				player.Pos.Y = float64(motionInfo.Pos.Y)
				player.Pos.Z = float64(motionInfo.Pos.Z)
				player.Rot.X = float64(motionInfo.Rot.X)
				player.Rot.Y = float64(motionInfo.Rot.Y)
				player.Rot.Z = float64(motionInfo.Rot.Z)

				// 处理耐力消耗
				g.ImmediateStamina(player, motionInfo.State)
			} else {
				// 非玩家实体在移动
				// 更新场景实体的位置信息
				pos := sceneEntity.GetPos()
				pos.X = float64(motionInfo.Pos.X)
				pos.Y = float64(motionInfo.Pos.Y)
				pos.Z = float64(motionInfo.Pos.Z)
				rot := sceneEntity.GetRot()
				rot.X = float64(motionInfo.Rot.X)
				rot.Y = float64(motionInfo.Rot.Y)
				rot.Z = float64(motionInfo.Rot.Z)
				if sceneEntity.GetEntityType() == constant.ENTITY_TYPE_GADGET {
					// 载具耐力消耗
					gadgetEntity := sceneEntity.GetGadgetEntity()
					if gadgetEntity.GetGadgetVehicleEntity() != nil {
						// 处理耐力消耗
						g.ImmediateStamina(player, motionInfo.State)
						// 处理载具销毁请求
						g.VehicleDestroyMotion(player, sceneEntity, motionInfo.State)
					}
				}
			}
			sceneEntity.SetMoveState(uint16(motionInfo.State))
			sceneEntity.SetLastMoveSceneTimeMs(entityMoveInfo.SceneTime)
			sceneEntity.SetLastMoveReliableSeq(entityMoveInfo.ReliableSeq)

			if motionInfo.State == proto.MotionState_MOTION_NOTIFY {
				continue
			}

			player.CombatInvokeHandler.AddEntry(entry.ForwardType, entry)
		case proto.CombatTypeArgument_COMBAT_ANIMATOR_PARAMETER_CHANGED:
			evtAnimatorParameterInfo := new(proto.EvtAnimatorParameterInfo)
			err := pb.Unmarshal(entry.CombatData, evtAnimatorParameterInfo)
			if err != nil {
				logger.Error("parse EvtAnimatorParameterInfo error: %v", err)
				continue
			}
			logger.Debug("EvtAnimatorParameterInfo: %v, ForwardType: %v", evtAnimatorParameterInfo, entry.ForwardType)

			// 这是否?
			if evtAnimatorParameterInfo.IsServerCache {
				evtAnimatorParameterInfo.IsServerCache = false
				// TODO 暂时只做3.2协议的兼容了 在GS这边处理不同版本的协议太烦人了 有机会全部改到GATE那边处理 GS所有接收和发送的都应该是3.2版本的协议
				newCombatData, err := pb.Marshal(evtAnimatorParameterInfo)
				if err != nil {
					logger.Error("build EvtAnimatorParameterInfo error: %v", err)
					continue
				}
				entry.CombatData = newCombatData
			}

			player.CombatInvokeHandler.AddEntry(entry.ForwardType, entry)
		case proto.CombatTypeArgument_COMBAT_ANIMATOR_STATE_CHANGED:
			evtAnimatorStateChangedInfo := new(proto.EvtAnimatorStateChangedInfo)
			err := pb.Unmarshal(entry.CombatData, evtAnimatorStateChangedInfo)
			if err != nil {
				logger.Error("parse EvtAnimatorStateChangedInfo error: %v", err)
				continue
			}
			logger.Debug("EvtAnimatorStateChangedInfo: %v, ForwardType: %v", evtAnimatorStateChangedInfo, entry.ForwardType)
			player.CombatInvokeHandler.AddEntry(entry.ForwardType, entry)
		default:
			player.CombatInvokeHandler.AddEntry(entry.ForwardType, entry)
		}
	}
}

func (g *GameManager) AoiPlayerMove(player *model.Player, oldPos *model.Vector, newPos *model.Vector, entityId uint32) {
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		logger.Error("get player world is nil, uid: %v", player.PlayerID)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	sceneBlockAoiMap := WORLD_MANAGER.GetSceneBlockAoiMap()
	aoiManager, exist := sceneBlockAoiMap[player.SceneId]
	if !exist {
		logger.Error("get scene block aoi is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerID)
		return
	}
	oldGid := aoiManager.GetGidByPos(float32(oldPos.X), 0.0, float32(oldPos.Z))
	newGid := aoiManager.GetGidByPos(float32(newPos.X), 0.0, float32(newPos.Z))
	if oldGid != newGid {
		// 跨越了block格子
		logger.Debug("player cross grid, oldGid: %v, newGid: %v, uid: %v", oldGid, newGid, player.PlayerID)
	}
	// 旧位置视野范围内的group
	oldVisionGroupMap := make(map[uint32]*gdconf.Group)
	oldGroupList := aoiManager.GetObjectListByPos(float32(oldPos.X), 0.0, float32(oldPos.Z))
	for groupId, groupAny := range oldGroupList {
		group := groupAny.(*gdconf.Group)
		distance2D := math.Sqrt(math.Pow(oldPos.X-float64(group.Pos.X), 2.0) + math.Pow(oldPos.Z-float64(group.Pos.Z), 2.0))
		if distance2D > ENTITY_LOD {
			continue
		}
		oldVisionGroupMap[uint32(groupId)] = group
	}
	// 新位置视野范围内的group
	newVisionGroupMap := make(map[uint32]*gdconf.Group)
	newGroupList := aoiManager.GetObjectListByPos(float32(newPos.X), 0.0, float32(newPos.Z))
	for groupId, groupAny := range newGroupList {
		group := groupAny.(*gdconf.Group)
		distance2D := math.Sqrt(math.Pow(newPos.X-float64(group.Pos.X), 2.0) + math.Pow(newPos.Z-float64(group.Pos.Z), 2.0))
		if distance2D > ENTITY_LOD {
			continue
		}
		newVisionGroupMap[uint32(groupId)] = group
	}
	// 消失的场景实体
	delEntityIdList := make([]uint32, 0)
	for groupId, group := range oldVisionGroupMap {
		_, exist := newVisionGroupMap[groupId]
		if exist {
			continue
		}
		// 旧有新没有的group即为消失的
		for _, monster := range group.MonsterList {
			entity := scene.GetEntityByObjectId(monster.ObjectId)
			if entity == nil {
				continue
			}
			scene.DestroyEntity(entity.GetId())
			delEntityIdList = append(delEntityIdList, entity.GetId())
		}
		for _, npc := range group.NpcList {
			entity := scene.GetEntityByObjectId(npc.ObjectId)
			if entity == nil {
				continue
			}
			scene.DestroyEntity(entity.GetId())
			delEntityIdList = append(delEntityIdList, entity.GetId())
		}
		for _, gadget := range group.GadgetList {
			entity := scene.GetEntityByObjectId(gadget.ObjectId)
			if entity == nil {
				continue
			}
			scene.DestroyEntity(entity.GetId())
			delEntityIdList = append(delEntityIdList, entity.GetId())
		}
	}
	// 出现的场景实体
	addEntityIdList := make([]uint32, 0)
	for groupId, group := range newVisionGroupMap {
		_, exist := oldVisionGroupMap[groupId]
		if exist {
			continue
		}
		// 新有旧没有的group即为出现的
		for _, monster := range group.MonsterList {
			entityId := g.CreateConfigEntity(scene, monster.ObjectId, monster)
			addEntityIdList = append(addEntityIdList, entityId)
		}
		for _, npc := range group.NpcList {
			entityId := g.CreateConfigEntity(scene, npc.ObjectId, npc)
			addEntityIdList = append(addEntityIdList, entityId)
		}
		for _, gadget := range group.GadgetList {
			entityId := g.CreateConfigEntity(scene, gadget.ObjectId, gadget)
			addEntityIdList = append(addEntityIdList, entityId)
		}
	}
	// 同步客户端消失和出现的场景实体
	g.RemoveSceneEntityNotifyToPlayer(player, proto.VisionType_VISION_MISS, delEntityIdList)
	g.AddSceneEntityNotify(player, proto.VisionType_VISION_MEET, addEntityIdList, false, false)
	// 场景区域触发器
	dbQuest := player.GetDbQuest()
	for _, group := range newVisionGroupMap {
		for _, region := range group.RegionList {
			shape := alg.NewShape()
			switch uint8(region.Shape) {
			case constant.REGION_SHAPE_SPHERE:
				shape.NewSphere(&alg.Vector3{X: region.Pos.X, Y: region.Pos.Y, Z: region.Pos.Z}, region.Radius)
			case constant.REGION_SHAPE_CUBIC:
				shape.NewCubic(&alg.Vector3{X: region.Pos.X, Y: region.Pos.Y, Z: region.Pos.Z},
					&alg.Vector3{X: region.Size.X, Y: region.Size.Y, Z: region.Size.Z})
			case constant.REGION_SHAPE_CYLINDER:
				shape.NewCylinder(&alg.Vector3{X: region.Pos.X, Y: region.Pos.Y, Z: region.Pos.Z},
					region.Radius, region.Height)
			case constant.REGION_SHAPE_POLYGON:
				vector2PointArray := make([]*alg.Vector2, 0)
				for _, vector := range region.PointArray {
					// z就是y
					vector2PointArray = append(vector2PointArray, &alg.Vector2{X: vector.X, Z: vector.Y})
				}
				shape.NewPolygon(&alg.Vector3{X: region.Pos.X, Y: region.Pos.Y, Z: region.Pos.Z},
					vector2PointArray, region.Height)
			}
			oldPosInRegion := shape.Contain(&alg.Vector3{
				X: float32(oldPos.X),
				Y: float32(oldPos.Y),
				Z: float32(oldPos.Z),
			})
			newPosInRegion := shape.Contain(&alg.Vector3{
				X: float32(newPos.X),
				Y: float32(newPos.Y),
				Z: float32(newPos.Z),
			})
			if !oldPosInRegion && newPosInRegion {
				// EVENT_ENTER_REGION
				logger.Debug("player enter region: %v, uid: %v", region, player.PlayerID)
				for _, trigger := range group.TriggerList {
					if trigger.Event != constant.LUA_EVENT_ENTER_REGION {
						continue
					}
					cond := CallLuaFunc(group.LuaState, trigger.Condition, &LuaCtx{
						uid: player.PlayerID,
					}, &LuaEvt{
						param1:         region.ConfigId,
						targetEntityId: entityId,
					})
					if !cond {
						continue
					}
					// TODO 这一块写得太炸裂了需要优化
					for _, triggerDataConfig := range gdconf.GetTriggerDataMap() {
						if triggerDataConfig.TriggerName == trigger.Name {
							for _, quest := range dbQuest.GetQuestMap() {
								questDataConfig := gdconf.GetQuestDataById(int32(quest.QuestId))
								if questDataConfig == nil {
									continue
								}
								for _, questCond := range questDataConfig.FinishCondList {
									if questCond.Type != constant.QUEST_FINISH_COND_TYPE_TRIGGER_FIRE {
										continue
									}
									if questCond.Param[0] != triggerDataConfig.TriggerId {
										continue
									}
									dbQuest.ForceFinishQuest(quest.QuestId)
								}
							}
						}
					}
				}
			} else if oldPosInRegion && !newPosInRegion {
				// EVENT_LEAVE_REGION
				logger.Debug("player leave region: %v, uid: %v", region, player.PlayerID)
			}
		}
	}
	g.AcceptQuest(player, true)
}

func (g *GameManager) AbilityInvocationsNotify(player *model.Player, payloadMsg pb.Message) {
	// logger.Debug("user ability invocations, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AbilityInvocationsNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}

	for _, entry := range req.Invokes {
		// logger.Debug("AbilityInvocationsNotify: %v", entry, player.PlayerID)

		// switch entry.ArgumentType {
		// case proto.AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_MODIFIER_CHANGE:
		//	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
		//	worldAvatar := world.GetWorldAvatarByEntityId(entry.EntityId)
		//	if worldAvatar != nil {
		//		for _, ability := range worldAvatar.abilityList {
		//			if ability.InstancedAbilityId == entry.Head.InstancedAbilityId {
		//				logger.Error("A: %v", ability)
		//			}
		//		}
		//		for _, modifier := range worldAvatar.modifierList {
		//			if modifier.InstancedAbilityId == entry.Head.InstancedAbilityId {
		//				logger.Error("B: %v", modifier)
		//			}
		//		}
		//		for _, modifier := range worldAvatar.modifierList {
		//			if modifier.InstancedModifierId == entry.Head.InstancedModifierId {
		//				logger.Error("C: %v", modifier)
		//			}
		//		}
		//	}
		// case proto.AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_NONE:
		// }

		// 处理耐力消耗
		g.HandleAbilityStamina(player, entry)
		player.AbilityInvokeHandler.AddEntry(entry.ForwardType, entry)
	}
}

func (g *GameManager) ClientAbilityInitFinishNotify(player *model.Player, payloadMsg pb.Message) {
	// logger.Debug("user client ability init finish, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.ClientAbilityInitFinishNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	invokeHandler := model.NewInvokeHandler[proto.AbilityInvokeEntry]()
	for _, entry := range req.Invokes {
		// logger.Debug("ClientAbilityInitFinishNotify: %v", entry)
		invokeHandler.AddEntry(entry.ForwardType, entry)
	}
	DoForward[proto.AbilityInvokeEntry](player, &proto.ClientAbilityInitFinishNotify{}, []string{"EntityId"}, "Invokes", invokeHandler)
}

func (g *GameManager) ClientAbilityChangeNotify(player *model.Player, payloadMsg pb.Message) {
	// logger.Debug("user client ability change, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.ClientAbilityChangeNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	invokeHandler := model.NewInvokeHandler[proto.AbilityInvokeEntry]()
	for _, entry := range req.Invokes {
		// logger.Debug("ClientAbilityChangeNotify: %v", entry)

		invokeHandler.AddEntry(entry.ForwardType, entry)
	}
	DoForward[proto.AbilityInvokeEntry](player, req, []string{"EntityId", "IsInitHash"}, "Invokes", invokeHandler)

	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	for _, abilityInvokeEntry := range req.Invokes {
		switch abilityInvokeEntry.ArgumentType {
		case proto.AbilityInvokeArgument_ABILITY_META_ADD_NEW_ABILITY:
			abilityMetaAddAbility := new(proto.AbilityMetaAddAbility)
			if config.GetConfig().Hk4e.ClientProtoProxyEnable {
				clientProtoObj := g.GetClientProtoObjByName("AbilityMetaAddAbility")
				if clientProtoObj == nil {
					logger.Error("get client proto obj is nil")
					continue
				}
				ok := utils.UnmarshalProtoObj(abilityMetaAddAbility, clientProtoObj, abilityInvokeEntry.AbilityData)
				if !ok {
					logger.Error("AbilityMetaAddAbility proto error")
					continue
				}
			} else {
				err := pb.Unmarshal(abilityInvokeEntry.AbilityData, abilityMetaAddAbility)
				if err != nil {
					logger.Error("parse AbilityMetaAddAbility error: %v", err)
					continue
				}
			}
			worldAvatar := world.GetWorldAvatarByEntityId(abilityInvokeEntry.EntityId)
			if worldAvatar == nil {
				continue
			}
			if abilityMetaAddAbility.Ability == nil {
				continue
			}
			abilityList := worldAvatar.GetAbilityList()
			abilityList = append(abilityList, abilityMetaAddAbility.Ability)
			worldAvatar.SetAbilityList(abilityList)
		case proto.AbilityInvokeArgument_ABILITY_META_MODIFIER_CHANGE:
			abilityMetaModifierChange := new(proto.AbilityMetaModifierChange)
			if config.GetConfig().Hk4e.ClientProtoProxyEnable {
				clientProtoObj := g.GetClientProtoObjByName("AbilityMetaModifierChange")
				if clientProtoObj == nil {
					logger.Error("get client proto obj is nil")
					continue
				}
				ok := utils.UnmarshalProtoObj(abilityMetaModifierChange, clientProtoObj, abilityInvokeEntry.AbilityData)
				if !ok {
					logger.Error("AbilityMetaModifierChange proto error")
					continue
				}
			} else {
				err := pb.Unmarshal(abilityInvokeEntry.AbilityData, abilityMetaModifierChange)
				if err != nil {
					logger.Error("parse AbilityMetaModifierChange error: %v", err)
					continue
				}
			}
			abilityAppliedModifier := &proto.AbilityAppliedModifier{
				ModifierLocalId:           abilityMetaModifierChange.ModifierLocalId,
				ParentAbilityEntityId:     0,
				ParentAbilityName:         abilityMetaModifierChange.ParentAbilityName,
				ParentAbilityOverride:     abilityMetaModifierChange.ParentAbilityOverride,
				InstancedAbilityId:        abilityInvokeEntry.Head.InstancedAbilityId,
				InstancedModifierId:       abilityInvokeEntry.Head.InstancedModifierId,
				ExistDuration:             0,
				AttachedInstancedModifier: abilityMetaModifierChange.AttachedInstancedModifier,
				ApplyEntityId:             abilityMetaModifierChange.ApplyEntityId,
				IsAttachedParentAbility:   abilityMetaModifierChange.IsAttachedParentAbility,
				ModifierDurability:        nil,
				SbuffUid:                  0,
				IsServerbuffModifier:      abilityInvokeEntry.Head.IsServerbuffModifier,
			}
			worldAvatar := world.GetWorldAvatarByEntityId(abilityInvokeEntry.EntityId)
			if worldAvatar == nil {
				continue
			}
			modifierList := worldAvatar.GetModifierList()
			modifierList = append(modifierList, abilityAppliedModifier)
			worldAvatar.SetModifierList(modifierList)
		default:
		}
	}
}

func (g *GameManager) EvtDoSkillSuccNotify(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user event do skill success, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.EvtDoSkillSuccNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	logger.Debug("EvtDoSkillSuccNotify: %v", req)

	// 处理技能开始的耐力消耗
	g.SkillStartStamina(player, req.CasterId, req.SkillId)
}

func (g *GameManager) EvtAvatarEnterFocusNotify(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user avatar enter focus, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.EvtAvatarEnterFocusNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	logger.Debug("EvtAvatarEnterFocusNotify: %v", req)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	g.SendToWorldA(world, cmd.EvtAvatarEnterFocusNotify, player.ClientSeq, req)
}

func (g *GameManager) EvtAvatarUpdateFocusNotify(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user avatar update focus, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.EvtAvatarUpdateFocusNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	logger.Debug("EvtAvatarUpdateFocusNotify: %v", req)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	g.SendToWorldA(world, cmd.EvtAvatarUpdateFocusNotify, player.ClientSeq, req)
}

func (g *GameManager) EvtAvatarExitFocusNotify(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user avatar exit focus, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.EvtAvatarExitFocusNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	logger.Debug("EvtAvatarExitFocusNotify: %v", req)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	g.SendToWorldA(world, cmd.EvtAvatarExitFocusNotify, player.ClientSeq, req)
}

func (g *GameManager) EvtEntityRenderersChangedNotify(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user entity render change, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.EvtEntityRenderersChangedNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	logger.Debug("EvtEntityRenderersChangedNotify: %v", req)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	g.SendToWorldA(world, cmd.EvtEntityRenderersChangedNotify, player.ClientSeq, req)
}

func (g *GameManager) EvtCreateGadgetNotify(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user create gadget, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.EvtCreateGadgetNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	logger.Debug("EvtCreateGadgetNotify: %v", req)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v", player.SceneId)
		return
	}
	scene.CreateEntityGadgetClient(&model.Vector{
		X: float64(req.InitPos.X),
		Y: float64(req.InitPos.Y),
		Z: float64(req.InitPos.Z),
	}, &model.Vector{
		X: float64(req.InitEulerAngles.X),
		Y: float64(req.InitEulerAngles.Y),
		Z: float64(req.InitEulerAngles.Z),
	}, req.EntityId, req.ConfigId, req.CampId, req.CampType, req.OwnerEntityId, req.TargetEntityId, req.PropOwnerEntityId)
	g.AddSceneEntityNotify(player, proto.VisionType_VISION_BORN, []uint32{req.EntityId}, true, true)
}

func (g *GameManager) EvtDestroyGadgetNotify(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user destroy gadget, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.EvtDestroyGadgetNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	logger.Debug("EvtDestroyGadgetNotify: %v", req)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v", player.SceneId)
		return
	}
	scene.DestroyEntity(req.EntityId)
	g.RemoveSceneEntityNotifyBroadcast(scene, proto.VisionType_VISION_MISS, []uint32{req.EntityId})
}
