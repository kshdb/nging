/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package common

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/webx-top/captcha"
	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	stdCode "github.com/webx-top/echo/code"
	hdlCaptcha "github.com/webx-top/echo/handler/captcha"
	"github.com/webx-top/echo/middleware/render"
	"github.com/webx-top/echo/middleware/tplfunc"
	"github.com/webx-top/echo/subdomains"
)

var ErrorProcessors = []render.ErrorProcessor{
	func(ctx echo.Context, err error) (processed bool, newErr error) {
		if errors.Is(err, db.ErrNoMoreRows) {
			return true, echo.NewError(ctx.T(`数据不存在`), stdCode.DataNotFound)
		}
		return false, err
	},
}

func ProcessError(ctx echo.Context, err error) error {
	for _, processor := range ErrorProcessors {
		if processor == nil {
			continue
		}
		var processed bool
		processed, err = processor(ctx, err)
		if processed {
			break
		}
	}
	return err
}

// Ok 操作成功
func Ok(v string) Successor {
	return NewOk(v)
}

// Err 获取错误信息
func Err(ctx echo.Context, err error) (ret interface{}) {
	if err == nil {
		flash := ctx.Flash()
		if flash != nil {
			if errMsg, ok := flash.(string); ok {
				ret = errors.New(errMsg)
			} else {
				ret = flash
			}
		}
	} else {
		ret = ProcessError(ctx, err)
	}
	return
}

// SendOk 记录成功信息
func SendOk(ctx echo.Context, msg string) {
	if ctx.IsAjax() || ctx.Format() != `html` {
		ctx.Data().SetInfo(msg, 1)
		return
	}
	ctx.Session().AddFlash(Ok(msg))
}

// SendFail 记录失败信息
func SendFail(ctx echo.Context, msg string) {
	if ctx.IsAjax() || ctx.Format() != `html` {
		ctx.Data().SetInfo(msg, 0)
		return
	}
	ctx.Session().AddFlash(msg)
}

// SendErr 记录错误信息 (SendFail的别名)
func SendErr(ctx echo.Context, err error) {
	err = ProcessError(ctx, err)
	SendFail(ctx, err.Error())
}

func GenCaptchaError(ctx echo.Context, hostAlias string, captchaName string, id string, args ...string) echo.Data {
	data := ctx.Data()
	data.SetZone(captchaName)
	data.SetData(CaptchaInfo(hostAlias, captchaName, id, args...))
	data.SetError(ErrCaptcha)
	return data
}

// VerifyCaptcha 验证码验证
func VerifyCaptcha(ctx echo.Context, hostAlias string, captchaName string, args ...string) echo.Data {
	idGet := ctx.Form
	idSet := func(id string) {
		ctx.Request().Form().Set(`captchaId`, id)
	}
	if len(args) > 0 {
		idGet = func(_ string, defaults ...string) string {
			return ctx.Form(args[0], defaults...)
		}
		idSet = func(id string) {
			ctx.Request().Form().Set(args[0], id)
		}
	}
	code := ctx.Form(captchaName)
	id := idGet("captchaId")
	if len(id) == 0 {
		return GenCaptchaError(ctx, hostAlias, captchaName, id, args...)
	}
	exists := captcha.Exists(id)
	if len(code) == 0 {
		return GenCaptchaError(ctx, hostAlias, captchaName, ``, args...)
	}
	if !exists && len(hdlCaptcha.DefaultOptions.CookieName) > 0 {
		id = ctx.GetCookie(hdlCaptcha.DefaultOptions.CookieName)
		if len(id) > 0 {
			if exists = captcha.Exists(id); exists {
				idSet(id)
			}
		}
	}
	if !exists && len(hdlCaptcha.DefaultOptions.HeaderName) > 0 {
		id = ctx.Header(hdlCaptcha.DefaultOptions.HeaderName)
		if len(id) > 0 {
			if exists = captcha.Exists(id); exists {
				idSet(id)
			}
		}
	}
	if !exists || !tplfunc.CaptchaVerify(code, idGet) {
		return GenCaptchaError(ctx, hostAlias, captchaName, ``, args...)
	}
	return ctx.Data()
}

// VerifyAndSetCaptcha 验证码验证并设置新验证码信息
func VerifyAndSetCaptcha(ctx echo.Context, hostAlias string, captchaName string, args ...string) echo.Data {
	data := VerifyCaptcha(ctx, hostAlias, captchaName, args...)
	if data.GetCode() != stdCode.CaptchaError {
		data.SetData(CaptchaInfo(hostAlias, captchaName, ``, args...))
	}
	return data
}

// CaptchaInfo 新验证码信息
func CaptchaInfo(hostAlias string, captchaName string, captchaID string, args ...string) echo.H {
	if len(captchaID) == 0 {
		captchaID = captcha.New()
	}
	captchaIdent := `captchaId`
	if len(args) > 0 {
		captchaIdent = args[0]
	}
	return echo.H{
		`captchaName`:  captchaName,
		`captchaIdent`: captchaIdent,
		`captchaID`:    captchaID,
		`captchaURL`:   subdomains.Default.URL(`/captcha/`+captchaID+`.png`, hostAlias),
	}
}

type ConfigFromDB interface {
	ConfigFromDB() echo.H
}

func MakeMap(values ...interface{}) echo.H {
	h := echo.H{}
	if len(values) == 0 {
		return h
	}
	var k string
	for i, j := 0, len(values); i < j; i++ {
		if i%2 == 0 {
			k = fmt.Sprint(values[i])
			continue
		}
		h.Set(k, values[i])
		k = ``
	}
	if len(k) > 0 {
		h.Set(k, nil)
	}
	return h
}

// GetLocalIP 获取本机网卡IP
func GetLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIPNet bool
	)
	ipv4 = `127.0.0.1`
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIPNet = addr.(*net.IPNet); isIPNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	return
}

var notWordRegexp = regexp.MustCompile(`[^\w]+`)

// LookPath 获取二进制可执行文件路径
func LookPath(bin string, otherPaths ...string) (string, error) {
	envVarName := `NGING_` + notWordRegexp.ReplaceAllString(strings.ToUpper(bin), `_`) + `_PATH`
	envVarValue := os.Getenv(envVarName)
	if len(envVarValue) > 0 {
		if com.IsFile(envVarValue) {
			return envVarValue, nil
		}
		envVarValue = filepath.Join(envVarValue, bin)
		if com.IsFile(envVarValue) {
			return envVarValue, nil
		}
	}
	findPath, err := exec.LookPath(bin)
	if err == nil {
		return findPath, err
	}
	if !errors.Is(err, exec.ErrNotFound) {
		return findPath, err
	}
	for _, binPath := range otherPaths {
		binPath = filepath.Join(binPath, bin)
		if com.IsFile(binPath) {
			return binPath, nil
		}
	}
	return findPath, err
}
