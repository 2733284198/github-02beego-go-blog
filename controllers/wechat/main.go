package wechat

import (
	"fmt"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

type MainController struct {
	BaseController
}

func (c *MainController) Hello() {

	wc := wechat.NewWechat(config)

	log.Infoln(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Ctx.Request, c.Ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		log.Infoln(msg)
		return ResponseMsgType(msg)

		//text := message.Reply{}.MsgData
		/*switch msg.MsgType {
			//文本消息
			case message.MsgTypeText:
				//回复消息：演示回复用户发送的消息
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

				//图片消息
			case message.MsgTypeImage:
				//do something
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

				//语音消息
			case message.MsgTypeVoice:
				//do something
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

				//视频消息
			case message.MsgTypeVideo:
				//do something
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

				//小视频消息
			case message.MsgTypeShortVideo:
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
				//do something

				//地理位置消息
			case message.MsgTypeLocation:
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
				//do something

				//链接消息
			case message.MsgTypeLink:
				//do something
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

				//事件推送消息
			case message.MsgTypeEvent:

				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			default:
				text := message.NewText(msg.Content)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}*/

		//return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()






	/*signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	token := "b8cf671eaa1a270a9b53ddb894dd9029"
	echostr := c.GetString("echostr")

	var tempArray  = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature对比
	if sha1String == signature {
		_, err := c.Ctx.ResponseWriter.Write([]byte(echostr))
		if err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			c.StopRun()
		}
		//fmt.Print(echostr)
	} else {
		c.Data["json"] = "验证失败"
		c.ServeJSON()
		c.StopRun()
		//fmt.Println("验证失败")
	}*/


	/*c.Data["json"] = echostr
	c.ServeJSON()
	c.StopRun()*/
	//配置微信参数
	/*c.EnableRender = false
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {


		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()*/
}