package enigma

import (
	"strconv"
	"strings"
	"testing"
)

func TestEnigma_Encode(t *testing.T) {
	type fields struct {
		model           Model
		reflectorConfig string // B | 15 | AA BB CC DD EE ...
		rotorConfig     string // I IV VII | A U C | 1 14 3
		plugboardConfig string // AA BB CC DD EE FF GG HH II JJ
	}
	tests := []struct {
		name          string
		fields        fields
		text          string
		want          string
		wantConfigErr bool
		wantEncodeErr bool
	}{
		{
			name: "basic M3 with double-stepping",
			fields: fields{
				model:           M3,
				reflectorConfig: "B | | ",
				rotorConfig:     "I II III | | ",
				plugboardConfig: "",
			},
			text:          strings.Repeat("A", 200),
			want:          "BDZGOWCXLTKSBTMCDLPBMUQOFXYHCXTGYJFLINHNXSHIUNTHEORXPQPKOVHCBUBTZSZSOOSTGOTFSODBBZZLXLCYZXIFGWFDZEEQIBMGFJBWZFCKPFMGBXQCIVIBBRNCOCJUVYDKMVJPFMDRMTGLWFOZLXGJEYYQPVPBWNCKVKLZTCBDLDCTSNRCOOVPTGBVBBISGJSO",
			wantConfigErr: false,
			wantEncodeErr: false,
		},
		{
			name: "M3 with full settings",
			fields: fields{
				model:           M3,
				reflectorConfig: "C | | ",
				rotorConfig:     "III VII VIII | D U S | 12 8 6",
				plugboardConfig: "AI BX CU DF EN GQ HM JL KT OP",
			},
			text:          "LOREMQQIPSUMQQDOLORQQSITQQAMETQQCONSECTETUERQQADIPISCINGQQELITQQAENEANQQVELQQMASSAQQQUISQQMAURISQQVEHICULAQQLACINIAQQCURABITURQQSAGITTISQQHENDRERITQQANTEQQNAMQQQUISQQNULLAQQETIAMQQQUISQQQUAMQQALIQUAMQQINQQLOREMQQSITQQAMETQQLEOQQACCUMSANQQLACINIA",
			want:          "JRKGDLRCOURDHDKHEEOOWVEJVPOOKOBHFFQXDNWDYEDDTKDWLRGLSJMBRRQYHQRUPUBVYHTIABJNKZYPRQVJXTXOZWOSIMQHDYWHUHKZGCVXDIYURDQGOIHNFMMDYMXDPFKXZQTXZMZGYOYBQKIFXPFSXHYBOWRSYQWLXHMIIEHUWPOJJSSBNOSCPELDEENEGTMXWZQRTXCKRQLGFQKBUOBKEBXVGWFYIRSFHBPRAWKIBEPLBMCEW",
			wantConfigErr: false,
			wantEncodeErr: false,
		},
		{
			name: "M4 (4-rotors)",
			fields: fields{
				model:           M4,
				reflectorConfig: "BThin | | ",
				rotorConfig:     "gamma VI I VII | L X A Q | 18 16 23 2",
				plugboardConfig: "",
			},
			text:          "LOREMQQIPSUMQQDOLORQQSITQQAMETQQCONSECTETUERQQADIPISCINGQQELITQQAENEANQQVELQQMASSAQQQUISQQMAURISQQVEHICULAQQLACINIAQQCURABITURQQSAGITTISQQHENDRERITQQANTEQQNAMQQQUISQQNULLAQQETIAMQQQUISQQQUAMQQALIQUAMQQINQQLOREMQQSITQQAMETQQLEOQQACCUMSANQQLACINIA",
			want:          "HZAFDYHADNXFLGTKODHHUCMCKFKFLOSTSMRPZNBLIZSYXGGTEGUHNQQEDLQHPWYYMGSGNEYVWTSSOULABUDOWBMDRKLDNOWUMBFXESNFHBEUIXFXGNUJBKWEYJUGMPXIXONQNKDWIIVOGCFACLZZXWKDRDKRRJXGYLCAPWSWWPWFFPICTUOVHPMUNXNKVRTPKWXDEXYGFWFPYYCDBZVKYCMGMCKDVLJOJJFFCSHGXYXZCPTBORTDL",
			wantConfigErr: false,
			wantEncodeErr: false,
		},
		{
			name: "UKW-D (rewirable reflector)",
			fields: fields{
				model:           M4UKWD,
				reflectorConfig: "D | | AQ BG CK DI EL FX HZ MW NV OT PU RS",
				rotorConfig:     "I II III | D U Z | 17 5 8",
				plugboardConfig: "", //"AB CD EF GI ZS WL QY JX MP VR",
			},
			text:          "THERESQQTWOQQMISSINGQQPIECESQQFIRSTQQTHEQQRINGQQSETTINGQQCHANGESQQTHEQQOUTPUTQQLETTERQQITQQDOESNTQQROTATEQQTHEQQWHOLEQQEXITQQPATTERNQQSECONDQQTHEQQROTORSQQAREQQADVANCEDQQBEFOREQQTHEQQLETTERQQISQQENCRYPTED",
			want:          "KRHAIKWYFOKTFNNPVCDJAFHFUGNFNIILPGSIURPSCJUKRWKJNBOJFDNHNGVVEJMLFEGQOEMQKFHHCMLPCDMVDXOADJYQTVQWASKPCDSOFVVLIABJHVCEDRRGZVIKWDWCBVJXUZUMGZUEWFBWVDSMPXLYJKCQHLWCYNGTRUWUFWDHGAOPLNOAZIPNRYSGPZWHDYTUBYWBZZIS",
			wantConfigErr: false,
			wantEncodeErr: false,
		},
		{
			name: "Swiss-K (movable reflector)",
			fields: fields{
				model:           SwissK,
				reflectorConfig: "K | F | ",
				rotorConfig:     "II-K I-K III-K | A X L | 2 19 4",
				plugboardConfig: "",
			},
			text:          "ALLQQENIGMAQQKQQMACHINESQQWEREQQDELIVEREDQQBYQQTHEQQGERMANSQQWITHQQTHEQQSTANDARDQQCOMMERCIALQQWHEELQQWIRINGQQALSOQQKNOWNQQFROMQQTHEQQENIGMAQQDQQSEEQQTHEQQTABLEQQBELOWQQIMMEDIATELYQQAFTERQQRECEPTIONQQHOWEVERQQTHEQQSWISSQQCHANGEDQQTHEQQWIRINGQQOFQQALLQQCIPHERQQWHEELS",
			want:          "MKXPZMCGHRSVAMKALKDJGSRJIKZRPPCFUHWOOBGXKAFQSRFFWMXOVWGVKUKJIJKWVGIGSYIUNYFACJOUGRTQIZSZRTNNHKNGHSIETRPWLKXLSMGOIBPZSYUPIECXWHINIJSRMBMJRJHOOEABFWJZHMXGCXICDNFNVLNIPJGDXIDVEHXSPGDMGEWCCYUGXXBIHJLUXTSMRKIZVDDGNDGLHJHOXVZSYOVPVCYOOBPFYVENEQQXGIXAILVHSSVXAZURPZMLCPFEJ",
			wantConfigErr: false,
			wantEncodeErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := createEnigma(tt.fields.model, tt.fields.reflectorConfig, tt.fields.rotorConfig, tt.fields.plugboardConfig)
			if (err != nil) != tt.wantConfigErr {
				t.Errorf("config error = %v, wantErr %v", err, tt.wantConfigErr)
				return
			}
			got, err := e.Encode(tt.text)
			if (err != nil) != tt.wantEncodeErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantEncodeErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encode()\nwant = %v\n got = %v", got, tt.want)
			}
		})
	}
}

// todo - encode -> decode (including preprocess + postprocess functions)

// todo - error cases (config & encode errors)

func createEnigma(model Model, reflectorConfigString string, rotorConfigString string, plugboardConfig string) (Enigma, error) {

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
	conf = strings.Split(reflectorConfigString, "|")
	reflectorConfig := ReflectorConfig{}
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

	return NewEnigmaWithSetup(model, rotorsConfig, reflectorConfig, plugboardConfig)
}

func parseConfig(configString string) []string {
	trimmed := strings.TrimSpace(configString)
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, " ")
}
