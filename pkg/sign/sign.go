package sign

import (
	"errors"
	"time"
)

type Document struct {
	ID         int
	Content    string
	Signatures []Signature
	CreatedAt  time.Time
}

type Signature struct {
	Signer    string
	Timestamp time.Time
}

func SignDocument(doc *Document, signer string) error {

	for _, sig := range doc.Signatures {
		if sig.Signer == signer {
			return errors.New("ce signataire a déjà signé ce document")
		}
	}

	signature := Signature{
		Signer:    signer,
		Timestamp: time.Now(),
	}

	doc.Signatures = append(doc.Signatures, signature)

	// stocker les modifs dans la base de donnée ici
	//verifier la mis a jour des documents ici

	return nil
}

func VerifySignature(signature string) bool {
	// a implementer
	return signatureIsValid(signature)
}

func ApplySignature(document *Document, signer string) error {
	//a implementer
	if document.IsSignedBy(signer) {
		return errors.New("ce signataire a déjà signé ce document")
	}

	document.AddSignature(signer)
	return nil
}

func GetDocumentByID(id int) (*Document, error) {
	// Mettez en œuvre la logique de récupération du document par son ID ici
	// Renvoyez le document trouvé ou une erreur si le document n'est pas trouvé

	// Exemple de logique de récupération de document
	document, err := retrieveDocumentFromDatabase(id)
	if err != nil {
		return nil, err
	}
	return document, nil
}

// Exemple de logique de vérification de signature simplifiée
func signatureIsValid(signature string) bool {
	// Implémentez votre logique de vérification ici
	// Renvoyez true si la signature est valide, sinon false
	// Par exemple, vérifiez une signature avec une clé publique ou un algorithme de hachage
	return true
}

// Exemple de méthode pour ajouter une signature au document
func (d *Document) AddSignature(signer string) {
	signature := Signature{
		Signer:    signer,
		Timestamp: time.Now(),
	}
	d.Signatures = append(d.Signatures, signature)
}

// Exemple de méthode pour vérifier si un signataire a déjà signé le document
func (d *Document) IsSignedBy(signer string) bool {
	for _, sig := range d.Signatures {
		if sig.Signer == signer {
			return true
		}
	}
	return false
}

// Exemple de méthode pour récupérer un document à partir de la base de données
func retrieveDocumentFromDatabase(id int) (*Document, error) {
	// Mettez en œuvre la logique de récupération du document depuis la base de données
	// Renvoyez le document trouvé ou une erreur si le document n'est pas trouvé

	// Exemple de récupération de document fictive (à remplacer par la logique réelle)
	if id == 1 {
		return &Document{
			ID:         1,
			Content:    "Contenu du document de test",
			Signatures: []Signature{},
			CreatedAt:  time.Now(),
		}, nil
	}

	return nil, errors.New("Document non trouvé")
}
