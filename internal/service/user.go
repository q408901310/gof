package service

import (
	"context"
	"gof/internal/consts"
	"gof/internal/dao"
	"gof/internal/model"
	"gof/internal/model/entity"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sUser struct{}
)

var (
	insUser = sUser{}
)

func User() *sUser {
	return &insUser
}

func (s *sUser) Register(ctx context.Context, in *model.UserRegisterIn) error {
	if in.Confirm != in.Password {
		return gerror.New("两次密码不同")
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var user *entity.User
		if err := gconv.Struct(in, &user); err != nil {
			return err
		}
		if err := s.CheckPassportUnique(ctx, user.Passport); err != nil {
			return err
		}
		if err := s.CheckNicknameUnique(ctx, user.Name); err != nil {
			return err
		}
		user.Password = s.EncryptPassword(user.Password)
		_, err := dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

func (s *sUser) Login(ctx context.Context, in *model.UserLoginIn) (sessionId string, err error) {
	userEntity, err := s.GetUserByPassportAndPassword(
		ctx,
		in.Passport,
		s.EncryptPassword(in.Password),
	)
	if err != nil {
		return
	}
	if userEntity == nil {
		return sessionId, gerror.New(`账号或密码错误`)
	}
	//sessionId, err = Session().SetUser(ctx, userEntity)
	//if err != nil {
	//	return "", err
	//}
	r := g.RequestFromCtx(ctx)
	sessionId = r.Session.MustId()
	_, err = dao.User.Ctx(ctx).Update(g.Map{"token": sessionId}, "passport", in.Passport)
	if err != nil {
		return "", err
	}
	r.Cookie.SetHttpCookie(&http.Cookie{
		Name:     r.Server.GetSessionIdName(),
		Value:    sessionId,
		Secure:   true,
		SameSite: http.SameSiteNoneMode, // 自定义SameSite，配合secure一起使用
	})
	return
}

func (s *sUser) Guest(ctx context.Context) (err error) {
	r := g.RequestFromCtx(ctx)
	v, _ := r.Session.Get(consts.UserSessionKey)
	g.Dump(r.Cookie.GetSessionId())

	if !v.IsNil() {
		var user *entity.User
		_ = v.Struct(&user)
		g.Dump(user)
	}
	return
}

func (s *sUser) Server(ctx context.Context, server uint) error {
	row, err := dao.Server.Ctx(ctx).One("id", server)
	if err != nil {
		return gerror.New("区服不存在")
	}
	if row.IsEmpty() {
		return gerror.New("区服不存在")
	}
	//TODO 区服状态和人数判断
	return nil
}

func (s *sUser) GetUser(ctx context.Context) (user *entity.User) {
	r := g.RequestFromCtx(ctx)
	v, err := r.Session.Get(consts.UserSessionKey)
	if err != nil {
		return nil
	}
	if !v.IsNil() {
		_ = v.Struct(&user)
	}
	return
}

// GetUserByPassportAndPassword 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Passport: passport,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}

// GetUserByToken 根据token查询用户信息，一般用于验证账号是否已登录。
func (s *sUser) GetUserByToken(ctx context.Context, token string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		"token": token,
	}).Scan(&user)
	return
}

// CheckPassportUnique 检测给定的账号是否唯一
func (s *sUser) CheckPassportUnique(ctx context.Context, passport string) error {
	n, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Passport, passport).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`账号"%s"已被占用`, passport)
	}
	return nil
}

// CheckNicknameUnique 检测给定的昵称是否唯一
func (s *sUser) CheckNicknameUnique(ctx context.Context, nickname string) error {
	n, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Name, nickname).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`昵称"%s"已被占用`, nickname)
	}
	return nil
}

// EncryptPassword 将密码按照内部算法进行加密
func (s *sUser) EncryptPassword(password string) string {
	return gmd5.MustEncrypt(password)
}

// Logout 注销
func (s *sUser) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
