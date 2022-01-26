package repository

const (
	findAllUsers    = `SELECT id, login, password FROM flower.user;`
	createUserQuery = `INSERT INTO flower.user(id, login, password) VALUES ($1, $2, $3) RETURNING id;`
	updateUserQuery = `UPDATE flower.user
						SET login = COALESCE(NULLIF($1, ''), login),
							password = COALESCE(NULLIF($2, ''), password)
						WHERE id = $3
						RETURNING id, login, password;`
	getUserByIdQuery    = `SELECT id, login, password FROM flower.user WHERE id = $1;`
	getUserByLoginQuery = `SELECT id, login, password FROM flower.user WHERE login = $1;`
	deleteByIdQuery     = `DELETE FROM flower.user WHERE id = $1;`
	//updateUserLoginQuery = `UPDATE flower.\"user\" SET login = $1 WHERE id = $2`
)
