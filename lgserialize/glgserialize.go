package lgserialize

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"github.com/Pinablink/labgob/lgmodel"
	"github.com/gofrs/uuid"
)

// Estrutura especializada em serializar, comprimir e salvar
// um Map em disco local
type LabGobSerialize struct {
	LabGobSerializeUUID uuid.UUID
	PathStream          string
	MMap                map[string]lgmodel.LabGobCadastro
}

//
func NewLabGobSerialize(refMap map[string]lgmodel.LabGobCadastro, strPath string) *LabGobSerialize {
	iuuid, _ := uuid.NewV4()

	return &LabGobSerialize{
		LabGobSerializeUUID: iuuid,
		PathStream:          strPath,
		MMap:                refMap,
	}
}

//
func (refLabGobSerialize *LabGobSerialize) Serialize() bool {

	var mBuffer bytes.Buffer
	var encoder *gob.Encoder = gob.NewEncoder(&mBuffer)

	err := encoder.Encode(refLabGobSerialize.MMap)

	if err != nil {
		fmt.Println("Houve um erro no processo de serialização")
		fmt.Println(err.Error())
		return false
	}

	fmt.Printf("\n\n\nMapa Serializado\n\nTamanho: %d Bytes\n%x\n\n", len(mBuffer.Bytes()), mBuffer.Bytes())

	compressedData, errCompress := compressBytes(mBuffer.Bytes())

	if errCompress != nil {
		fmt.Println("Houve erro no processo de compactação dos dados serializados")
		fmt.Println(errCompress.Error())
		return false
	}

	errSave := saveCompressBytes(refLabGobSerialize.PathStream, compressedData)

	if errSave != nil {
		fmt.Println("Houve erro no processo de persistência dos dados serializados")
		fmt.Println(errSave.Error())
		return false
	}

	fmt.Printf("\n\n\n\nRecurso salvo com sucesso\n")

	return true
}

//

//
func (refLabGobSerialize *LabGobSerialize) Deserialize() (map[string]lgmodel.LabGobCadastro, error) {

	var strPathStream string = refLabGobSerialize.PathStream

	dataByte, errReadFile := os.ReadFile(strPathStream)

	if errReadFile != nil {
		fmt.Println("Ocorreu um erro na obtenção dos dados serializados")
		return nil, errReadFile
	}

	var bufferReader *bytes.Buffer = bytes.NewBuffer(dataByte)
	gzipReader, errGzipReader := gzip.NewReader(bufferReader)

	if errGzipReader != nil {
		fmt.Println("Ocorreu um erro na descompressão dos dados serializados")
		return nil, errGzipReader
	}

	var bytesBufferWrite bytes.Buffer
	_, errCopy := io.Copy(&bytesBufferWrite, gzipReader)

	if errCopy != nil {
		fmt.Println("Ocorreu um erro na leitura dos dados serializados")
		return nil, errCopy
	}

	// Ponto de deserialização do objeto
	var mapRet map[string]lgmodel.LabGobCadastro = make(map[string]lgmodel.LabGobCadastro)
	var bufferDeserialize bytes.Buffer
	bufferDeserialize.Write(bytesBufferWrite.Bytes())
	var decoder *gob.Decoder = gob.NewDecoder(&bufferDeserialize)

	errDecode := decoder.Decode(&mapRet)

	if errDecode != nil {
		fmt.Println("Ocorreu um erro no decode da struct de dados")
		return nil, errDecode
	}

	return mapRet, nil
}

//
func compressBytes(streamBytes []byte) (compressedData []byte, errCompress error) {

	var bytesBufferCompress bytes.Buffer
	writeCompress, _ := gzip.NewWriterLevel(&bytesBufferCompress, gzip.BestCompression)
	_, err := writeCompress.Write(streamBytes)

	writeCompress.Close()

	if err != nil {
		fmt.Println("Ocorreu erro no processamento de compressão dos dados.")
		return nil, err
	}

	fmt.Printf("\n\n\nMapa Serializado Comprimido\n\nTamanho: %d Bytes\n%x\n\n", len(bytesBufferCompress.Bytes()), bytesBufferCompress.Bytes())

	return bytesBufferCompress.Bytes(), nil
}

//
func saveCompressBytes(strPathStream string, bytesCompressed []byte) error {

	file, err := os.Create(strPathStream)

	if err != nil {
		return err
	}

	_, errFW := file.Write(bytesCompressed)

	if errFW != nil {
		return err
	}

	file.Close()

	return nil
}
