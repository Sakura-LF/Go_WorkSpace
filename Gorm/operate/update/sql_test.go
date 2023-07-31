package update

import "testing"

func TestSql(t *testing.T) {
	Sql()
}

func TestSqlExc(t *testing.T) {
	SqlExc()
}

func TestRows(t *testing.T) {
	Rows()
}

func TestSessionIssue(t *testing.T) {
	SessionIssue()
}

func TestSessionTest(t *testing.T) {
	SessionTest()
}

func TestSessionOption(t *testing.T) {
	SessionOption()
}

func TestDryRun(t *testing.T) {
	DryRun()
}

func TestContextTOCancel(t *testing.T) {
	ContextTOCancel()
}
