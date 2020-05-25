package decorator_pattern

type DataSource interface {
	WriteData(data string)
	ReadData() (data string)
}

type FileDataSource struct {
}

func (f *FileDataSource) WriteData(data string) {
	// 将数据写入文件
}

func (f *FileDataSource) ReadData() (data string) {
	// 从文件读取数据
	return data
}

type DataSourceDecorator struct {
	wrappee DataSource
}

func (d *DataSourceDecorator) WriteData(data string) {
	d.wrappee.WriteData(data)
}

func (d *DataSourceDecorator) ReadData() (data string) {
	return d.wrappee.ReadData()
}

// 之后还可以对DataSourceDecorator进行进一步的封装
