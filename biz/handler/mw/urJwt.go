package mw

type usJwtLogin struct {
	Mobile     string `form:"mobile,required" json:"mobile,required"`
	VerifyCode string `form:"verifyCode,required" json:"verifyCode,required"`
	VenueId    string `form:"venueId,required" json:"venueId,required"`
}
