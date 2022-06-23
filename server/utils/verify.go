package utils

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify	   = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
	Banner             	   = Rules{"Title": {NotEmpty()}, "Cover": {NotEmpty()}, "Content": {NotEmpty()}}
	ListBanner             = Rules{"Page": {NotEmpty(),Ge("1")}, "PageSize": {NotEmpty(),Ge("1"),Le("5")}}
	InfoBanner             = Rules{"ID": {NotEmpty(),Ge("1")}}
	EditBanner             = Rules{"ID": {NotEmpty()}}
	ListBlog             = Rules{"Page": {NotEmpty(),Ge("1")}, "PageSize": {NotEmpty(),Ge("1"),Le("50")}}
	CreateComment		=Rules{"BlogId": {NotEmpty()},"Name": {NotEmpty()},"Message":{NotEmpty()}}
	ListComment		=Rules{"BlogId": {NotEmpty()},"Page": {NotEmpty(),Ge("1")}, "PageSize": {NotEmpty(),Ge("1"),Le("50")}}

	Article             	   = Rules{"Title": {NotEmpty()}, "Cover": {NotEmpty()}, "Content": {NotEmpty()}}
	ListArticle             = Rules{"Page": {NotEmpty(),Ge("1")}, "PageSize": {NotEmpty(),Ge("1"),Le("5")}}
	InfoArticle            = Rules{"ID": {NotEmpty(),Ge("1")}}
	EditArticle             = Rules{"ID": {NotEmpty()}}
)
