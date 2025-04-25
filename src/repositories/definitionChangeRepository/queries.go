package definitionChangeRepository

var RegisterDefinitionChange = func() string {
	return `
INSERT INTO DefinitionChanges (Id, UserId, DefinitionId, OldValue, NewValue) VALUES (@Id, @UserId,
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

var GetDefinitionChangeState = func() string {
	return `
SELECT NewValue FROM DefinitionChanges 
WHERE DefinitionId = @DefinitionId
`
}
