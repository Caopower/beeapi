package controllers

type IndexController struct {
	BaseController
}
// @router / [get]
func (this *IndexController) Index() {
	this.TplName="index.html"
}
// @router /follow [get]
func (this *IndexController) IndexFollow() {
	this.TplName="follow.html"
}
// @router /message [get]
func (this *IndexController) IndexMessage() {
	this.TplName="message.html"
}
//@router /details [get]
func (this *IndexController) IndexDetails(){
	this.TplName="details.html"
}
//@router /comment [get]
func (this *IndexController) IndexComment(){
	this.TplName="comment.html"
}
//@router /info [get]
func (this *IndexController) IndexInfo(){
	this.TplName="info.html"
}