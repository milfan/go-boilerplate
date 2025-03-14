package entities

import (
	"time"

	"github.com/milfan/go-boilerplate/configs/constants"
)

type BaseEntity struct {
	createdBy string
	createdAt time.Time
	updatedBy string
	updatedAt time.Time
}

func New(createdBy string) BaseEntity {
	now := time.Now()
	return BaseEntity{
		createdAt: now,
		createdBy: createdBy,
		updatedAt: now,
		updatedBy: createdBy,
	}
}

func (b *BaseEntity) CreatedBy() string {
	return b.createdBy
}

func (b *BaseEntity) SetCreatedBy(val string) {
	b.updatedBy = val
}

func (b *BaseEntity) CreatedAt() time.Time {
	return b.createdAt
}

func (b *BaseEntity) SetCreatedAt(val time.Time) {
	b.createdAt = val
	b.updatedAt = val
}

func (m *BaseEntity) CreatedAtAsISOString() string {
	return m.createdAt.UTC().Format(constants.ISODateTimeFormat)
}

func (m *BaseEntity) CreatedAtAsSQLTimestampFormat() string {
	return m.createdAt.UTC().Format(constants.SQLTimestampFormat)
}

func (b *BaseEntity) UpdatedBy() string {
	return b.updatedBy
}

func (b *BaseEntity) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *BaseEntity) SetUpdatedAt() {
	b.updatedAt = time.Now()
}

func (b *BaseEntity) SetUpdatedBy(val string) {
	b.updatedBy = val
	b.SetUpdatedAt()
}

func (m *BaseEntity) UpdatedAtISOString() string {
	return m.updatedAt.UTC().Format(constants.ISODateTimeFormat)
}

func (m *BaseEntity) UpdatedAtSQLTimestampFormat() string {
	return m.updatedAt.UTC().Format(constants.SQLTimestampFormat)
}
