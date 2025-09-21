package events

import "context"

func (r *Repository) DeleteEvent(ctx context.Context, id int) error {
	_, err := r.Exec(ctx, `DELETE FROM events WHERE id = $1`, id)
	return err
}
