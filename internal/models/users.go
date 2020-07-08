package models

//Row
type SQLUser struct {
    Id      uint
    Nome 	string
	Email	string
}

//ScanRow
func (u *SQLUser) ScanRow(r Row) error {
    return r.Scan(
        &u.Id,
        &u.Nome,
        &u.Email,

    )
}

//List
type SQLUserList struct {
    Res []*SQLUser
	Len uint
}

//List ScanRow
func (list *SQLUserList) ScanRow(r Row) error {
    a := new(SQLUser)
    if err := a.ScanRow(r); err != nil {
        return err
    }

    list.Res = append(list.Res, a)
	list.Len++

    return nil
}
