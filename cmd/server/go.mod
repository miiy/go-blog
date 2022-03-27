module goblog.com/cmd/server

go 1.17

replace (
	goblog.com/pkg => ../../pkg
	goblog.com/service/article => ../../service/article
	goblog.com/service/auth => ../../service/auth
	goblog.com/service/example => ../../service/example
	goblog.com/service/feedback => ../../service/feedback
	goblog.com/service/search => ../../service/search
	goblog.com/service/tag => ../../service/tag
	goblog.com/service/userpost => ../../service/userpost
	goblog.com/service/usertag => ../../service/usertag
	goblog.com/service/usertagpost => ../../service/usertagpost
)

require github.com/google/wire v0.5.0
