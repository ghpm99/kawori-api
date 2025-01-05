package tag

import (
	"database/sql"
	"errors"
	"fmt"
	"kawori/api/pkg/database/queries"
	"kawori/api/pkg/utils"
)

type Repository struct {
	dbContext *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database}
}

func (repository *Repository) CreateTagRepository(transaction *sql.Tx, tag Tag) (Tag, error) {

	fail := func(err error) (Tag, error) {
		return tag, fmt.Errorf("CreateTagRepository: %v", err)
	}

	defer transaction.Rollback()

	result, err := transaction.Exec(
		queries.CreateTagQuery,
		&tag.Name,
		&tag.Color,
		&tag.UserId,
	)

	if err != nil {
		return fail(err)
	}

	tagId, err := result.LastInsertId()

	if err != nil {
		return fail(err)
	}

	tag.Id = int(tagId)

	return tag, nil
}

func (repository *Repository) UpdateTagRepository(transaction *sql.Tx, tag Tag) (bool, error) {

	fail := func(err error) (bool, error) {
		return false, fmt.Errorf("UpdateTagRepository: %v", err)
	}

	defer transaction.Rollback()

	result, err := transaction.Exec(
		queries.UpdateTagQuery,
		&tag.Id,
		&tag.UserId,
		&tag.Name,
		&tag.Color,
	)

	if err != nil {
		return fail(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fail(err)
	}

	if rowsAffected > 1 {
		transaction.Rollback()
		return fail(errors.New("multiple rows affected"))
	}

	return rowsAffected == 1, nil
}

func (repository *Repository) GetTagRepository(pagination utils.Pagination, filters TagFilter) (GetTagReturn, error) {

	fail := func(err error) (GetTagReturn, error) {
		return GetTagReturn{}, fmt.Errorf("GetTagRepository: %v", err)
	}

	data, err := repository.dbContext.Query(
		queries.GetAllTagsQuery,
		pagination.PageSize,
		pagination.Page,
		filters.UserId,
		filters.Name,
	)

	if err != nil {
		return fail(err)
	}
	var tagsArray []Tag
	for data.Next() {
		var tag Tag

		if err := data.Scan(
			&tag.Id,
			&tag.Name,
			&tag.Color,
			&tag.UserId,
		); err != nil {
			return fail(err)

		}
		tagsArray = append(tagsArray, tag)
	}
	if err := data.Err(); err != nil {
		return fail(err)
	}

	if pagination.Page > 1 {
		pagination.HasPrev = true
	}

	if tagsArray == nil {
		tagsArray = []Tag{}
	}

	return GetTagReturn{
		data:     tagsArray,
		pageInfo: pagination,
	}, nil
}
