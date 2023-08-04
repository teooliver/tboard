package data

func CreateMockSprint() {
	// Create User
	user := NewUser("john", "john@email.com")

	// Create Sprint
	sprint := NewSprint("sprint_test", user.id, 2)

	// Create Board
	board1 := NewBoard("board1", user.id)

	// Create Columns
	todoColumn := NewColumn("Todo", user.id)
	blockedColumn := NewColumn("Blocked", user.id)
	inDevelopmentColumn := NewColumn("In Development", user.id)
	reviewColumn := NewColumn("Review", user.id)
	doneColumn := NewColumn("Done", user.id)

	// Create Tasks
	NewTask("Task 01", "description for task 01", user.id).moveToColumn(todoColumn.id)
	NewTask("Task 02", "description for task 02", user.id).moveToColumn(todoColumn.id)
	NewTask("Task 03", "description for task 03", user.id).moveToColumn(todoColumn.id)
	NewTask("Task 04", "description for task 04", user.id).moveToColumn(todoColumn.id)
	NewTask("Task 05", "description for task 05", user.id).moveToColumn(todoColumn.id)

	// Add board to sprint
	sprint.addBoard(board1.id)

	// Add Columns to Board
	board1.addColumn(todoColumn.id)
	board1.addColumn(blockedColumn.id)
	board1.addColumn(inDevelopmentColumn.id)
	board1.addColumn(reviewColumn.id)
	board1.addColumn(doneColumn.id)

	// Add Tasks To Todo Column

}
