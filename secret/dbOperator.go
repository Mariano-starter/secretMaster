package secret

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/molin0000/secretMaster/rlp"
	"github.com/syndtr/goleveldb/leveldb"
	sc "golang.org/x/text/encoding/simplifiedchinese"
	"gopkg.in/ini.v1"
)

var db *leveldb.DB

func getDb() *leveldb.DB {
	if db == nil {
		_db, err := leveldb.OpenFile("secret.db", nil)
		if err != nil {
			fmt.Printf("open db error: %+v", err)
			panic(err)
		}
		db = _db
	}
	return db
}

func (b *Bot) keys(fromQQ uint64) []byte {
	return []byte(strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(fromQQ), 10))
}

func (b *Bot) ruleKey(fromQQ uint64) []byte {
	return []byte(strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(fromQQ), 10) + "_water")
}

func (b *Bot) moneyKey(fromQQ uint64) []byte {
	return []byte(strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(fromQQ), 10) + "_money")
}

func (b *Bot) advKey(fromQQ uint64) []byte {
	return []byte(strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(fromQQ), 10) + "_adventure")
}

func (b *Bot) godKey(secretID uint64) []byte {
	return []byte("god_" + strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(secretID), 10))
}

func (b *Bot) rnameKey(fromQQ uint64) []byte {
	return []byte("rname_" + strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(fromQQ), 10))
}

func (b *Bot) externKey(fromQQ uint64) []byte {
	return []byte("extern_" + strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(fromQQ), 10))
}

func (b *Bot) switchKey() []byte {
	return []byte("switch_" + strconv.FormatInt(int64(b.Group), 10))
}

func (b *Bot) getKeyPrefix() []byte {
	return []byte(strconv.FormatInt(int64(b.Group), 10) + "_")
}

func (b *Bot) getSwitch() bool {
	verify, err := getDb().Get(b.switchKey(), nil)
	if err != nil {
		return true
	}
	var v BotSwitch
	rlp.DecodeBytes(verify, &v)
	return v.Enable
}

func (b *Bot) setSwitch(enable bool) {
	p := &BotSwitch{b.Group, enable}
	buf, _ := rlp.EncodeToBytes(p)
	getDb().Put(b.switchKey(), buf, nil)
}

func (b *Bot) getPersonFromDb(fromQQ uint64) *Person {
	verify, _ := getDb().Get(b.keys(fromQQ), nil)
	var v Person
	rlp.DecodeBytes(verify, &v)
	return &v
}

func (b *Bot) setPersonToDb(fromQQ uint64, p *Person) {
	buf, _ := rlp.EncodeToBytes(p)
	getDb().Put(b.keys(fromQQ), buf, nil)
}

func (b *Bot) getWaterRuleFromDb(fromQQ uint64) *WaterRule {
	ret, err := getDb().Get(b.ruleKey(fromQQ), nil)
	if err != nil {
		return &WaterRule{Group: b.Group, QQ: fromQQ, DayCnt: 1, Days: uint64(time.Now().Unix() / (3600 * 24))}
	}
	var w WaterRule
	rlp.DecodeBytes(ret, &w)
	if w.Days != uint64(time.Now().Unix()/(3600*24)) {
		w.DayCnt = 0
		w.Days = uint64(time.Now().Unix() / (3600 * 24))
	}
	return &w
}

func (b *Bot) setWaterRuleToDb(fromQQ uint64, w *WaterRule) {
	buf, _ := rlp.EncodeToBytes(w)
	getDb().Put(b.ruleKey(fromQQ), buf, nil)
}

func (b *Bot) getMoneyFromDb(fromQQ uint64, chatCnt uint64) *Money {
	bind := b.getMoneyBind()
	if bind.HasUpdate && fileExists(bind.IniPath) {
		fmt.Println("get from ini.")
		return b.getMoneyFromIni(fromQQ)
	}
	fmt.Println("get from db.")

	ret, err := getDb().Get(b.moneyKey(fromQQ), nil)
	fmt.Println(ret, err)

	if err != nil {
		return &Money{Group: b.Group, QQ: fromQQ, Money: chatCnt}
	}
	var m Money
	rlp.DecodeBytes(ret, &m)
	fmt.Println(m)
	return &m
}

func (b *Bot) setMoneyToDb(fromQQ uint64, m *Money) {
	bind := b.getMoneyBind()
	if bind.HasUpdate {
		fmt.Printf("QQ:%d, Money:%d", m.QQ, m.Money)
		b.setMoneyToIni(fromQQ, m)
		return
	}
	buf, _ := rlp.EncodeToBytes(m)
	getDb().Put(b.moneyKey(fromQQ), buf, nil)
}

func cString(str string) string {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return gbstr
}

func goString(str string) string {
	utf8str, _ := sc.GB18030.NewDecoder().String(str)
	return utf8str
}

func (b *Bot) getMoneyFromIni(fromQQ uint64) *Money {
	bind := b.getMoneyBind()
	cfg, err := ini.Load(bind.IniPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return &Money{Group: b.Group, QQ: fromQQ, Money: 0}
	}

	fmt.Println("Data Path:", cfg.Section(cString(strconv.FormatUint(fromQQ, 10))).Key(cString(bind.IniKey)).String())
	amount, err := cfg.Section(strconv.FormatUint(fromQQ, 10)).Key(cString(bind.IniKey)).Uint64()
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return &Money{Group: b.Group, QQ: fromQQ, Money: 0}
	}

	return &Money{Group: b.Group, QQ: fromQQ, Money: amount}
}

func (b *Bot) setMoneyToIni(fromQQ uint64, m *Money) {
	bind := b.getMoneyBind()
	cfg, err := ini.Load(bind.IniPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	fmt.Printf("QQ:%d, Money:%d", m.QQ, m.Money)
	cfg.Section(cString(strconv.FormatUint(fromQQ, 10))).Key(cString(bind.IniKey)).SetValue(strconv.FormatUint(m.Money, 10))
	cfg.SaveTo(bind.IniPath)
}

func (b *Bot) getAdvFromDb(fromQQ uint64) *Adventure {
	ret, err := getDb().Get(b.advKey(fromQQ), nil)
	if err != nil {
		return &Adventure{Group: b.Group, QQ: fromQQ, Days: uint64(time.Now().Unix() / (3600 * 24)), DayCnt: 0}
	}
	var m Adventure
	rlp.DecodeBytes(ret, &m)
	if m.Days != uint64(time.Now().Unix()/(3600*24)) {
		m.DayCnt = 0
	}
	return &m
}

func (b *Bot) setAdvToDb(fromQQ uint64, m *Adventure) {
	buf, _ := rlp.EncodeToBytes(m)
	getDb().Put(b.advKey(fromQQ), buf, nil)
}

func (b *Bot) getGodFromDb(secretID uint64) uint64 {
	ret, err := getDb().Get(b.godKey(secretID), nil)
	if err != nil {
		return 0
	}
	var god uint64
	rlp.DecodeBytes(ret, &god)
	return god
}

func (b *Bot) setGodToDb(secretID uint64, god *uint64) {
	buf, _ := rlp.EncodeToBytes(god)
	getDb().Put(b.godKey(secretID), buf, nil)
}

func (b *Bot) setRNameToDb(fromQQ uint64, r *RespectName) {
	buf, _ := rlp.EncodeToBytes(r)
	getDb().Put(b.rnameKey(fromQQ), buf, nil)
}

func (b *Bot) getRNameFromDb(fromQQ uint64) string {
	ret, err := getDb().Get(b.rnameKey(fromQQ), nil)
	if err != nil {
		return "无"
	}
	var god RespectName
	rlp.DecodeBytes(ret, &god)
	return "\n" + god.Name
}

func (b *Bot) setExternToDb(fromQQ uint64, r *ExternProperty) {
	buf, _ := rlp.EncodeToBytes(r)
	getDb().Put(b.externKey(fromQQ), buf, nil)
}

func (b *Bot) getExternFromDb(fromQQ uint64) *ExternProperty {
	ret, err := getDb().Get(b.externKey(fromQQ), nil)
	if err != nil {
		return &ExternProperty{Luck: 0, Magic: 200, Days: uint64(time.Now().Unix() / 3600 / 24)}
	}
	var v ExternProperty
	rlp.DecodeBytes(ret, &v)
	if v.Days != uint64(time.Now().Unix()/(3600*24)) {
		v.Magic = 200
		v.Days = uint64(time.Now().Unix() / (3600 * 24))
	}
	fmt.Printf("getExternFromDb:%+v", v)
	tree := b.getPersonValue("SkillTree", fromQQ, &SkillTree{}).(*SkillTree)
	fmt.Printf("tree:%+v", tree)

	for i := 0; i < len(tree.Skills); i++ {
		if tree.Skills[i] != nil && tree.Skills[i].Name == skillList[0].Name {
			v.Luck = tree.Skills[i].Level
			break
		}
	}

	return &v
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (b *Bot) moneyBindKey() []byte {
	return []byte("moneyBind_" + strconv.FormatInt(int64(b.Group), 10))
}

func (b *Bot) getMoneyBind() *MoneyBind {
	verify, err := getDb().Get(b.moneyBindKey(), nil)
	if err != nil {
		return &MoneyBind{"/home/user/coolq/data/app/com.qmt.demo/玩家数据.ini", "cQQ", "金币数量", false}
	}
	var v MoneyBind
	rlp.DecodeBytes(verify, &v)
	return &v
}

func (b *Bot) setMoneyBind(p *MoneyBind) {
	buf, _ := rlp.EncodeToBytes(p)
	getDb().Put(b.moneyBindKey(), buf, nil)
}

// Common DB operate---------

func (b *Bot) groupKey(keyPrefix string) []byte {
	return []byte(keyPrefix + "_" + strconv.FormatInt(int64(b.Group), 10))
}

func (b *Bot) setGroupValue(keyPrefix string, p interface{}) {
	buf, _ := rlp.EncodeToBytes(p)
	getDb().Put(b.groupKey(keyPrefix), buf, nil)
}

func (b *Bot) getGroupValue(keyPrefix string, defaultValue interface{}) interface{} {
	data, err := getDb().Get(b.groupKey(keyPrefix), nil)
	if err != nil {
		return defaultValue
	}

	rlp.DecodeBytes(data, defaultValue)
	return defaultValue
}

func (b *Bot) personKey(keyPrefix string, fromQQ uint64) []byte {
	return []byte(keyPrefix + "_" + strconv.FormatInt(int64(b.Group), 10) + "_" + strconv.FormatInt(int64(fromQQ), 10))
}

func (b *Bot) setPersonValue(keyPrefix string, fromQQ uint64, p interface{}) {
	buf, err := rlp.EncodeToBytes(p)
	if err != nil {
		fmt.Println(err)
	}
	getDb().Put(b.personKey(keyPrefix, fromQQ), buf, nil)
}

func (b *Bot) getPersonValue(keyPrefix string, fromQQ uint64, defaultValue interface{}) interface{} {
	data, err := getDb().Get(b.personKey(keyPrefix, fromQQ), nil)
	fmt.Printf("getPersonValue, %s, %d, %+v, %v", keyPrefix, fromQQ, data, err)
	if err != nil {
		fmt.Println("数据库中没有这个字段:", b.personKey(keyPrefix, fromQQ))
		return defaultValue
	}

	rlp.DecodeBytes(data, defaultValue)
	return defaultValue
}
