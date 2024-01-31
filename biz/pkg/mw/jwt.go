/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mw

import (
	"context"
	"errors"
	"fmt"
	"hertz_admin/biz/dal/entity"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "id"
)

type User struct {
	Id          float64
	Name        string
	Permissions []interface{}
}

func InitJwt() {
	var err error

	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": "success",
				"data": utils.H{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				},
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Account  string `form:"mobile" json:"mobile" query:"mobile" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			sysUser, err := entity.QuerySysUserByMobile(loginStruct.Account)
			if err != nil {
				return nil, err
			}
			if sysUser.Password != loginStruct.Password {
				return nil, errors.New("密码不正确")
			}

			menus, _, err := entity.QuerySysMenu(nil, 1, 1000)

			var list []interface{}
			for _, menu := range menus {
				list = append(list, menu.ApiUrl)
			}

			return &User{
				Id:          float64(sysUser.Id),
				Name:        sysUser.RealName,
				Permissions: list,
			}, nil
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			user := data.(*User)
			url := string(c.Request.Path())

			flag := false
			for _, permission := range user.Permissions {
				if fmt.Sprint(permission) == url {
					flag = true
					break
				}
			}
			return flag
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				Id:          claims[IdentityKey].(float64),
				Name:        claims["name"].(string),
				Permissions: claims["permissions"].([]interface{}),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					IdentityKey:   v.Id,
					"name":        v.Name,
					"permissions": v.Permissions,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
