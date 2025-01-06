package utils

//
//// Read ExcelData
//func Read(filename string) (*model.ExcelData, error) {
//	excel, err := excelize.OpenFile(filename)
//	if err != nil {
//		return nil, err
//	}
//	defer func() {
//		err := excel.Close()
//		if err != nil {
//			zap.S().Error(err)
//			return
//		}
//	}()
//	rows, err := excel.GetRows("Sheet1")
//	if err != nil {
//		return nil, err
//	}
//	if len(rows) < 2 {
//		return nil, errors.New(errors.DataNotFound)
//	}
//	titles := rows[0]
//	titleMap := make(map[int]string, len(titles))
//	for i, title := range titles {
//		titleMap[i] = title
//	}
//
//	dataList := make([]map[int]string, len(rows)-1)
//	for i := 1; i < len(rows); i++ {
//		cells := rows[i]
//
//		cellMap := make(map[int]string, len(titles))
//		for j, cell := range cells {
//			cellMap[j] = cell
//		}
//		dataList[i] = cellMap
//	}
//
//	excelData := model.ExcelData{
//		Titles: titleMap,
//		Datas:  dataList,
//	}
//	return &excelData, nil
//}
