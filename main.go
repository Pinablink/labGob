package main

import (
	"fmt"

	"github.com/Pinablink/labgob/lgmodel"
	"github.com/Pinablink/labgob/lgserialize"
	"github.com/Pinablink/lingue"
	"github.com/Pinablink/lingue/oslabel"
	"github.com/gofrs/uuid"
)

// Detecta qual SO em execução e chama o clear da tela de comando
var cmdClear *lingue.Lingue = lingue.NewLingue()

func main() {

	cmdClear.ExecCommand(oslabel.CLEAR_CMD)
	fmt.Println("Lab Gob")

	// Inicializando um de Mapa de Dados
	var mapCadastro map[string]lgmodel.LabGobCadastro = make(map[string]lgmodel.LabGobCadastro)

	// CRIAR ESTRUTURA DE DADOS
	// Cadastro Usuário 1
	emailTeste1uuid, _ := uuid.NewV4()
	var emailTeste1 lgmodel.LabCadastroEmail = lgmodel.LabCadastroEmail{
		LabCadastroEmailUUID: emailTeste1uuid,
		Email:                "emailTeste1@gmail.com",
	}

	cadastroTeste1uuid, _ := uuid.NewV4()
	sliceCadastroEmail1 := make([]lgmodel.LabCadastroEmail, 0)
	sliceCadastroEmail1 = append(sliceCadastroEmail1, emailTeste1)
	var cadastroTeste1 lgmodel.LabGobCadastro = lgmodel.LabGobCadastro{
		LabCadastroUUID: cadastroTeste1uuid,
		Nome:            "Teste de usuario 1",
		ListMail:        sliceCadastroEmail1,
	}

	// Cadastro Usuário 2
	emailTeste2uuid, _ := uuid.NewV4()
	var emailTeste2 lgmodel.LabCadastroEmail = lgmodel.LabCadastroEmail{
		LabCadastroEmailUUID: emailTeste2uuid,
		Email:                "emailTeste2@gmail.com",
	}

	cadastroTeste2uuid, _ := uuid.NewV4()
	sliceCadastroEmail2 := make([]lgmodel.LabCadastroEmail, 0)
	sliceCadastroEmail2 = append(sliceCadastroEmail2, emailTeste2)
	var cadastroTeste2 lgmodel.LabGobCadastro = lgmodel.LabGobCadastro{
		LabCadastroUUID: cadastroTeste2uuid,
		Nome:            "Teste de usuario 2",
		ListMail:        sliceCadastroEmail2,
	}

	// Adicionando estrutura ao mapa
	mapCadastro["emailTeste1@gmail.com"] = cadastroTeste1
	mapCadastro["emailTeste2@gmail.com"] = cadastroTeste2

	fmt.Printf("\n\n\nMapa Criado\n%s\n\n\n", mapCadastro)

	// Preparando para serializar
	var obSerialize *lgserialize.LabGobSerialize = lgserialize.NewLabGobSerialize(mapCadastro, "C:\\lab_gob\\testeLabGob.mgob")
	// Serializando o map
	obSerialize.Serialize()

	// Deserializando o map
	/*mapCadastroDeserialized, errDes := obSerialize.Deserialize()

	if errDes != nil {
		fmt.Println(errDes.Error())
	} else {
		fmt.Println("Mapa descomprimido")
		fmt.Println(mapCadastroDeserialized)
	}
	*/
}
