package definitionChangeRepository

var RegisterDefinitionChange = func() string {
	return `
INSERT INTO DefinitionChanges (Id, UserId, DefinitionID, OldValue, NewValue) VALUES (@Id, @UserId,
@DefinitionId, @OldValue, @NewValue)
`
}

var UpdateDefinitionChangeState = func() string {
	return `
UPDATE DefinitionChanges 
SET State = @State
WHERE DefinitionId = @DefinitionId
`
}
