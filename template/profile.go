package template

import "github.com/rezandry/kejarmimpi/models"

//Profile is making template for get data profile
func Profile(user *models.User) map[string]interface{} {

	data := map[string]interface{}{
		"id":           user.ID,
		"name":         user.Name,
		"photo":        user.Photo,
		"email":        user.Email,
		"job":          user.Job,
		"address":      user.Location,
		"status":       user.Headline,
		"bio":          user.Bio,
		"biography":    user.Biography,
		"education":    user.Education,
		"career":       user.Career,
		"otherCareer":  user.OtherCareer,
		"organization": user.Organization,
		"facebook":     user.Facebook,
		"instagram":    user.Instagram,
		"linkedin":     user.Linkedin,
		"twitter":      user.Twitter,
		"token":        user.Token,
	}
	return data
}
