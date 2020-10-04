package queryDict

const (
	GETALLUSER      = `select userID, userFirstname,userLastname,userAddress from tb_user where userStatus = 0`
	GETSPECIFICUSER = `select userID, userFirstname,userLastname,userAddress from tb_user where userStatus = 1 and userFirstname = ?`
	ADDNEWUSER      = `insert into tb_user values ( ?, ?, ?, ?, 0 )`
	UPDATEUSER      = `update tb_user set userFirstname = ?, userLastname = ?, userAddress = ? where userID = ?`
	DELETEUSER      = `update tb_user set userStatus = 1 where userID = ?`
)
