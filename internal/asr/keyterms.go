package asr

// Keyterm lists for Deepgram Nova-3 keyterm prompting.
// Each environment gets a focused set of terms the model
// would not reliably recognize on its own. Common words
// like "fever", "surgery", "temperature" are omitted
// because Nova-3 already handles them well.
//
// Usage: pass the result of GlobalKeyterms() merged with
// the relevant environment slice to Transcribe().

// GlobalKeyterms returns terms that appear across multiple
// farm environments and should always be included.
func GlobalKeyterms() []string {
	return []string{
		// Drug names (brand and generic)
		"Draxxin",
		"tulathromycin",
		"Banamine",
		"flunixin",
		"Excede",
		"ceftiofur",
		"Penicillin",
		"Penicilina",
		"Dexamethasone",
		"Dexametasona",
		"Lutalyse",
		"prostaglandina",

		// Routes of administration
		"subcutaneous",
		"subQ",
		"intramammary",
		"intramamario",
		"subcutáneo",
		"intravenoso",
		"intramuscular",
		"calcio intravenoso",

		// Withdrawal and food safety
		"withdrawal period",
		"tiempo de retiro",
		"retiro de carne",
		"retiro de leche",
		"bulk tank",
		"tanque de leche",
		"dump bucket",
		"cubeta de descarte",

		// Common cross-environment veterinary Spanish
		"recién parida",
		"vaca tratada",
		"vaca caída",
		"vaca tirada",
	}
}

// MilkingParlorKeyterms covers mastitis detection,
// withdrawal compliance, and equipment concerns.
var MilkingParlorKeyterms = []string{
	"mastitis",
	"CMT test",
	"California Mastitis Test",
	"prueba CMT",
	"somatic cell count",
	"SCC",
	"conteo de células somáticas",
	"strip milk",
	"ordeñar a mano",
	"teat dip",
	"sellador de pezones",
	"clumpy milk",
	"grumos",
	"milking unit",
	"pezonera",
	"unidad de ordeño",
	"cuarto",
	"banda en la pata",
	"periodo seco",
	"ordeños",
}

// GeneralBarnKeyterms covers sick cow identification,
// routine herd health, and downer cow scenarios.
var GeneralBarnKeyterms = []string{
	"displaced abomasum",
	"desplazamiento de abomaso",
	"DA",
	"hypocalcemia",
	"hipocalcemia",
	"milk fever",
	"fiebre de leche",
	"downer cow",
	"ruminating",
	"cud chewing",
	"rumiando",
	"off feed",
	"sin apetito",
	"ping",
	"recumbent",
	"postrada",
	"echada",
	"splayed legs",
	"patas abiertas",
	"hip lifter",
	"levantador de cadera",
	"headlock",
}

// HoofTrimmingKeyterms covers lesion assessment,
// treatment, and lameness scoring.
var HoofTrimmingKeyterms = []string{
	"sole ulcer",
	"úlcera de suela",
	"hoof block",
	"wooden block",
	"taco",
	"outer claw",
	"inner claw",
	"peña de afuera",
	"peña de adentro",
	"digital dermatitis",
	"hairy heel wart",
	"dermatitis digital",
	"verruga",
	"white line disease",
	"enfermedad de línea blanca",
	"lameness",
	"cojera",
	"hoof trimming",
	"recorte de pezuñas",
	"absceso",
	"hinchazón",
}

// CalvingKeyterms covers dystocia, difficult births,
// and newborn calf care.
var CalvingKeyterms = []string{
	"dystocia",
	"distocia",
	"parto difícil",
	"OB chains",
	"calving chains",
	"cadenas de parto",
	"breech",
	"de nalgas",
	"colostrum",
	"calostro",
	"presentation",
	"presentación",
	"fetlock",
	"menudillo",
	"amniotic sac",
	"water bag",
	"bolsa de agua",
	"OB sleeve",
	"guante OB",
	"calf jack",
	"saca becerros",
	"jack de parto",
	"traction",
	"tracción",
}

// TreatmentPenKeyterms covers medication administration
// and post-surgical monitoring. Drug names are in the
// global set so this focuses on delivery terminology.
var TreatmentPenKeyterms = []string{
	"jugular vein",
	"vena yugular",
	"milliliters",
	"mililitros",
	"mL",
	"gauge",
	"calibre",
	"neumonía",
	"pneumonia",
}

// BreedingKeyterms covers heat detection, AI timing,
// pregnancy checks, and retained placenta.
var BreedingKeyterms = []string{
	"artificial insemination",
	"inseminación artificial",
	"standing heat",
	"estrus",
	"celo",
	"preg check",
	"pregnancy check",
	"checar preñez",
	"inseminada",
	"preñez confirmada",
	"days in gestation",
	"días de gestación",
	"corpus luteum",
	"cuerpo lúteo",
	"CL",
	"OvSynch",
	"synchronization protocol",
	"protocolo de sincronización",
	"GnRH injection",
	"inyección de GnRH",
	"prostaglandin",
	"retained placenta",
	"retained membranes",
	"placenta retenida",
	"metritis",
	"uterine discharge",
	"flujo uterino",
	"semen tank",
	"tanque de semen",
	"breeding window",
	"ventana de inseminación",
	"standing to be mounted",
	"dejarse montar",
}

// BiosecurityKeyterms covers isolation and
// contagious disease protocols.
var BiosecurityKeyterms = []string{
	"isolation pen",
	"corral de aislamiento",
	"biosecurity",
	"bioseguridad",
	"disinfection",
	"desinfección",
	"desinfectar",
	"nasal discharge",
	"moco nasal",
	"descarga nasal",
	"cuarentena",
	"cloro diluido",
}

// KeytermsForEnvironment returns the merged global +
// environment-specific keyterm list ready to pass
// to the Deepgram API.
func KeytermsForEnvironment(env string) []string {
	global := GlobalKeyterms()

	var local []string
	switch env {
	case "milking_parlor":
		local = MilkingParlorKeyterms
	case "general_barn":
		local = GeneralBarnKeyterms
	case "hoof_trimming":
		local = HoofTrimmingKeyterms
	case "calving":
		local = CalvingKeyterms
	case "treatment_pen":
		local = TreatmentPenKeyterms
	case "breeding":
		local = BreedingKeyterms
	case "biosecurity":
		local = BiosecurityKeyterms
	default:
		return global
	}

	// Deduplicate in case a local term overlaps with global
	seen := make(map[string]bool, len(global)+len(local))
	merged := make([]string, 0, len(global)+len(local))

	for _, t := range global {
		if !seen[t] {
			seen[t] = true
			merged = append(merged, t)
		}
	}
	for _, t := range local {
		if !seen[t] {
			seen[t] = true
			merged = append(merged, t)
		}
	}

	return merged
}
