package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/bitly/go-simplejson"
	"github.com/catsworld/cqhttp-go-sdk/cq"
	"github.com/catsworld/golib/nyastring"
	"github.com/catsworld/golib/random"
	"github.com/catsworld/qq-bot-api"
)

type BotConfig struct {
	Token        string
	APIEndpoint  string
	PollEndpoint string
	Master       string
}
type TomlConfig struct {
	QQ BotConfig
}

var (
	RootDir, ChangeLogText, VersionText, HelpText string
	Conf                                          TomlConfig
	Words                                         *simplejson.Json
	QQ                                            *qqbotapi.BotAPI
	err                                           error
)

func main() {
	if RootDir, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		log.Fatal(err)
	}
	if _, err = toml.DecodeFile(RootDir+"/config.toml", &Conf); err != nil {
		log.Fatal(err)
	}
	raw, err := ioutil.ReadFile(RootDir + "/words.json")
	if err != nil {
		log.Fatal(err)
	}
	if Words, err = simplejson.NewJson([]byte(raw)); err != nil {
		log.Fatal(err)
	}
	raw, err = ioutil.ReadFile(RootDir + "/CHANGELOG.md")
	if err != nil {
		log.Fatal(err)
	}
	ChangeLogText, VersionText = nyastring.GetChangeLogAndVersion(string(raw))
	raw, err = ioutil.ReadFile(RootDir + "/HELP.md")
	if err != nil {
		log.Fatal(err)
	}
	HelpText = nyastring.GetHelp(string(raw))

	QQ, err = qqbotapi.NewBotAPI(Conf.QQ.Token, Conf.QQ.APIEndpoint, Conf.QQ.PollEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("已登录:", QQ.Self.UserName)

	u := qqbotapi.NewUpdate(0)
	u.Timeout = 60
	Updates, err := QQ.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	for u := range Updates {
		if u.Message == nil {
			continue
		}
		OnMessage(u.Message)
	}
}
func GetWordByString(w *simplejson.Json, s string, r string) string {
	p := w.Get(s)
	l := len(p.MustArray())
	for i := 0; i < l; i++ {
		v := p.GetIndex(i)
		if v.Get("if").MustString() == r {
			return v.Get("word").MustString()
		}
	}
	return ""
}

func OnMessage(r *qqbotapi.Message) {
	log.Println(r.Text)

	args := Message2Args(r)
	if len(args) == 0 {
		return
	}

	mt := ""
	if args[0] == "/botlog" {
		ChangeLog(r)
	} else if args[0] == "/botver" {
		Version(r)
	} else if args[0] == "/bothelp" {
		Help(r)
	} else if args[0] == "/蛮神心脏" {
		mt = "1.5.1 新宿御苑 均160AP"
	} else if args[0] == "/凤凰羽毛" {
		mt = "周日 剑阶上级本 均175.4AP\n第三章 丰饶之海 均193.5AP"
	} else if args[0] == "/世界树之种" {
		mt = "第三章 丰饶之海 均56.9AP"
	} else if args[0] == "/英雄之证" {
		mt = "第三章 海盗船 均20.1AP"
	} else if args[0] == "/凶骨" {
		mt = "序章 冬木-X-C 均21.7AP"
	} else if args[0] == "/龙之牙" {
		mt = "第五章 德明 均27.1AP"
	} else if args[0] == "/虚影之尘" {
		mt = "第五章 夏洛特 均31.4AP"
	} else if args[0] == "/愚者之锁" {
		mt = "第六章 死之荒野 均29.7AP"
	} else if args[0] == "/万死的毒针" {
		mt = "第七章 芦苇原 均33.2AP"
	} else if args[0] == "/魔术髓液" {
		mt = "1.5.1 新宿站 均32.5AP"
	} else if args[0] == "/鬼魂提灯" {
		mt = "第七章 库萨 均52.8AP"
	} else if args[0] == "/八连双晶" {
		mt = "第六章 圣都市街 均74.1AP"
	} else if args[0] == "/蛇之宝玉" {
		mt = "第七章 巴比伦 均78.4AP"
	} else if args[0] == "/无间齿轮" {
		mt = "1.5.1 枪身塔 均45.1AP"
	} else if args[0] == "/禁断书页" {
		mt = "1.5.1 新宿二丁目 均68.0AP"
	} else if args[0] == "/陨蹄铁" {
		mt = "第六章 无之大地 均50.6AP"
	} else if args[0] == "/大骑士勋章" {
		mt = "第六章 卡美洛 王城"
	} else if args[0] == "/追忆的贝壳" {
		mt = "第七章 观测所 均51.3AP"
	} else if args[0] == "/混沌之爪" {
		mt = "第五章 得梅因 均88.7AP"
	} else if args[0] == "/龙之逆鳞" {
		mt = "第七章 尼普尔 均159.1AP"
	} else if args[0] == "/精灵根" {
		mt = "第六章 圣都市街 均163.9AP"
	} else if args[0] == "/战马的幼角" {
		mt = "第六章 东之村 均109.2AP"
	} else if args[0] == "/血之泪石" {
		mt = "1.5.1 新宿二丁目 均114.1AP"
	} else if args[0] == "/黑兽脂" {
		mt = "第七章 北之高台 均89.7AP"
	} else if args[0] == "/封魔之灯" {
		mt = "第六章 隐秘之村 均121.4AP"
	} else if args[0] == "/智慧之圣甲虫像" {
		mt = "第六章 大神殿 均203.7AP"
	} else if args[0] == "/原初胎毛" {
		mt = "第七章 鲜血神殿 均108.2AP"
	} else if args[0] == "/咒兽胆石" {
		mt = "第七章 艾比夫山 均173.6AP"
	} else if args[0] == "/国服活动" {
		mt = "目录\n/活动奖励\n/奖励礼装及从者\n/副本介绍\n/攻略建议\n/卡池分析\n/外链网址\n请输入目录名称获取详细资讯w"
	} else if args[0] == "/活动奖励" {
		mt = nyastring.GetRdWord(Words, "REWARD")
	} else if args[0] == "/奖励礼装及从者" {
		mt = nyastring.GetRdWord(Words, "SERVANT FREE")
	} else if args[0] == "/副本介绍" {
		mt = nyastring.GetRdWord(Words, "Carbon")
	} else if args[0] == "/攻略建议" {
		mt = nyastring.GetRdWord(Words, "Strategy")
	} else if args[0] == "/卡池分析" {
		mt = nyastring.GetRdWord(Words, "Gacha")
	} else if args[0] == "/外链网址" {
		mt = "http://bbs.nga.cn/read.php?tid=11332208"
	} else if args[0] == "/英灵立绘"{
		for len(args[1])<4{
			args[1] = "0" + args[1]
		}
		mt = "[CQ:image,file=file://C:\\picture\\" + args[1] + ".png]"
		// fmt.Sprintf("%03d",random.Random(1,100))
		log.Println(mt)
	} else if args[0] == "/骗氪"{
		mt = "[CQ:image,file=file://C:\\PK\\FGO\\ " + fmt.Sprintf("(%d)",random.Random(1,442)) + ".png]"
		log.Println(mt)
	} else if args[0] == "/召唤"{
		var i int = random.Random(1,10);
		if i <= 6{
			mt = "[CQ:image,file=file://C:\\3\\ " + fmt.Sprintf("(%d)",random.Random(1,32)) + ".png]"
			log.Println(mt)
		}else if i <= 9{
			mt = "[CQ:image,file=file://C:\\4\\ " + fmt.Sprintf("(%d)",random.Random(1,61)) + ".png]"
			log.Println(mt)
		}else{
			mt = "[CQ:image,file=file://C:\\5\\ " + fmt.Sprintf("(%d)",random.Random(1,59)) + ".png]"
			log.Println(mt)
		}
	}
	m := qqbotapi.NewMessage(r.Chat.ID, r.Chat.Type, mt)
	QQ.Send(m)
}
func Help(r *qqbotapi.Message) {
	args := Message2Args(r)
	if len(args) > 0 && args[0] == "/khkhelp" {
		m := qqbotapi.NewMessage(r.Chat.ID, r.Chat.Type, VersionText)
		QQ.Send(m)
	}
}

func Version(r *qqbotapi.Message) {
	args := Message2Args(r)
	if len(args) > 0 && args[0] == "/khkver" {
		m := qqbotapi.NewMessage(r.Chat.ID, r.Chat.Type, VersionText)
		QQ.Send(m)
	}
}

func ChangeLog(r *qqbotapi.Message) {
	args := Message2Args(r)
	if len(args) > 0 && args[0] == "/khklog" {
		m := qqbotapi.NewMessage(r.Chat.ID, r.Chat.Type, ChangeLogText+nyastring.GetRdWord(Words, "changelogegg"))
		QQ.Send(m)
	}
}

func Message2Args(r *qqbotapi.Message) []string {
	c := r.Text
	ret := nyastring.SplitCommand(c)
	if len(ret) > 0 {
		ret[0] = strings.Replace(ret[0], cq.At(QQ.Self.UserName), "", -1)
	}
	return ret
}
