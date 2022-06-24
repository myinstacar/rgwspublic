package rgwspublic

import (
	"encoding/json"
	"encoding/xml"
)

// XMLResponse is where we parse an http response
type XMLResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    XMLBody
}

// XMLBody is the body of a response
type XMLBody struct {
	XMLName           xml.Name `xml:"Body"`
	AFMMethodResponse AFMMethodResponse
}

// XMLRGWSPublicAfmMethodResponse is the response to a publicAFM method
type AFMMethodResponse struct {
	XMLName        xml.Name `xml:"rgWsPublic2AfmMethodResponse" json:"-"`
	Result         ResultData
	AFMCalledByRec AFMCalledByRecData
	BasicRec       BasicRecData
}

// XMLRGWSPublicVersionInfoResponse holds version info
type XMLRGWSPublicVersionInfoResponse struct {
	XMLName xml.Name `xml:"rgWsPublicVersionInfoResponse" json:"-"`
	Result  string   `xml:"result"`
}

type ResultData struct {
	XMLName    xml.Name `xml:"result" json:"-"`
	ResultType ResultTypeData
}

type ResultTypeData struct {
	XMLName            xml.Name `xml:"rg_ws_public2_result_rtType" json:"-"`
	CallSeqID          CallSeqIDData
	ErrorRec           ErrorRecData
	AFMCalledByRecData AFMCalledByRecData
}

type CallSeqIDData struct {
	CallSeqID string `xml:"call_seq_id" json:"call_seq_id"`
}

type ErrorRecData struct {
	XMLName          xml.Name `xml:"error_rec" json:"-"`
	ErrorCode        string   `xml:"error_code" json:"error_code"`
	ErrorDescription string   `xml:"error_description" json:"error_description"`
}

// AFMDCalledByData is the data relative to who did the search
type AFMCalledByRecData struct {
	XMLName             xml.Name `xml:"afm_called_by_rec" json:"-"`
	TokenUsername       string   `xml:"token_username" json:"token_username"`
	TokenAFM            string   `xml:"token_afm" json:"token_afm"`
	TokenAFMFullName    string   `xml:"token_afm_fullname" json:"token_afm_fullname"`
	AFMCalledBy         string   `xml:"afm_called_by" json:"afm_called_by"`
	AFMCalledByFullName string   `xml:"afm_called_by_fullname" json:"afm_called_by_fullname"`
	Activities          []FirmActivities
}

// BasicRecData is the data relative to an entity's VAT search
type BasicRecData struct {
	XMLName                     xml.Name `xml:"basic_rec" json:"-"`
	AFM                         string   `xml:"afm" json:"afm"`                                              // ΑΦΜ
	DOY                         string   `xml:"doy" json:"doy"`                                              // ΚΩΔΙΚΟΣ ΔΟΥ
	DOYDescription              string   `xml:"doy_descr" json:"doy_description"`                            // ΠΕΡΙΓΡΑΦΗ ΔΟΥ
	InitialFlagDescription      string   `xml:"i_ni_flag_descr" json:"initial_flag_description"`             // ΦΠ /ΜΗ ΦΠ
	DeactivationFlag            string   `xml:"deactivation_flag" json:"deactivation_flag"`                  // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ:1=ΕΝΕΡΓΟΣ ΑΦΜ 2=ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	DeactivationFlagDescription string   `xml:"deactivation_flag_desc" json:"deactivation_flag_description"` // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ(ΠΕΡΙΓΡΑΦΗ): ΕΝΕΡΓΟΣ ΑΦΜ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	FirmFlagDescription         string   `xml:"firm_flag_descr" json:"firm_flag_description"`                // ΤΙΜΕΣ: ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΜΗ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΠΡΩΗΝ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ
	Onomasia                    string   `xml:"onomasia" json:"onomasia"`                                    // ΕΠΩΝΥΜΙΑ
	CommercialTitle             string   `xml:"commer_title" json:"commercial_title"`                        // ΤΙΤΛΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	LegalStatusDescription      string   `xml:"legal_status_descr" json:"legal_status_descr"`                // ΠΕΡΙΓΡΑΦΗ ΜΟΡΦΗΣ ΜΗ Φ.Π.
	PostalAddress               string   `xml:"postal_address" json:"postal_address"`                        // ΟΔΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAddressNo             string   `xml:"postal_address_no" json:"postal_address_no"`                  // ΑΡΙΘΜΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalZipCode               string   `xml:"postal_zip_code" json:"postal_zip_code"`                      // ΤΑΧ. ΚΩΔ. ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAreaDescription       string   `xml:"postal_area_description" json:"postal_area_description"`      // ΠΕΡΙΟΧΗ ΕΠΙΧΕΙΡΗΣΗΣ
	RegistrationDate            string   `xml:"regist_date" json:"registration_date"`                        // ΗΜ/ΝΙΑ ΕΝΑΡΞΗΣ
	StopDate                    string   `xml:"stop_date" json:"stop_date"`                                  // ΗΜ/ΝΙΑ ΔΙΑΚΟΠΗΣ
	NormalVATSystemFlag         string   `xml:"normal_vat_system_flag" json:"normal_vat_system_flag"`
	Activities                  FirmActivities
}

// FirmActivities is the activities of the entity
type FirmActivities struct {
	XMLName xml.Name `xml:"firm_act_tab" json:"-"`
	Items   []Item
}

type Item struct {
	XMLName             xml.Name `xml:"item" json:"-"`
	FirmActCode         string   `xml:"firm_act_code" json:"firm_act_code"`                   // ΚΩΔΙΚΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActDescriptionn string   `xml:"firm_act_descr" json:"firm_act_description"`           // ΠΕΡΙΓΡΑΦΗ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	FirmActKind         string   `xml:"firm_act_kind" json:"firm_activity_kind"`              // ΕΙΔΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: 1=ΚΥΡΙΑ, 2=ΔΕΥΤΕΡΕΥΟΥΣΑ, 3=ΛΟΙΠΗ, 4=ΒΟΗΘΗΤΙΚΗ
	FirmActKindDescr    string   `xml:"firm_act_kind_descr" json:"firm_act_kind_description"` // ΠΕΡΙΓΡΑΦΗ ΕΙΔΟΥΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: ΚΥΡΙΑ, ΔΕΥΤΕΡΕΥΟΥΣΑ, ΛΟΙΠΗ, ΒΟΗΘΗΤΙΚΗ
}

const (
	// Endpoint is the url for WSDL service
	Endpoint                                       = "https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2?WSDL"
	RG_WS_PUBLIC_AFM_CALLED_BY_BLOCKED             = "Ο χρήστης που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."
	RG_WS_PUBLIC_AFM_CALLED_BY_NOT_FOUND           = "Ο Α.Φ.Μ. για τον οποίο γίνεται η κλήση δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_EPIT_NF                           = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν ανήκει και δεν ανήκε ποτέ σε νομικό πρόσωπο, νομική οντότητα, ή φυσικό πρόσωπο με εισόδημα από επιχειρηματική δραστηριότητα."
	RG_WS_PUBLIC_FAILURES_TOLERATED_EXCEEDED       = "Υπέρβαση μέγιστου επιτρεπτού ορίου πρόσφατων αποτυχημένων κλήσεων. Προσπαθήστε εκ νέου σε μερικές ώρες."
	RG_WS_PUBLIC_MAX_DAILY_USERNAME_CALLS_EXCEEDED = "Υπέρβαση μέγιστου επιτρεπτού ορίου ημερήσιων κλήσεων ανά χρήστη (ανεξαρτήτως εξουσιοδοτήσεων)."
	RG_WS_PUBLIC_MONTHLY_LIMIT_EXCEEDED            = "Υπέρβαση του Μέγιστου Επιτρεπτού Μηνιαίου Ορίου Κλήσεων."
	RG_WS_PUBLIC_MSG_TO_TAXISNET_ERROR             = "Δημιουργήθηκε πρόβλημα κατά την ενημέρωση των εισερχόμενων μηνυμάτων στο MyTAXISnet."
	RG_WS_PUBLIC_NO_INPUT_PARAMETERS               = "Δε δόθηκαν υποχρεωτικές παράμετροι εισόδου για την κλήση της υπηρεσίας."
	RG_WS_PUBLIC_SERVICE_NOT_ACTIVE                = "Η υπηρεσία δεν είναι ενεργή."
	RG_WS_PUBLIC_TAXPAYER_NF                       = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_TOKEN_AFM_BLOCKED                 = "Ο χρήστης (ή ο εξουσιοδοτημένος τρίτος) που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_AUTHORIZED          = "Ο τρέχον χρήστης δεν έχει εξουσιοδοτηθεί από τον Α.Φ.Μ. για χρήση της υπηρεσίας."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_FOUND               = "Ο Α.Φ.Μ. του τρέχοντος χρήστη δε βρέθηκε στους έγκυρους Α.Φ.Μ. του Μητρώου TAXIS."
	RG_WS_PUBLIC_TOKEN_AFM_NOT_REGISTERED          = "Ο τρέχον χρήστης δεν έχει εγγραφεί για χρήση της υπηρεσίας."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_ACTIVE         = "Ο κωδικός χρήστη (username) που χρησιμοποιήθηκε έχει ανακληθεί."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED  = "Ο συνδυασμός χρήστη/κωδικού πρόσβασης που δόθηκε δεν είναι έγκυρος."
	RG_WS_PUBLIC_TOKEN_USERNAME_NOT_DEFINED        = "Δεν ορίσθηκε ο χρήστης που καλεί την υπηρεσία."
	RG_WS_PUBLIC_TOKEN_USERNAME_TOO_LONG           = "Διαπιστώθηκε υπέρβαση του μήκους του ονόματος του χρήστη (username) της υπηρεσίας"
	RG_WS_PUBLIC_WRONG_AFM                         = "O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν είναι έγκυρος."
)

func (a *ResultTypeData) JSON() (string, error) {

	var j []byte
	j, err := json.MarshalIndent(&a, "", "\t")
	if err != nil {
		return "", err
	}

	return string(j), nil
}

func (a *ResultTypeData) String() string {

	var s string

	// s = fmt.Sprintf("XMLName:%s\n", a.XMLName)
	// s += fmt.Sprintf("AFM:%s\n", a.AFM)
	// s += fmt.Sprintf("DOY:%s\n", a.DOY)
	// s += fmt.Sprintf("DOYDesc:%s\n", a.DOYDesc)
	// s += fmt.Sprintf("INiFlagDescr:%s\n", a.INiFlagDescr)
	// s += fmt.Sprintf("DeactivationFlag:%s\n", a.DeactivationFlag)
	// s += fmt.Sprintf("DeactivationFlagDescr:%s\n", a.DeactivationFlagDescr)
	// s += fmt.Sprintf("FirmFlagDescr:%s\n", a.FirmFlagDescr)
	// s += fmt.Sprintf("Onomasia:%s\n", a.Onomasia)
	// s += fmt.Sprintf("CommerTitle:%s\n", a.CommerTitle)
	// s += fmt.Sprintf("LegalStatusDescr:%s\n", a.LegalStatusDescr)
	// s += fmt.Sprintf("PostalAddress:%s\n", a.PostalAddress)
	// s += fmt.Sprintf("PostalAddressNo:%s\n", a.PostalAddressNo)
	// s += fmt.Sprintf("PostalZipCode:%s\n", a.PostalZipCode)
	// s += fmt.Sprintf("PostalAreaDescription:%s\n", a.PostalAreaDescription)
	// s += fmt.Sprintf("RegistDate:%s\n", a.RegistDate)
	// s += fmt.Sprintf("StopDate:%s\n", a.StopDate)

	// s += fmt.Sprintf("ACTIVITIES:--------------------\n")
	// for k, v := range a.Activities {
	// 	s += fmt.Sprintf("ACTIVITY #%d\n", k)
	// 	s += fmt.Sprintf("FirmActCode: %s\n", v.FirmActCode)
	// 	s += fmt.Sprintf("FirmActDescr: %s\n", v.FirmActDescr)
	// 	s += fmt.Sprintf("FirmActKind: %s\n", v.FirmActKind)
	// 	s += fmt.Sprintf("FirmActKindDescr: %s\n", v.FirmActKindDescr)
	// }

	// s += fmt.Sprintf("ErrorDescr: %s\n", a.Error.ErrorDescr)
	// s += fmt.Sprintf("ErrorCode: %s\n", a.Error.ErrorCode)

	return s

}
