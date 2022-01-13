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
		name   string
		fields enigmaSpec
		text   string
		want   string
	}{
		{
			name: "basic M3 with double-stepping",
			fields: enigmaSpec{
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
			fields: enigmaSpec{
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
			fields: enigmaSpec{
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
			fields: enigmaSpec{
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
			fields: enigmaSpec{
				model:           Commercial,
				rotorConfig:     "III-K I-K II-K | G Z J | 6 18 4",
				reflectorConfig: " | Y | ",
			},
			text: "WHENQQBLETCHLEYQQPARKQQWASQQFIRSTQQOPENEDQQASQQAQQMUSEUMQQAROUNDQQTWOQQTHOUSANDQQTHEYQQHADQQANQQENIGMAQQONQQDISPLAYQQTHATQQCOULDQQBEQQTOUCHEDQQBYQQTHEQQPUBLICQQITQQWASQQPARTQQOFQQTHEQQSOCALLEDQQCRYPTOQQTRAILQQTHATQQALLOWEDQQVISITORSQQTOQQFOLLOWQQTHEQQFLOWQQOFQQANQQENIGMAQQMESSAGE",
			want: "ISZZXPSFLMUMSNFXOGHEQIINTXJCAHQLBRELBJWAQWRJIUWUJILFKOPUOLUEXOKVFXLQCOKGNKVHYLBGDRYNGOPVQWIXNVXHOYDEAULBABSTTTZMRCFGXVFSOFZQPKRQKGKREOAXYLCBCZRHMUIRCHCGCNQIEABYWSNWMHOJVQGHWZETBYKBWJMLPRWKMNDMMARELELXKEFIWREMOSJLFESCDCRVOWVVFAMDUAQBRFQLRILGAZYCEPIIZLSXMWPJJMLHRGGMWCYDCTKCEOQJGMZC",
		},
		{
			name: "Swiss-K (movable reflector)",
			fields: enigmaSpec{
				model:           SwissK,
				rotorConfig:     "II-SK I-SK III-SK | A X L | 2 19 4",
				reflectorConfig: " | F | ",
			},
			text: "ALLQQENIGMAQQKQQMACHINESQQWEREQQDELIVEREDQQBYQQTHEQQGERMANSQQWITHQQTHEQQSTANDARDQQCOMMERCIALQQWHEELQQWIRINGQQALSOQQKNOWNQQFROMQQTHEQQENIGMAQQDQQSEEQQTHEQQTABLEQQBELOWQQIMMEDIATELYQQAFTERQQRECEPTIONQQHOWEVERQQTHEQQSWISSQQCHANGEDQQTHEQQWIRINGQQOFQQALLQQCIPHERQQWHEELS",
			want: "MKXPZMCGHRSVAMKALKDJGSRJIKZRPPCFUHWOOBGXKAFQSRFFWMXOVWGVKUKJIJKWVGIGSYIUNYFACJOUGRTQIZSZRTNNHKNGHSIETRPWLKXLSMGOIBPZSYUPIECXWHINIJSRMBMJRJHOOEABFWJZHMXGCXICDNFNVLNIPJGDXIDVEHXSPGDMGEWCCYUGXXBIHJLUXTSMRKIZVDDGNDGLHJHOXVZSYOVPVCYOOBPFYVENEQQXGIXAILVHSSVXAZURPZMLCPFEJ",
		},
		{
			name: "Tripitz",
			fields: enigmaSpec{
				model:       Tripitz,
				rotorConfig: "III-T VIII-T I-T | W W W | 13 25 2",
			},
			text: "THEQQENIGMAQQTQQTIRPITZQQWASQQAQQSPECIALQQVERSIONQQOFQQTHEQQENIGMAQQKQQTHATQQWASQQMADEQQFORQQTHEQQJAPANESEQQARMYQQDURINGQQWWIIQQTHEQQWHEELSQQWEREQQWIREDQQDIFFERENTLYQQANDQQEACHQQHADQQFIVEQQTURNOVERQQNOTCHESQQQQTHEQQTABLEQQBELOWQQSHOWSQQTHEQQWIRINGQQOFQQTHEQQWHEELSQQTHEQQETWQQANDQQUKW",
			want: "NSLLDBIGRLEJHUKZRVIOYXAPGYDZLIKWILEVAGJKXBJBQTMTKSHSHXPVCJYUWJFLPHSJQIGEUBIKHPBONFFBHYTSIHJCUDFOPNEYTVLBVCWIGXADLLZRFGHCNCYHMPYGFJONRBXMAQANGKXOLZLTXMVWHNZLQDNJQDLXGATRRNGOIHNQMKVYPJFUSAPIAQDHVJUATOXYFSNTVWEHIYXEXZJMGICNRLDKKNEAWGRHKDRNBCLSTJFXNZYBCEGBWCSRLCIRAOHYNHEDCEIZILFMTAPMGEFD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := createEnigma(tt.fields.model, tt.fields.rotorConfig, tt.fields.reflectorConfig, tt.fields.plugboardConfig)
			if err != nil {
				t.Errorf("config error = %v", err)
				return
			}
			got, err := e.Encode(tt.text)
			if err != nil {
				t.Errorf("Encode() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Encode()\nwant = %v\n got = %v", tt.want, got)
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
						t.Errorf("Encode() error while encoding = %v", err)
						return
					}
				}

				e.RotorsReset()
				decoded, err := e.Encode(encoded)
				if err != nil {
					if err != nil {
						t.Errorf("Encode() error while decoding = %v", err)
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

// todo - error cases (config & encode errors)

func createEnigma(model Model, rotorConfigString string, reflectorConfigString string, plugboardConfig string) (Enigma, error) {
	// Rotors
	conf := strings.Split(rotorConfigString, "|")
	rotorTypes := parseConfig(conf[0])
	wheelPositions := parseConfig(conf[1])
	ringPositions := parseConfig(conf[2])

	slots := []RotorSlot{Fourth, Left, Middle, Right}
	firstSlotIndex := 0
	if model.GetRotorCount() == 3 {
		firstSlotIndex = 1
	}
	rotorsConfig := make(map[RotorSlot]RotorConfig)
	for si, i := firstSlotIndex, 0; si < len(slots); si, i = si+1, i+1 {
		config := RotorConfig{}
		if rotorTypes != nil {
			config.RotorType = RotorType(rotorTypes[i])
		}
		if wheelPositions != nil {
			config.WheelPosition = wheelPositions[i][0]
		}
		if ringPositions != nil {
			config.RingPosition, _ = strconv.Atoi(ringPositions[i])
		}
		rotorsConfig[slots[si]] = config
	}

	// Reflector
	reflectorConfig := ReflectorConfig{}
	if reflectorConfigString != "" {
		conf = strings.Split(reflectorConfigString, "|")
		refType := ReflectorType(strings.TrimSpace(conf[0]))
		if refType != "" {
			reflectorConfig.ReflectorType = refType
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
