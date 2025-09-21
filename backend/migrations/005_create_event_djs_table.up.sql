CREATE TABLE event_djs (
    event_id INT REFERENCES events(id) ON DELETE CASCADE,
    dj_id INT REFERENCES djs(id) ON DELETE CASCADE,
    order_in_lineup INT,
    PRIMARY KEY (event_id, dj_id)
);

CREATE INDEX idx_event_djs_event_id_order
    ON event_djs(event_id, order_in_lineup);
