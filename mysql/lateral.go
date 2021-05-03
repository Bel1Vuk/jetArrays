package mysql

import "github.com/go-jet/jet/v2/internal/jet"

func LATERAL(selectStmt SelectStatement) lateralImpl {
	return lateralImpl{
		selectStmt: selectStmt,
	}
}

type lateralImpl struct {
	selectStmt SelectStatement
}

func (l lateralImpl) AS(alias string) SelectTable {
	subQuery := &selectTableImpl{
		SelectTable: jet.NewLateral(l.selectStmt, alias),
	}

	subQuery.readableTableInterfaceImpl.parent = subQuery

	return subQuery
}
