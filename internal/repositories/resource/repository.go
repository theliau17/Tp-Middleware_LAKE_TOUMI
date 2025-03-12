package resource

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"middleware/example/internal/models"
)

type Repository interface {
	GetAll() ([]models.Resource, error)
	GetByID(id string) (*models.Resource, error)
	Create(r *models.Resource) error
	Update(r *models.Resource) error
	Delete(id string) error
}

type repository struct {
	db *sql.DB
}

// NewRepository retourne une instance du repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]models.Resource, error) {
	rows, err := r.db.Query("SELECT id, name, file_path, created_at FROM resources")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []models.Resource
	for rows.Next() {
		var res models.Resource
		var createdAtStr string
		if err := rows.Scan(&res.ID, &res.Name, &res.FilePath, &createdAtStr); err != nil {
			return nil, err
		}
		res.CreatedAt, _ = time.Parse(time.RFC3339, createdAtStr)
		resources = append(resources, res)
	}
	return resources, nil
}

func (r *repository) GetByID(id string) (*models.Resource, error) {
	row := r.db.QueryRow("SELECT id, name, file_path, created_at FROM resources WHERE id = ?", id)
	var res models.Resource
	var createdAtStr string
	if err := row.Scan(&res.ID, &res.Name, &res.FilePath, &createdAtStr); err != nil {
		return nil, err
	}
	res.CreatedAt, _ = time.Parse(time.RFC3339, createdAtStr)
	return &res, nil
}

func (r *repository) Create(res *models.Resource) error {
	res.ID = uuid.New().String()
	res.CreatedAt = time.Now().UTC()
	_, err := r.db.Exec(`INSERT INTO resources (id, name, file_path, created_at) VALUES (?, ?, ?, ?)`,
		res.ID, res.Name, res.FilePath, res.CreatedAt.Format(time.RFC3339))
	return err
}

func (r *repository) Update(res *models.Resource) error {
	_, err := r.db.Exec(`UPDATE resources SET name = ?, file_path = ? WHERE id = ?`,
		res.Name, res.FilePath, res.ID)
	return err
}

func (r *repository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM resources WHERE id = ?`, id)
	return err
}
