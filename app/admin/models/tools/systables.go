package tools

import (
	"strings"

	"fiy/app/admin/models"
	orm "fiy/common/global"
)

type SysTables struct {
	TableId             int          `gorm:"primary_key;auto_increment" json:"tableId"`    //表编码
	TBName              string       `gorm:"column:table_name;size:255;" json:"tableName"` //表名称
	TableComment        string       `gorm:"size:255;" json:"tableComment"`                //表备注
	ClassName           string       `gorm:"size:255;" json:"className"`                   //类名
	TplCategory         string       `gorm:"size:255;" json:"tplCategory"`                 //
	PackageName         string       `gorm:"size:255;" json:"packageName"`                 //包名
	ModuleName          string       `gorm:"size:255;" json:"moduleName"`                  //模块名
	BusinessName        string       `gorm:"size:255;" json:"businessName"`                //
	FunctionName        string       `gorm:"size:255;" json:"functionName"`                //功能名称
	FunctionAuthor      string       `gorm:"size:255;" json:"functionAuthor"`              //功能作者
	PkColumn            string       `gorm:"size:255;" json:"pkColumn"`
	PkGoField           string       `gorm:"size:255;" json:"pkGoField"`
	PkJsonField         string       `gorm:"size:255;" json:"pkJsonField"`
	Options             string       `gorm:"size:255;" json:"options"`
	TreeCode            string       `gorm:"size:255;" json:"treeCode"`
	TreeParentCode      string       `gorm:"size:255;" json:"treeParentCode"`
	TreeName            string       `gorm:"size:255;" json:"treeName"`
	Tree                bool         `gorm:"size:1;" json:"tree"`
	Crud                bool         `gorm:"size:1;" json:"crud"`
	Remark              string       `gorm:"size:255;" json:"remark"`
	IsDataScope         int          `gorm:"size:1;" json:"isDataScope"`
	IsActions           int          `gorm:"size:1;" json:"isActions"`
	IsAuth              int          `gorm:"size:1;" json:"isAuth"`
	IsLogicalDelete     string       `gorm:"size:1;" json:"isLogicalDelete"`
	LogicalDelete       bool         `gorm:"size:1;" json:"logicalDelete"`
	LogicalDeleteColumn string       `gorm:"size:128;" json:"logicalDeleteColumn"`
	CreateBy            string       `gorm:"size:128;" json:"createBy"`
	UpdateBy            string       `gorm:"size:128;" json:"updateBy"`
	DataScope           string       `gorm:"-" json:"dataScope"`
	Params              Params       `gorm:"-" json:"params"`
	Columns             []SysColumns `gorm:"-" json:"columns"`

	models.BaseModel
}

func (SysTables) TableName() string {
	return "sys_tables"
}

type Params struct {
	TreeCode       string `gorm:"-" json:"treeCode"`
	TreeParentCode string `gorm:"-" json:"treeParentCode"`
	TreeName       string `gorm:"-" json:"treeName"`
}

func (e *SysTables) GetPage(pageSize int, pageIndex int) ([]SysTables, int, error) {
	var doc []SysTables

	table := orm.Eloquent.Table("sys_tables")

	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}

	var count int64

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	//table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, int(count), nil
}

func (e *SysTables) Get() (SysTables, error) {
	var doc SysTables
	var err error
	table := orm.Eloquent.Table("sys_tables")

	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableId != 0 {
		table = table.Where("table_id = ?", e.TableId)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	var col SysColumns
	col.TableId = doc.TableId
	if doc.Columns, err = col.GetList(); err != nil {
		return doc, err
	}

	return doc, nil
}

func (e *SysTables) GetTree() ([]SysTables, error) {
	var doc []SysTables
	var err error
	table := orm.Eloquent.Table("sys_tables")

	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableId != 0 {
		table = table.Where("table_id = ?", e.TableId)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	for i := 0; i < len(doc); i++ {
		var col SysColumns
		//col.FkCol = append(col.FkCol, SysColumns{ColumnId: 0, ColumnName: "请选择"})
		col.TableId = doc[i].TableId
		if doc[i].Columns, err = col.GetList(); err != nil {
			return doc, err
		}

	}

	return doc, nil
}

func (e *SysTables) Create() (SysTables, error) {
	var doc SysTables
	result := orm.Eloquent.Table("sys_tables").Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	for i := 0; i < len(e.Columns); i++ {
		e.Columns[i].TableId = doc.TableId

		e.Columns[i].Create()
	}

	return doc, nil
}

func (e *SysTables) Update() (update SysTables, err error) {
	//if err = orm.Eloquent.Table("sys_tables").First(&update, e.TableId).Error; err != nil {
	//	return
	//}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("sys_tables").Where("table_id = ?", e.TableId).Updates(&e).Error; err != nil {
		return
	}

	tableNames := make([]string, 0)
	for i := range e.Columns {
		if e.Columns[i].FkTableName != "" {
			tableNames = append(tableNames, e.Columns[i].FkTableName)
		}
	}

	tables := make([]SysTables, 0)
	tableMap := make(map[string]*SysTables)
	if len(tableNames) > 0 {
		if err = orm.Eloquent.Table("sys_tables").Where("table_name in (?)", tableNames).Find(&tables).Error; err != nil {
			return
		}
		for i := range tables {
			tableMap[tables[i].TBName] = &tables[i]
		}
	}

	for i := 0; i < len(e.Columns); i++ {
		if e.Columns[i].FkTableName != "" {
			t, ok := tableMap[e.Columns[i].FkTableName]
			if ok {
				e.Columns[i].FkTableNameClass = t.ClassName
				e.Columns[i].FkTableNamePackage = t.ModuleName
			} else {
				tableNameList := strings.Split(e.Columns[i].FkTableName, "_")
				e.Columns[i].FkTableNameClass = ""
				e.Columns[i].FkTableNamePackage = ""
				for a := 0; a < len(tableNameList); a++ {
					strStart := string([]byte(tableNameList[a])[:1])
					strEnd := string([]byte(tableNameList[a])[1:])
					e.Columns[i].FkTableNameClass += strings.ToUpper(strStart) + strEnd
					e.Columns[i].FkTableNamePackage += strings.ToLower(strStart) + strings.ToLower(strEnd)
				}
			}
		}
		_, _ = e.Columns[i].Update()
	}
	return
}

func (e *SysTables) Delete() (success bool, err error) {
	tx := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Table("sys_tables").Delete(SysTables{}, "table_id = ?", e.TableId).Error; err != nil {
		success = false
		return
	}
	if err = tx.Table("sys_columns").Delete(SysColumns{}, "table_id = ?", e.TableId).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (e *SysTables) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Unscoped().Table(e.TableName()).Where(" table_id in (?)", id).Delete(&SysColumns{}).Error; err != nil {
		return
	}
	Result = true
	return
}
