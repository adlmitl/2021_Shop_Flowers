package repository

// Queries to db psql.
const (
	findAllFlowers    = `SELECT id, name, price FROM flower.flower;`
	createFlowerQuery = `INSERT INTO flower.flower(id, name, price) VALUES ($1, $2, $3) RETURNING id;`
	updateFlowerQuery = `UPDATE flower.flower
						SET name = COALESCE(NULLIF($2, ''), name),
							price = $3
						WHERE id = $1
						RETURNING id, name, price;`
	getFlowerByIdQuery    = `SELECT id, name, price FROM flower.flower WHERE id = $1;`
	getFlowerByNameQuery  = `SELECT id, name, price FROM flower.flower WHERE name = $1;`
	deleteFlowerByIdQuery = `DELETE FROM flower.flower WHERE id = $1;`
)
