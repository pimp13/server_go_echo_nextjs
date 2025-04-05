package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Name      string    `bun:"name,nullzero,notnull"`
	Email     string    `bun:"email,nullzero,notnull,unique"`
	Password  string    `bun:"password,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
