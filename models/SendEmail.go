package models

import (
	"encoding/json"
	"fmt"
	"github.com/FREE-WE-1/backend/global"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

func SendEmailValidate(em []string) (string, error) {
	e := email.NewEmail()
	e.From = fmt.Sprintf("啊对对队 <2755137031@qq.com>")
	e.To = em
	// 生成6位随机验证码
	e.Subject = "找回密码"
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")
	//设置文件发送的内容
	content := fmt.Sprintf(`
	<div>
		<div>
			尊敬的 %s，您好！
		</div>
		<div style="padding: 8px 40px 8px 40px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
	`, em[0], t, vCode)
	e.HTML = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2755137031@qq.com", "hngwgsbfpnuddfbb", "smtp.qq.com"))
	return vCode, err
}

func GetValidateCode(c *gin.Context, id int, em []string) int32 {
	// 获取目的邮箱
	vCode, err := SendEmailValidate(em)
	if err != nil {
		log.Println(err)
		return 400
	}
	// 验证码存入redis 并设置过期时间5分钟
	//_ = global.RedisClient.Do(global.RedisClient.Context(), "set", "vCode", vCode)
	//_ = global.RedisClient.Do(global.RedisClient.Context(), "expire", "vCode", 300)
	var ret EmailToken
	ret.Token = string(vCode)

	var jsonStr []byte
	jsonStr, err = json.Marshal(ret)
	if err != nil {
		panic(err) // never happens
	}
	err = global.RedisClient.Set(c, string(id), jsonStr, 0).Err()
	if err != nil {
		panic(err)
	}

	//global.RedisClient.Set(c, string(id), vCode, 0)
	//c.Set("email-token", string(vCode))
	if err != nil {
		log.Println(err.Error())
		return 400
	}
	return 200
}
