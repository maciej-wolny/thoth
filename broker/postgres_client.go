package main

import (
	"database/sql"
	"fmt"
	"strings"
	"thoth/avro"
)

type postgresClient struct {
	dbConn *sql.DB
}

func NewPostgresClient(config *DatabaseConfig) (*postgresClient, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return &postgresClient{
		dbConn: db,
	}, nil
}

type TelemetryData struct {
	Timestamp int    `json:"timestamp"`
	SensorId  string `json:"sensorId"`
	Id        int    `json:"id"`
	Version   string `json:"version"`
	Objects   []struct {
		Label string `json:"label"`
		Score int    `json:"score"`
		Ymin  int    `json:"ymin"`
		Ymax  int    `json:"ymax"`
		Xmin  int    `json:"xmin"`
		Xmax  int    `json:"xmax"`
	} `json:"objects"`
}

func (pc postgresClient) BatchInsertTelemetryData(batch *avro.TelemetryDataBatch) error {

	valueStrings := make([]string, 0, batch.Size)
	valueArgs := make([]interface{}, 0, batch.Size*9)
	i := 0
	for _, post := range batch.TelemetryDataBatch {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*10+1, i*10+2, i*10+3, i*10+4, i*10+5, i*10+6, i*10+7, i*10+8, i*10+9, i*10+10))
		for _, classificationResults := range post.Objects {
			valueArgs = append(valueArgs, fmt.Sprint(post.Timestamp))
			valueArgs = append(valueArgs, fmt.Sprint(post.Id))
			valueArgs = append(valueArgs, fmt.Sprint(post.SensorId))
			valueArgs = append(valueArgs, fmt.Sprint(post.Version))
			//valueArgs = append(valueArgs, "true")
			valueArgs = append(valueArgs, fmt.Sprint(classificationResults.Label))
			valueArgs = append(valueArgs, fmt.Sprint(classificationResults.Score))
			valueArgs = append(valueArgs, fmt.Sprint(classificationResults.Ymin))
			valueArgs = append(valueArgs, fmt.Sprint(classificationResults.Xmin))
			valueArgs = append(valueArgs, fmt.Sprint(classificationResults.Ymax))
			valueArgs = append(valueArgs, fmt.Sprint(classificationResults.Xmax))
		}
		i++

	}
	stmt := fmt.Sprintf(
		"INSERT INTO \"telemetry\" ("+
			"timestamp, "+
			"id, "+
			"sensor_id, "+
			"version, "+
			"classification_result_label, "+
			"classification_result_score, "+
			"classification_result_ymin, "+
			"classification_result_xmin, "+
			"classification_result_ymax, "+
			"classification_result_xmax) "+
			"VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := pc.dbConn.Exec(stmt, valueArgs...)
	return err
}

func (pc postgresClient) Close() {
	_ = pc.dbConn.Close()
}
