package repository

const (
	findAllUsers    = `SELECT id, login, password FROM flower.\"user\";`
	createUserQuery = `INSERT INTO flower."user"(id, login, password) VALUES ($1, $2, $3);`
	updateUserQuery = `UPDATE \"user\"
						SET login = COALESCE(NULLIF($1, ''), login),
							password = COALESCE(NULLIF($2, ''), password)
						WHERE id = $3
						RETURNING *`
)
