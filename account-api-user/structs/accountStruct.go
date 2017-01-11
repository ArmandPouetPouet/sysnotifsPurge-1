package structs

//User struct
type User struct {
	tableName struct{} `sql:"sql7_adm.iddb" `
	UserID    int      `sql:"IDU" json:"id"`
	FirstName string   `sql:"Prenom" json:"first_name"`
	LastName  string   `sql:"Nom" json:"last_name"`
	Gender    string   `sql:"Title" json:"gender"`
	NickName  string   `sql:"Nick" json:"nick_name"`
	Email     string   `json:"email"`
}
