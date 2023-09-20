package mysql

import (
	"context"
	"time"

	"github.com/ping-yee/Awesome-Walking-Skeleton/service/repository"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) repository.GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (p *GameRepository) GetGameById(ctx context.Context, id int) (*repository.Game, error) {
	game := new(repository.Game)

	result := p.db.Table("games").Select("id", "Name").First(game, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (p *GameRepository) CreateGame(ctx context.Context, game *repository.Game) (*repository.Game, error) {

	result := p.db.Table("games").Create(game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}
