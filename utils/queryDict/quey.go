package queryDict

const (
	GETALLUSER      = `select u.identityID, u.userName, u.birthDate, j.jobName,e.educationName from tbUser u inner join tbJob j on u.jobID = j.jobID inner join tbEducation e on u.educationID = e.educationID;`
	GETSPECIFICUSER = `select u.identityID, u.userName, u.birthDate, j.jobName,e.educationName from tbUser u inner join tbJob j on u.jobID = j.jobID inner join tbEducation e on u.educationID = e.educationID where identityID = ?;`
	ADDNEWUSER      = `insert into tbUser (identityID, userName, birthDate,jobID,educationID) values (?,?,?,?,?);`
	UPDATEUSER      = `update tbUser set userName = ?, birthDate = ?, jobID = ?, educationID = ? where identityID = ?;`
)
