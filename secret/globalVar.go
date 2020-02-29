package secret

var secretInfo = [...]SecretInfo{
	{SecretName: "黑夜", SecretLevelName: [10]string{"不眠者", "午夜诗人", "梦魇", "安魂师", "灵巫", "守夜人", "恐惧主教", "隐秘之仆", "未知", "黑夜"}},
	{SecretName: "死神", SecretLevelName: [10]string{"收尸人", "掘墓人", "通灵者", "死灵导师", "看门人", "不死者", "摆渡人", "死亡执政官", "未知", "死神"}},
	{SecretName: "战神", SecretLevelName: [10]string{"战士", "格斗家", "武器大师", "黎明骑士", "守护者", "猎魔人", "银骑士", "未知", "未知", "战神"}},
	{SecretName: "暗", SecretLevelName: [10]string{"秘祈人", "倾听者", "隐修士", "蔷薇主教", "牧羊人", "黑骑士", "三首圣堂", "未知", "未知", "未知"}},
	{SecretName: "心灵", SecretLevelName: [10]string{"观众", "读心者", "心理医生", "催眠师", "梦境行者", "操纵师", "织梦者", "洞察者", "作家", "空想家"}},
	{SecretName: "风暴", SecretLevelName: [10]string{"水手", "暴怒之民", "航海家", "风眷者", "海洋歌者", "灾难主祭", "海王", "天灾", "雷神", "暴君"}},
	{SecretName: "智慧", SecretLevelName: [10]string{"阅读者", "推理学员", "侦探", "博学者", "秘术导师", "预言家", "未知", "未知", "未知", "未知"}},
	{SecretName: "太阳", SecretLevelName: [10]string{"歌颂者", "祈光人", "太阳神官", "公证人", "光之祭司", "无暗者", "正义导师", "逐光者", "未知", "太阳"}},
	{SecretName: "异种", SecretLevelName: [10]string{"囚犯", "疯子", "狼人", "活尸", "怨灵", "木偶", "沉默门徒", "古代邪物", "神孽", "未知"}},
	{SecretName: "深渊", SecretLevelName: [10]string{"罪犯", "冷血者", "连环杀手", "恶魔", "欲望使徒", "魔鬼", "呓语者", "未知", "未知", "深渊"}},
	{SecretName: "秘密", SecretLevelName: [10]string{"窥秘人", "格斗教授", "巫师", "卷轴教授", "星象师", "神秘学者", "预言大师", "贤者", "知识皇帝", "未知"}},
	{SecretName: "工匠", SecretLevelName: [10]string{"通识者", "考古学家", "鉴定师", "机械专家", "天文学家", "炼金术士", "奥秘学者", "知识导师", "启蒙者", "完美者"}},
	{SecretName: "愚者", SecretLevelName: [10]string{"占卜家", "小丑", "魔法师", "无面人", "秘偶大师", "诡法师", "古代学者", "奇迹师", "诡秘侍者", "愚者"}},
	{SecretName: "门", SecretLevelName: [10]string{"学徒", "戏法大师", "占星人", "记录官", "旅行家", "秘法师", "漫游者", "旅法师", "星之匙", "未知"}},
	{SecretName: "时间", SecretLevelName: [10]string{"偷盗者", "诈骗师", "解密学者", "盗火人", "窃梦家", "寄生者", "欺瞒导师", "命运木马", "时之虫", "错误"}},
	{SecretName: "大地", SecretLevelName: [10]string{"耕种者", "医师", "丰收祭司", "生物学家", "德鲁伊", "古代炼金师", "未知", "未知", "未知", "未知"}},
	{SecretName: "月亮", SecretLevelName: [10]string{"药师", "驯兽师", "吸血鬼", "魔药教授", "深红学者", "巫王", "召唤大师", "未知", "未知", "未知"}},
	{SecretName: "命运", SecretLevelName: [10]string{"怪物", "机器", "幸运儿", "灾祸教士", "赢家", "厄运法师", "混乱行者", "先知", "水银之蛇", "命运之轮"}},
	{SecretName: "审判", SecretLevelName: [10]string{"仲裁人", "治安官", "审讯者", "法官", "惩戒骑士", "律令法师", "混乱猎手", "平衡者", "秩序之手", "审判者"}},
	{SecretName: "黑皇帝", SecretLevelName: [10]string{"律师", "野蛮人", "贿赂者", "腐化男爵", "混乱导师", "堕落伯爵", "狂乱法师", "熵之公爵", "弑序亲王", "黑皇帝"}},
	{SecretName: "战争", SecretLevelName: [10]string{"猎人", "挑衅者", "纵火家", "阴谋家", "收割者", "铁血骑士", "战争主教", "天气术士", "征服者", "红祭司"}},
	{SecretName: "魔女", SecretLevelName: [10]string{"刺客", "教唆者", "女巫", "欢愉", "痛苦", "绝望", "不老", "灾难", "末日", "原初"}},
}

var secretGroup = [...][]uint64{
	{0, 1, 2},
	{3, 4, 5, 6, 7},
	{8, 9},
	{10, 11},
	{12, 13, 14},
	{15, 16},
	{17},
	{18, 19},
	{20, 21},
}

var fight = [...]string{
	"不堪一击",
	"毫不足虑",
	"不足挂齿",
	"初学乍练",
	"勉勉强强",
	"初窥门径",
	"初出茅庐",
	"略知一二",
	"普普通通",
	"平平常常",
	"平淡无奇",
	"粗懂皮毛",
	"半生不熟",
	"登堂入室",
	"略有小成",
	"已有小成",
	"鹤立鸡群",
	"驾轻就熟",
	"青出於蓝",
	"融会贯通",
	"心领神会",
	"炉火纯青",
	"了然於胸",
	"波澜不惊",
	"略有大成",
	"已有大成",
	"豁然贯通",
	"不同凡响",
	"非比寻常",
	"出类拔萃",
	"罕有敌手",
	"波澜老成",
	"冠绝一时",
	"锐不可当",
	"断蛟伏虎",
	"百战不殆",
	"技冠群雄",
	"超群绝伦",
	"神乎其技",
	"出神入化",
	"傲视群雄",
	"超尘拔俗",
	"无出其右",
	"登峰造极",
	"无与伦比",
	"靡坚不摧",
	"所向披靡",
	"一代宗师",
	"世外高人",
	"精深奥妙",
	"神功盖世",
	"望而生遁",
	"勇冠三军",
	"横扫千军",
	"万敌莫开",
	"举世无双",
	"独步天下",
	"指点江山",
	"惊世骇俗",
	"撼天动地",
	"震古铄今",
	"亘古一人",
	"席卷八荒",
	"超凡入圣",
	"威镇寰宇",
	"空前绝后",
	"天人合一",
	"深藏不露",
	"深不可测",
	"高深莫测",
	"岿然如峰",
	"势如破竹",
	"势涌彪发",
	"叱咤喑呜",
	"跌宕昭彰",
	"拿云攫石",
	"摧朽拉枯",
	"鹰撮霆击",
	"鳌掷鲸吞",
	"潮鸣电掣",
	"风云变色",
	"风行雷厉",
	"拔山举鼎",
	"山呼海啸",
	"回山倒海",
	"倒峡泻河",
	"熏天赫地",
	"拔地参天",
	"万仞之巅",
	"气贯长虹",
	"气吞山河",
	"欺霜傲雪",
	"立海垂云",
	"移山竭海",
	"凤翥龙翔",
	"摘星逐月",
	"气凌霄汉",
	"扭转乾坤",
	"洗尽铅华",
	"返璞归真",
}

// 类型：1）-钱＆经验5%；2）-经验10%；3）-钱10%；4）+钱＆经验10%；5）+钱15%；6）+经验15%；7）无事发生35%
var adventureEvents = [7]AdventureEvent{
	AdventureEvent{Type: 1, Probability: 5, Messages: []string{
		"在探险过程中遇到危险不得不使用保命物品才得以逃脱，损失惨重。",
		"被异教徒盯上并暴力劫掠，损失惨重。",
		"不幸遭遇了官方非凡者对普通非凡者的追捕，仓皇逃窜。",
	}},
	AdventureEvent{Type: 2, Probability: 10, Messages: []string{
		"由于错误的祷告，被不知名的存在盯上，受到打击。",
		"飞鱼与酒酒吧的列朗齐味道不错，你一口气喝了好几大杯。晕头转向中有混混想用所谓本地土特产腌肉把戏强买强卖。你愤怒地想要回击却被他的同伙团团围住，双拳难敌四手，被迫挨打。",
		"被蒸汽与机械教会的狂热研究者抓去实验新制作的非凡物品，腰酸背痛地艰难逃脱。",
		"没时间解释了，列奥德罗！一道闪电劈下，",
		"在探险过程中用尽了力量不得不返回，灵性枯竭。",
		"被0-008安排到大佬面前送死，就这样结束了吗…就在这时一位灰雾萦绕的圣洁天使突然从天而降用羽翼将你层层包裹。你失去了部分记忆，但最终逃脱了安排。赞美愚者。",
	}},
	AdventureEvent{Type: 3, Probability: 10, Messages: []string{
		"在非凡者聚会上买到假货，该死。",
		"交房租的时间到了，老天。",
		"偶遇极光会吃手指的邪教徒，逃跑的路上不幸丢失了些金钱。",
		"飞鱼与酒酒吧的列朗齐味道不错，你一口气喝了好几大杯。晕头转向中被混混用所谓本地土特产腌肉的强买强卖把戏骗走了些钱。腌肉味道还行，可你还是心疼钱包。",
		"路过拜亚姆土著孤儿幼儿园，看到知名大海盗烈焰正在指挥工人们给孩子们修建新的操场。你忍不住给混血儿们捐了点钱。虽然钱包瘪了一点但是心情很愉快。",
		"魔女的滋味真不错啊…一觉醒来佳人已不知所踪。",
		"被一位叫佛尔思的女士拉进咖啡馆询问问题，据她说这是要搜集小说的素材…？你讲得口干舌燥，不由连着喝了好几杯茶。",
		"和奥黛丽一起救济了战乱中贝克兰德的百姓们，今晚能够多吃一点东西的孩子们又多了几个。",
		"被命运议员附加幸运，兴致勃勃地去赌博却损失金钱。",
	}},
	AdventureEvent{Type: 4, Probability: 10, Messages: []string{
		"探索了一座遗迹，了解到一些隐秘的知识并发现了一些财宝。",
		"成功打击了一名异教徒，获得了对方的一些物品。",
		"在贝克兰德一家著名报红冰淇淋店买了一只甜筒。正准备享用时却发现旁边夫妻怀中一位可爱的宝宝正含着拇指眼巴巴地望着你。鬼使神差中你把手中还未品尝的甜筒就这么直接递了过去……当晚你发现自己随手买的彩票中了大奖。",
		"和一位身材矮小的金发女士合作，成功抓捕了一名官方悬赏的在逃罪犯。",
		"和一位黑夜教会长相俊美的红手套先生合作，成功抓捕了一名官方悬赏的在逃罪犯。",
		"你搭乘的客船遇上了海盗袭击。千钧一发之际幽蓝复仇者号出现在海天交接之处，震慑走了向平民挥刀的歹徒们。你向蓝发的船长脱帽致意，获得经验以及海盗落下的金钱。",
	}},
	AdventureEvent{Type: 5, Probability: 15, Messages: []string{
		"完成一次非凡任务委托，获得报酬。",
		"猎杀了一只凶猛的非凡生物，获得材料。",
		"不小心踩死了一只虚弱的老鼠，获得了一份偷盗者非凡特性。拿着它总觉得自己会不太安全的样子…你将非凡特性转卖了。",
		"完成了丰收教堂发布的药材收购委托，获得报酬和乌特拉夫斯基神父的感谢。",
		"资助了拜亚姆混血儿慈善基金下的几处产业，结果居然在几个礼拜后获得了分红…？你想起现在的拜亚姆反抗军在由谁管理，不禁明白了什么…",
		"完成一次非凡任务委托，获得报酬。",
	}},

	AdventureEvent{Type: 6, Probability: 15, Messages: []string{
		"在一出失落的古堡中发现一本记载了神秘学知识的书籍，获得知识。",
		"参加非凡者组织的聚会，有幸的到了高阶阅读者的指点，获得知识。",
		"你遇见了贝克兰德最美丽的明珠奥黛丽·霍尔小姐，并有幸同她本人……的宠物苏茜玩耍了一会，获得幸福感。",
		"没时间解释了，赞美太阳！圣洁的日光中，",
		"有幸围观了疯狂冒险家格尔曼·斯帕罗提现海盗的全部过程。",
		"没时间解释了，赞美女神！寂静的绯红中，",
		"猎人的滋味真不错啊…你听了一整晚的单口相声。",
		"被知识教会的传教士拦下，被强制灌输了一段四皇之战的历史。",
		"和一名路边的猎人对骂，并成功取得了阶段性胜利。没有人可以在骂战里赢过祖安出生的你，没有人。",
		"君子爱财，取之有道。你背下了酒吧告示板上所有海盗的悬赏令，并期待以后能派上用场。",
		"从人贩子手里解救了一名土著小女孩，分别时她紧紧地拥抱了你。",
		"帮胖药师达克威尔抓住了卖假药的，分别时他送给你一包木乃伊粉作为感谢。",
		"和月亮一起打扫了丰收教堂，听他絮叨了半天乌特拉夫斯基神父有多可怕。",
		"成功催稿佛尔斯！天啊，你似乎完成了一件不可能做到的事。",
		"被值夜者怀疑而入梦检查，庆幸的是你只做了春梦。官方非凡者对你的信任度和了解程度提高了！然而本插件并没有势力好感这个属性……",
	}},
	AdventureEvent{Type: 7, Probability: 35, Messages: []string{
		"美好的一天，何不来点东拜朗甜姜红茶？",
		"糟糕的一天，并不想出门！",
		"糟糕的一天，你只想摸鱼！",
		"美好的一天，何不来点咕噜树树汁？",
		"美好的一天，何不来点冰镇淡啤酒？",
		"无聊的一天，你开始用空气子弹打蚊子，然而并没有打中。",
		"无聊的一天，你开始用火鸦打苍蝇，然而并没有打中。",
		"无聊的一天，你开始入梦路边的乌鸦，并发现它正准备计划着毁灭世界。",
		"无聊的一天，你开始入梦邻居家的黑猫，它说自己认识真正的黑皇帝。",
		"无聊的一天，你开始从灵界召唤特质为没用的灵体，结果并没有什么用。",
		"无聊的一天，你开始冥想传说中占卜家的神话形态，结果睁眼以后看什么都像虫子。",
		"斗邪恶。斗邪恶。斗邪恶。不不不，不能再这么懒散下去了。该努力扮演尽快晋升才对！……………………斗邪恶。斗邪恶。斗邪恶。",
		"灵性直觉忽然预警！你在自己家躲了一整天没敢出门，完全不知道自己究竟躲过了什么。不过在这个世界里，无知反而意味着安全。",
	}},
}

var menus = Menu{
	0,
	"帮助",
	"回复 帮助 可显示帮助信息。",
	"支持私聊查询，私聊格式为 指令@群号 其余详细介绍请见：https://github.com/molin0000/secretMaster/blob/master/README.md`",
	[]Menu{
		{7, "设置", "回复 设置 可查看插件高级设置。", "", []Menu{
			{71, ".master", "查看和设置管理员，设置本群的GM，格式为：.master;QQ号码，必须先设置管理员，才能使用其它设置指令, 查看指令.master。", "", nil},
			{72, "序列战争开", "在本群开启序列战争插件。", "", nil},
			{73, "序列战争关", "在本群关闭序列战争插件。", "", nil},
			{74, "货币映射", "映射序列战争货币到其它插件ini文件，格式为：货币映射;ini路径;cQQ;关键字。", "", nil},
			{75, "查看映射", "查看货币映射参数。", "", nil},
			{76, "货币升级", "货币映射后，使用此命令使映射生效。", "", nil},
			{77, "货币降级", "取消货币映射，还原为插件内部金镑使用。", "", nil},
			{78, "GM", "GM指令格式为：GM;exp;100;QQ号 目前支持exp,money,magic,god数值支持加和减。", "", nil},
			{79, "版本", "查询当前插件版本号。", "", nil},
		}},
		{1, "资料", "回复 资料 可显示与人物信息相关的指令。", "", []Menu{
			{11, "属性", "回复 属性 可查询当前人物的属性信息。", "", nil},
			{12, "尊名", "序列3以后可以自定义尊名显示，方法为 @机器人 尊名xxxxoooo。", "", nil},
			{13, "道具", "回复 道具 可查看当前人物持有的道具。", "", nil},
			{14, "技能", "回复 技能 可查看当前人物所有的技能。", "", nil},
			{15, "排行", "回复 排行 可查询当前群内的非凡者排行榜。", "", nil},
			{16, "查询", "回复 查询+排名 可查询排行榜上人物的属性信息。", "", nil},
			{17, "神位", "回复 神位 可查询当前22条序列的神位信息。", "", nil},
		}},
		{2, "序列", "回复 序列 可显示与途径、序列相关的指令。", "", []Menu{
			{21, "途径", "回复 途径 可查询途径列表。", "", nil},
			{22, "更换", "回复 更换+途径序号 可更改当前人物的非凡途径。", "", nil},
			{23, "晋升", "回复 晋升 可尝试晋升下一序列。（暂时不可用，自动晋升）", "", nil},
			{24, "自杀", "回复 自杀 可删除当前人物，重新开始。", "", nil},
		}},
		{3, "教会", "回复 教会 可显示与教会相关的指令。", "教会支持的技能包括：幸运光环，精力充沛，灵性协调。", []Menu{
			{31, "寻访", "回复 寻访 可查询当前已经创建的全部教会和隐秘组织。", "", nil},
			{32, "创建", "拥有至高权杖和1000金镑可创建教会或隐秘组织，格式为：创建;组织名称;组织简介;技能名称;入会费", "", nil},
			{33, "解散", "回复 解散;名称 可解散自己创建的教会或隐秘组织。", "", nil},
			{34, "加入", "回复 加入+名称 可加入现有的教会或隐秘组织。", "", nil},
			{35, "退出", "回复 退出+名称 可退出当前已经加入的组织。", "", nil},
			{36, "祈祷", "回复 祈祷 可消耗1份灵性材料，得到3天的神灵庇护。(可使用教会技能，信徒每次祈祷为教主增加10点经验)", "", nil},
		}},
		{4, "商店", "回复 商店 可查看神秘黑市的商品清单。", "欢迎光临星火之潮贝克兰德分店，请在浏览商品时，戴好您的斗篷和面具（找到想要购买的商品后，只需要回复购买加商品名称即可，如：购买探险卷轴）", []Menu{
			{41, "探险卷轴", "（100金镑）走内部渠道，可立即开始1次探险。", "", nil},
			{42, "灵性药剂", "（70金榜）不明物质酿造的灵性药剂，据说口感极差。", "", nil},
			{43, "至高权杖", "（1000金镑）这是成立非凡者组织或教会必不可少的证明物。", "", nil},
			{44, "灵性材料", "（50金榜）随机得到1份灵性材料，可用于发动非凡仪式。", "", nil},
			{45, "红剧场门票", "（200金镑100灵性）听说很多非凡者看完表演都升级了。", "", nil},
		}},
		{5, "探险", "回复 探险 可主动触发每日3次的奇遇探险经历。", "", nil},
		{6, "生活", "回复 娱乐 可查看娱乐菜单。", "", []Menu{
			{61, "银行", "可将金镑存入获得利息收益。指令格式为：银行;存款;3500 银行;查账 银行;取款;3500", "", nil},
			{62, "工作", "可选择一份工作获得稳定收益。指令格式为：工作;女装直播 工作;停止 （每日灵性-20，金镑+20，经验+10）", "", nil},
			{63, "钓鱼", "可在风暴之海进行垂钓。灵性-5，随机加金镑(0~20)。", "", nil},
			{64, "彩票", "极小概率中奖。金镑-2，1%几率赢得1000金镑。", "", nil},
		}},
	},
}

var skillList = []Skill{
	{0, "幸运光环", 1, 5},
	{1, "精力充沛", 1, 5},
	{2, "灵性协调", 1, 5},
}

var version = &Version{"序列战争", "v2.0.0", "2020-02-29"}
