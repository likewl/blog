package middleware

func GetImg() string {
	var Img string
	imgs := make(map[int]string)
	imgs[1] = "/static/images/头像1.png"
	imgs[2] = "/static/images/头像2.png"
	imgs[3] = "/static/images/头像3.png"
	imgs[4] = "/static/images/头像4.png"
	imgs[4] = "/static/images/头像5.png"
	imgs[4] = "/static/images/头像6.png"
	for _, img := range imgs {
		Img = img
		return Img
		break
	}
	return Img
}
