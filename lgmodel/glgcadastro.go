package lgmodel

import "github.com/gofrs/uuid"

//
type LabCadastroEmail struct {
	LabCadastroEmailUUID uuid.UUID
	Email                string
}

//
type LabGobCadastro struct {
	LabCadastroUUID uuid.UUID
	Nome            string
	ListMail        []LabCadastroEmail
}
