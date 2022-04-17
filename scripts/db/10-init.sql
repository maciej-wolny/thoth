CREATE USER helsing_edge_user WITH PASSWORD '123';
CREATE DATABASE helsing_edge OWNER helsing_edge_user;

CREATE TABLE IF NOT EXISTS telemetry
(
    timestamp                   INTEGER,
    sensor_id                   INTEGER,
    id                          BIGINT,
    version                     SMALLINT,
    classification_result_label VARCHAR(100),
    classification_result_score INTEGER,
    classification_result_ymin  INTEGER,
    classification_result_xmin  INTEGER,
    classification_result_ymax  INTEGER,
    classification_result_xmax  INTEGER
--     was_inserted  BOOL
);

INSERT INTO telemetry (timestamp, sensorId, id, version, classification_result_label, classification_result_score, classification_result_ymin, classification_result_xmin, classification_result_ymax, classification_result_xmax) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?),(?, ?, ?, ?, ?, ?, ?, ?, ?)