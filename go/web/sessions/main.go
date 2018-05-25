package main

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
)

func main() {
	fmt.Println(PassWord)
	store := sessions.NewCookieStore([]byte(PassWord))
	m := martini.Classic()
	m.Use(render.Renderer())
	store.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   0,
	})
	m.Use(sessions.Sessions("my_session", store))
	m.Use(sessionauth.SessionUser(GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/login"
	sessionauth.RedirectParam = "from"
	m.Get("/login", func(r render.Render) {
		r.HTML(200, "login", nil)
	})
	m.Post("/login", binding.Bind(User{}), func(session sessions.Session, postedUser User, r render.Render, req *http.Request) {
		// You should verify credentials against a database or some other mechanism at this point.
		// Then you can authenticate this session.
		fmt.Printf("USER IN: %+v\n", postedUser)
		err := sessionauth.AuthenticateSession(session, &postedUser)
		if err != nil {
			r.JSON(500, err)
		}

		params := req.URL.Query()
		redirect := params.Get(sessionauth.RedirectParam)
		r.Redirect(redirect)
		return
	})
	m.Get("/private", sessionauth.LoginRequired, func(r render.Render, user sessionauth.User) {
		fmt.Printf("SESSION: %+v\n", user)
		r.HTML(200, "private", user.(*User))
	})
	m.Get("/logout", sessionauth.LoginRequired, func(session sessions.Session, user sessionauth.User, r render.Render) {
		fmt.Printf("USER OUT: %+v\n", user)
		sessionauth.Logout(session, user)
		r.Redirect("/login")
	})
	m.RunOnAddr(":3001")
}
