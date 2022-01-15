package enigma

import (
	"strconv"
	"strings"
	"testing"
)

type enigmaSpec struct {
	model           Model
	rotorConfig     string // I IV VII | A U C | 1 14 3
	reflectorConfig string // B | 15 | AA BB CC DD EE ...
	plugboardConfig string // AA BB CC DD EE FF GG HH II JJ
}

func TestEnigma_Encode(t *testing.T) {
	tests := []struct {
		name string
		spec enigmaSpec
		text string
		want string
	}{
		{
			name: "basic M3 with double-stepping",
			spec: enigmaSpec{
				model:           M3,
				rotorConfig:     "I II III | | ",
				reflectorConfig: "B | | ",
				plugboardConfig: "",
			},
			text: strings.Repeat("A", 200),
			want: "BDZGOWCXLTKSBTMCDLPBMUQOFXYHCXTGYJFLINHNXSHIUNTHEORXPQPKOVHCBUBTZSZSOOSTGOTFSODBBZZLXLCYZXIFGWFDZEEQIBMGFJBWZFCKPFMGBXQCIVIBBRNCOCJUVYDKMVJPFMDRMTGLWFOZLXGJEYYQPVPBWNCKVKLZTCBDLDCTSNRCOOVPTGBVBBISGJSO",
		},
		{
			name: "M3 with full settings",
			spec: enigmaSpec{
				model:           M3,
				rotorConfig:     "III VII VIII | D U S | 12 8 6",
				reflectorConfig: "C | | ",
				plugboardConfig: "AI BX CU DF EN GQ HM JL KT OP",
			},
			text: "LOREMQQIPSUMQQDOLORQQSITQQAMETQQCONSECTETUERQQADIPISCINGQQELITQQAENEANQQVELQQMASSAQQQUISQQMAURISQQVEHICULAQQLACINIAQQCURABITURQQSAGITTISQQHENDRERITQQANTEQQNAMQQQUISQQNULLAQQETIAMQQQUISQQQUAMQQALIQUAMQQINQQLOREMQQSITQQAMETQQLEOQQACCUMSANQQLACINIA",
			want: "JRKGDLRCOURDHDKHEEOOWVEJVPOOKOBHFFQXDNWDYEDDTKDWLRGLSJMBRRQYHQRUPUBVYHTIABJNKZYPRQVJXTXOZWOSIMQHDYWHUHKZGCVXDIYURDQGOIHNFMMDYMXDPFKXZQTXZMZGYOYBQKIFXPFSXHYBOWRSYQWLXHMIIEHUWPOJJSSBNOSCPELDEENEGTMXWZQRTXCKRQLGFQKBUOBKEBXVGWFYIRSFHBPRAWKIBEPLBMCEW",
		},
		{
			name: "M4 (4-rotors)",
			spec: enigmaSpec{
				model:           M4,
				rotorConfig:     "gamma VI I VII | L X A Q | 18 16 23 2",
				reflectorConfig: "BThin | | ",
				plugboardConfig: "",
			},
			text: "LOREMQQIPSUMQQDOLORQQSITQQAMETQQCONSECTETUERQQADIPISCINGQQELITQQAENEANQQVELQQMASSAQQQUISQQMAURISQQVEHICULAQQLACINIAQQCURABITURQQSAGITTISQQHENDRERITQQANTEQQNAMQQQUISQQNULLAQQETIAMQQQUISQQQUAMQQALIQUAMQQINQQLOREMQQSITQQAMETQQLEOQQACCUMSANQQLACINIA",
			want: "HZAFDYHADNXFLGTKODHHUCMCKFKFLOSTSMRPZNBLIZSYXGGTEGUHNQQEDLQHPWYYMGSGNEYVWTSSOULABUDOWBMDRKLDNOWUMBFXESNFHBEUIXFXGNUJBKWEYJUGMPXIXONQNKDWIIVOGCFACLZZXWKDRDKRRJXGYLCAPWSWWPWFFPICTUOVHPMUNXNKVRTPKWXDEXYGFWFPYYCDBZVKYCMGMCKDVLJOJJFFCSHGXYXZCPTBORTDL",
		},
		{
			name: "UKW-D (rewirable reflector)",
			spec: enigmaSpec{
				model:           M4UKWD,
				rotorConfig:     "I II III | D U Z | 17 5 8",
				reflectorConfig: "D | | AQ BG CK DI EL FX HZ MW NV OT PU RS",
				plugboardConfig: "",
			},
			text: "THERESQQTWOQQMISSINGQQPIECESQQFIRSTQQTHEQQRINGQQSETTINGQQCHANGESQQTHEQQOUTPUTQQLETTERQQITQQDOESNTQQROTATEQQTHEQQWHOLEQQEXITQQPATTERNQQSECONDQQTHEQQROTORSQQAREQQADVANCEDQQBEFOREQQTHEQQLETTERQQISQQENCRYPTED",
			want: "KRHAIKWYFOKTFNNPVCDJAFHFUGNFNIILPGSIURPSCJUKRWKJNBOJFDNHNGVVEJMLFEGQOEMQKFHHCMLPCDMVDXOADJYQTVQWASKPCDSOFVVLIABJHVCEDRRGZVIKWDWCBVJXUZUMGZUEWFBWVDSMPXLYJKCQHLWCYNGTRUWUFWDHGAOPLNOAZIPNRYSGPZWHDYTUBYWBZZIS",
		},
		{
			name: "Commercial",
			spec: enigmaSpec{
				model:           Commercial,
				rotorConfig:     "III-K I-K II-K | G Z J | 6 18 4",
				reflectorConfig: " | Y | ",
			},
			text: "WHENQQBLETCHLEYQQPARKQQWASQQFIRSTQQOPENEDQQASQQAQQMUSEUMQQAROUNDQQTWOQQTHOUSANDQQTHEYQQHADQQANQQENIGMAQQONQQDISPLAYQQTHATQQCOULDQQBEQQTOUCHEDQQBYQQTHEQQPUBLICQQITQQWASQQPARTQQOFQQTHEQQSOCALLEDQQCRYPTOQQTRAILQQTHATQQALLOWEDQQVISITORSQQTOQQFOLLOWQQTHEQQFLOWQQOFQQANQQENIGMAQQMESSAGE",
			want: "ISZZXPSFLMUMSNFXOGHEQIINTXJCAHQLBRELBJWAQWRJIUWUJILFKOPUOLUEXOKVFXLQCOKGNKVHYLBGDRYNGOPVQWIXNVXHOYDEAULBABSTTTZMRCFGXVFSOFZQPKRQKGKREOAXYLCBCZRHMUIRCHCGCNQIEABYWSNWMHOJVQGHWZETBYKBWJMLPRWKMNDMMARELELXKEFIWREMOSJLFESCDCRVOWVVFAMDUAQBRFQLRILGAZYCEPIIZLSXMWPJJMLHRGGMWCYDCTKCEOQJGMZC",
		},
		{
			name: "Swiss-K (movable reflector)",
			spec: enigmaSpec{
				model:           SwissK,
				rotorConfig:     "II-SK I-SK III-SK | A X L | 2 19 4",
				reflectorConfig: " | F | ",
			},
			text: "ALLQQENIGMAQQKQQMACHINESQQWEREQQDELIVEREDQQBYQQTHEQQGERMANSQQWITHQQTHEQQSTANDARDQQCOMMERCIALQQWHEELQQWIRINGQQALSOQQKNOWNQQFROMQQTHEQQENIGMAQQDQQSEEQQTHEQQTABLEQQBELOWQQIMMEDIATELYQQAFTERQQRECEPTIONQQHOWEVERQQTHEQQSWISSQQCHANGEDQQTHEQQWIRINGQQOFQQALLQQCIPHERQQWHEELS",
			want: "MKXPZMCGHRSVAMKALKDJGSRJIKZRPPCFUHWOOBGXKAFQSRFFWMXOVWGVKUKJIJKWVGIGSYIUNYFACJOUGRTQIZSZRTNNHKNGHSIETRPWLKXLSMGOIBPZSYUPIECXWHINIJSRMBMJRJHOOEABFWJZHMXGCXICDNFNVLNIPJGDXIDVEHXSPGDMGEWCCYUGXXBIHJLUXTSMRKIZVDDGNDGLHJHOXVZSYOVPVCYOOBPFYVENEQQXGIXAILVHSSVXAZURPZMLCPFEJ",
		},
		{
			name: "Tripitz",
			spec: enigmaSpec{
				model:       Tripitz,
				rotorConfig: "III-T VIII-T I-T | W W W | 13 25 2",
			},
			text: "THEQQENIGMAQQTQQTIRPITZQQWASQQAQQSPECIALQQVERSIONQQOFQQTHEQQENIGMAQQKQQTHATQQWASQQMADEQQFORQQTHEQQJAPANESEQQARMYQQDURINGQQWWIIQQTHEQQWHEELSQQWEREQQWIREDQQDIFFERENTLYQQANDQQEACHQQHADQQFIVEQQTURNOVERQQNOTCHESQQQQTHEQQTABLEQQBELOWQQSHOWSQQTHEQQWIRINGQQOFQQTHEQQWHEELSQQTHEQQETWQQANDQQUKW",
			want: "NSLLDBIGRLEJHUKZRVIOYXAPGYDZLIKWILEVAGJKXBJBQTMTKSHSHXPVCJYUWJFLPHSJQIGEUBIKHPBONFFBHYTSIHJCUDFOPNEYTVLBVCWIGXADLLZRFGHCNCYHMPYGFJONRBXMAQANGKXOLZLTXMVWHNZLQDNJQDLXGATRRNGOIHNQMKVYPJFUSAPIAQDHVJUATOXYFSNTVWEHIYXEXZJMGICNRLDKKNEAWGRHKDRNBCLSTJFXNZYBCEGBWCSRLCIRAOHYNHEDCEIZILFMTAPMGEFD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := createEnigma(tt.spec.model, tt.spec.rotorConfig, tt.spec.reflectorConfig, tt.spec.plugboardConfig)
			if err != nil {
				t.Errorf("config error = %v", err)
				return
			}
			got, err := e.Encode(tt.text)
			if err != nil {
				t.Errorf("encode error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("want = %v\n got = %v", tt.want, got)
			}
		})
	}
}

func TestEnigma_Decode(t *testing.T) {
	models := []struct {
		name string
		spec enigmaSpec
	}{
		{
			name: "M3",
			spec: enigmaSpec{
				model:           M3,
				rotorConfig:     "VI I V | B T H | 12 1 5",
				reflectorConfig: "C | | ",
				plugboardConfig: "AB CD EF GH IJ KL MN OP", // partial plug boards are allowed
			},
		},
		{
			name: "UKW-D",
			spec: enigmaSpec{
				model:           M4UKWD,
				rotorConfig:     "II VIII I | Q N G | 11 11 9",
				reflectorConfig: "D | | IQ HG CL DA NK FX BZ MW EV OR PU TS",
				plugboardConfig: "ZY WV JF ES LO",
			},
		},
	}
	texts := []string{
		"Simple text with punctuation, nothing special.",
		"Words with special characters on edges, like complex,xerox, vixen opaque Everest, or quirk, or vast molotov. Yearly-yield is also very life-threatening.",
		"The Enigma machine is a cipher device developed and used in the early to mid-twentieth century to protect commercial, diplomatic, and military communication. It was employed extensively by Nazi Germany during World War II, in all branches of the German military. The Enigma machine was considered so secure that it was used to encipher the most top-secret messages. The Enigma has an electromechanical rotor mechanism that scrambles the twenty-six letters of the alphabet. In typical use, one person enters text on the Enigmas keyboard and another person writes down which of twenty-six lights above the keyboard illuminated at each key press. If plain text is entered, the illuminated letters are the encoded ciphertext. Entering ciphertext transforms it back into readable plaintext. The rotor mechanism changes the electrical connections between the keys and the lights with each keypress. The security of the system depends on machine settings that were generally changed daily, based on secret key lists distributed in advance, and on other settings that were changed for each message. The receiving station would have to know and use the exact settings employed by the transmitting station to successfully decrypt a message. While Nazi Germany introduced a series of improvements to Enigma over the years, and these hampered decryption efforts, they did not prevent Poland from cracking the machine prior to the war, enabling the Allies to exploit Enigma-enciphered messages as a major source of intelligence. Many commentators say the flow of Ultra communications intelligence from the decryption of Enigma, Lorenz, and other ciphers, shortened the war substantially, and might even have altered its outcome.",
	}

	for _, model := range models {
		t.Run(model.name, func(t *testing.T) {
			e, err := createEnigma(model.spec.model, model.spec.rotorConfig, model.spec.reflectorConfig, model.spec.plugboardConfig)
			if err != nil {
				t.Errorf("config error = %v", err)
				return
			}
			for i, text := range texts {
				e.RotorsReset()
				encoded, err := e.Encode(Preprocess(text))
				if err != nil {
					if err != nil {
						t.Errorf("encode error = %v", err)
						return
					}
				}

				e.RotorsReset()
				decoded, err := e.Encode(encoded)
				if err != nil {
					if err != nil {
						t.Errorf("encode error while decoding = %v", err)
						return
					}
				}

				want := strings.ToUpper(text)
				got := Postprocess(decoded)
				if got != want {
					t.Errorf("Enigma spec %s, text number %d encode-decode error\nwant = %v\n got = %v", model.name, i, want, got)
				}
			}
		})
	}
}

func TestEnigma_ConfigurationError(t *testing.T) {
	tests := []struct {
		name string
		spec enigmaSpec
	}{
		{
			name: "unsupported model",
			spec: enigmaSpec{"M5", "", "", ""},
		},
		{
			name: "unsupported rotor",
			spec: enigmaSpec{M3, "I II I-K | |", "", ""},
		},
		{
			name: "unsupported 4th rotor",
			spec: enigmaSpec{M4, "V I II III | |", "", ""},
		},
		{
			name: "duplicate rotor model",
			spec: enigmaSpec{M3, "I II II | |", "", ""},
		},
		{
			name: "invalid rotor wheel position",
			spec: enigmaSpec{M3, "I II III | A X x | 1 2 3", "", ""},
		},
		{
			name: "invalid rotor ring position",
			spec: enigmaSpec{M3, "I II III | A B C | 27 6 12", "", ""},
		},
		{
			name: "unsupported reflector",
			spec: enigmaSpec{Commercial, "", "B | |", ""},
		},
		{
			name: "unsupported reflector position",
			spec: enigmaSpec{M3, "", "B | X |", ""},
		},
		{
			name: "invalid reflector position",
			spec: enigmaSpec{M4, "", "B | x |", ""},
		},
		{
			name: "unsupported reflector wiring",
			spec: enigmaSpec{M3, "", "B | | AQ BG CK DI EL FX HZ MW NV OT PU RS", ""},
		},
		{
			name: "invalid reflector wiring (wrong format)",
			spec: enigmaSpec{M4UKWD, "", "D | | AQB BG CK DI EL FX HZ MW NV OT PU RS", ""},
		},
		{
			name: "invalid reflector wiring (incomplete)",
			spec: enigmaSpec{M4UKWD, "", "D | | AQ BG CK", ""},
		},
		{
			name: "invalid reflector wiring (unsupported letter)",
			spec: enigmaSpec{M4UKWD, "", "D | | AQ BG cK DI EL FX HZ MW NV OT PU RS", ""},
		},
		{
			name: "invalid reflector wiring (fixed letter)",
			spec: enigmaSpec{M4UKWD, "", "D | | YJ BG CK DI EL FX HZ MW NV OT PU RS", ""},
		},
		{
			name: "invalid reflector wiring (duplicate)",
			spec: enigmaSpec{M4UKWD, "", "D | | AQ BQ CK DI EL FX HZ MW NV OT PU RS", ""},
		},
		{
			name: "unsupported plugboard",
			spec: enigmaSpec{SwissK, "", "", "AB CD EF"},
		},
		{
			name: "invalid plugboard wiring (wrong format)",
			spec: enigmaSpec{M3, "", "", "ABC DE FG"},
		},
		{
			name: "invalid plugboard wiring (unsupported letter)",
			spec: enigmaSpec{M3, "", "", "Ai BX CU DF EN GQ HM JL KT OP"},
		},
		{
			name: "invalid plugboard wiring (duplicate)",
			spec: enigmaSpec{M4, "", "", "AI AI CU DF EN GQ HM JL KT OP"},
		},
		{
			name: "invalid plugboard wiring (to itself)",
			spec: enigmaSpec{M4, "", "", "AA BX CU DF EN GQ HM JL KT OP"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := createEnigma(tt.spec.model, tt.spec.rotorConfig, tt.spec.reflectorConfig, tt.spec.plugboardConfig)
			if err == nil {
				t.Errorf("expected config error, got none")
			}
		})
	}
}

func TestEnigma_EncodeAlphabet(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{
			name: "unsupported letter",
			text: "SOMEQQTEXTQQWTIHQQINVALIDČQQLETTER",
		},
		{
			name: "number",
			text: "INQQTHEQQ1STQQCENTURY",
		},
		{
			name: "unsupported symbol",
			text: "PUNCTUATION.ISQQNOTQQSUPPORTED",
		},
		{
			name: "unsupported symbol with preprocess",
			text: Preprocess("Only some basic punctuation symbols are supported by Preprocess()."),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := createEnigma(M3, "I II III | B C D | 3 4 5", "C | |", "AB CD EF")
			if err != nil {
				t.Errorf("config error = %v", err)
				return
			}

			_, err = e.Encode(tt.text)
			if err == nil {
				t.Errorf("expected encode error, got none")
			}
		})
	}
}

func createEnigma(model Model, rotorConfigString, reflectorConfigString, plugboardConfig string) (Enigma, error) {
	// Rotors
	rotorsConfig := make(map[RotorSlot]RotorConfig)
	if rotorConfigString != "" {
		conf := strings.Split(rotorConfigString, "|")
		rotorModels := parseConfig(conf[0])
		wheelPositions := parseConfig(conf[1])
		ringPositions := parseConfig(conf[2])

		slots := []RotorSlot{Fourth, Left, Middle, Right}
		firstSlotIndex := 0
		if !model.HasRotorSlot(Fourth) {
			firstSlotIndex = 1
		}
		for si, i := firstSlotIndex, 0; si < len(slots); si, i = si+1, i+1 {
			config := RotorConfig{}
			if rotorModels != nil {
				config.Model = RotorModel(rotorModels[i])
			}
			if wheelPositions != nil {
				config.WheelPosition = wheelPositions[i][0]
			}
			if ringPositions != nil {
				config.RingPosition, _ = strconv.Atoi(ringPositions[i])
			}
			rotorsConfig[slots[si]] = config
		}
	}

	// Reflector
	reflectorConfig := ReflectorConfig{}
	if reflectorConfigString != "" {
		conf := strings.Split(reflectorConfigString, "|")
		refModel := ReflectorModel(strings.TrimSpace(conf[0]))
		if refModel != "" {
			reflectorConfig.Model = refModel
		}
		refPosition := strings.TrimSpace(conf[1])
		if refPosition != "" {
			reflectorConfig.WheelPosition = refPosition[0]
		}
		refWiring := strings.TrimSpace(conf[2])
		if refWiring != "" {
			reflectorConfig.Wiring = refWiring
		}
	}

	return NewEnigmaWithSetup(model, rotorsConfig, reflectorConfig, plugboardConfig)
}

func parseConfig(configString string) []string {
	trimmed := strings.TrimSpace(configString)
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, " ")
}
