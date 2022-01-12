package enigma

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestEnigma_Encode(t *testing.T) {
	type fields struct {
		model           model
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

func createEnigma(model model, reflectorConfig string, rotorConfig string, plugboardConfig string) (Enigma, error) {
	e := NewEnigma(model)

	// Reflector
	conf := strings.Split(reflectorConfig, "|")
	refType := strings.TrimSpace(conf[0])
	if refType != "" {
		if err := e.ReflectorSelect(reflectorType(refType)); err != nil {
			return Enigma{}, fmt.Errorf("reflector select error: %w", err)
		}
	}

	refPosition := strings.TrimSpace(conf[1])
	if refPosition != "" {
		if err := e.ReflectorSet(byte(refPosition[0])); err != nil {
			return Enigma{}, fmt.Errorf("reflector set error: %w", err)
		}
	}

	refWiring := strings.TrimSpace(conf[2])
	if refWiring != "" {
		if err := e.ReflectorRewire(refWiring); err != nil {
			return Enigma{}, fmt.Errorf("reflector wiring error: %w", err)
		}
	}

	// Rotors
	conf = strings.Split(rotorConfig, "|")
	rotorTypes := strings.TrimSpace(conf[0])
	slots := []RotorSlot{Fourth, Left, Middle, Right}
	firstRotorSlotIndex := 0
	if e.GetRotorCount() == 3 {
		firstRotorSlotIndex = 1
	}
	if rotorTypes != "" {
		types := make(map[RotorSlot]RotorType)
		i := firstRotorSlotIndex
		for _, rType := range strings.Split(rotorTypes, " ") {
			types[slots[i]] = RotorType(rType)
			i++
		}
		if err := e.RotorsSelect(types); err != nil {
			return Enigma{}, fmt.Errorf("rotor select error: %w", err)
		}
	}

	wheelConfig := strings.TrimSpace(conf[1])
	if wheelConfig != "" {
		i := firstRotorSlotIndex
		for _, val := range strings.Split(wheelConfig, " ") {
			if err := e.RotorSetWheel(slots[i], val[0]); err != nil {
				return Enigma{}, fmt.Errorf("rotor %d wheel set error: %w", slots[i], err)
			}
			i++
		}
	}

	ringConfig := strings.TrimSpace(conf[2])
	if ringConfig != "" {
		i := firstRotorSlotIndex
		for _, val := range strings.Split(ringConfig, " ") {
			pos, err := strconv.Atoi(val)
			if err != nil {
				return Enigma{}, fmt.Errorf("invalid ring config %s: %w", val, err)
			}
			if err := e.RotorSetRing(slots[i], pos); err != nil {
				return Enigma{}, fmt.Errorf("rotor %d ring set error: %w", slots[i], err)
			}
			i++
		}
	}

	// Plugboard
	if plugboardConfig != "" {
		if err := e.PlugboardSet(plugboardConfig); err != nil {
			return Enigma{}, fmt.Errorf("plugboard set error: %w", err)
		}
	}

	return e, nil
}
