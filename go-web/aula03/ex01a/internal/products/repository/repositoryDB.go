package products

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type repositoryDbProduct struct {
	db *sql.DB
}

func NewRepositoryDB(db *sql.DB) Repository {
	return &repositoryDbProduct{db: db}
}

func (r *repositoryDbProduct) GetAll() ([]Product, error) {

	var ps []Product

	rows, err := r.db.Query("SELECT * FROM products")
	if err != nil {
		return ps, err
	}

	defer rows.Close()

	for rows.Next() {
		var product Product

		err := rows.Scan(&product.ID, &product.Name, &product.Tipo, &product.Count, &product.Price)
		if err != nil {
			return ps, err
		}

		ps = append(ps, product)
	}

	return ps, nil
}

func (r *repositoryDbProduct) GetId(id int) (Product, error) {

	var ps Product

	stmt, err := r.db.Prepare("SELECT * FROM products Where id = ?")
	if err != nil {
		return ps, err
	}
	err = stmt.QueryRow(id).Scan(&ps.ID,
		&ps.Name,
		&ps.Tipo,
		&ps.Count,
		&ps.Price)
	defer stmt.Close()

	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (r *repositoryDbProduct) LastID() (int, error) {

	stmt, err := r.db.Prepare("Select * from products")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	lid, err := res.LastInsertId()
	return int(lid), err
}

func (r *repositoryDbProduct) Store(id int, name, tipo string, count int, price float64) (Product, error) {

	stmt, err := r.db.Prepare("INSERT INTO products (name, typee, count, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, tipo, count, price)
	if err != nil {
		return Product{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return Product{}, fmt.Errorf("produto %d não cadastrado", id)
	}

	lid, _ := res.LastInsertId()
	ps, _ := r.GetId(int(lid))
	return ps, nil
}

func (r *repositoryDbProduct) Update(id int, name, tipo string, count int, price float64) (Product, error) {

	stmt, err := r.db.Prepare("UPDATE products SET name=? , typee=?, count=?, price = ? WHERE id=?")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, tipo, count, price, id)
	if err != nil {
		return Product{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	ps, _ := r.GetId(id)
	return ps, nil
}

func (r *repositoryDbProduct) UpdateName(id int, name string) (Product, error) {

	stmt, err := r.db.Prepare("UPDATE products SET name=? WHERE id=?")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, id)

	if err != nil {
		return Product{}, err
	}
	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}

	ps, _ := r.GetId(id)

	return ps, nil
}

func (r *repositoryDbProduct) Delete(id int) error {

	stmt, err := r.db.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close() // Impedir vazamento de memória

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf("produto %d não encontrado", id)
	}

	return nil
}

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASS_DB")
	host := os.Getenv("HOST_DB")
	port := os.Getenv("PORT_DB")
	table := os.Getenv("TABLE")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, table)
	log.Println(dataSource)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("could not ping the database: ", err)
		return nil, err
	}

	log.Println("connected")
	return db, nil
}
