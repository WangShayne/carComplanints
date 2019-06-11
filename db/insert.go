package db

import "carComplanints/common"

func InsertComplanint(com *Complanint) int64 {
	sqlStr := "INSERT INTO complanints(sn,brandName,brandId,seriesName,seriesId,modelName,modelId,sketch,sketchURL,date) VALUE(?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.DB.Prepare(sqlStr)
	common.CheckErr(err)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	result, err := stmt.Exec(
		com.SN,
		com.Brand,
		com.BrandID,
		com.Series,
		com.SeriesID,
		com.Model,
		com.ModelID,
		com.Sketch,
		com.SketchURL,
		com.Date,
	)
	CheckErr(err)

	id, err := result.LastInsertId()

	CheckErr(err)

	return id

}
