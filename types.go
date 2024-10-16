package tradingview

// SocketInterface ...
type SocketInterface interface {
	AddSymbol(symbol string) error
	RemoveSymbol(symbol string) error
	Init() error
	Close() error
}

// SocketMessage ...
type SocketMessage struct {
	Message string      `json:"m"`
	Payload interface{} `json:"p"`
}

// QuoteMessage ...
type QuoteMessage struct {
	Symbol string     `mapstructure:"n"`
	Status string     `mapstructure:"s"`
	Data   *QuoteData `mapstructure:"v"`
}

// QuoteData ...
// QuoteData ...
// QuoteData ...
type QuoteData struct {
	Price           *float64  `mapstructure:"lp"`
	Bid             *float64  `mapstructure:"bid"`
	Ask             *float64  `mapstructure:"ask"`
	PriceChange     *float64  `mapstructure:"ch"`
	PriceChangePerc *float64  `mapstructure:"chp"`
	LastPriceTime   *int64    `mapstructure:"lp_time"`
	PriceScale      *int      `mapstructure:"pricescale"` // Yeni alan eklendi
	OriginalName    *string   `mapstructure:"original_name"`
	ProName         *string   `mapstructure:"pro_name"`
	ShortName       *string   `mapstructure:"short_name"`
	Type            *string   `mapstructure:"type"`
	TypeSpecs       *[]string `mapstructure:"typespecs"`
	UpdateMode      *string   `mapstructure:"update_mode"`
	Volume          *float64  `mapstructure:"volume"`
}

type DUMessage struct {
	Symbol  string  `mapstructure:"n"`     // Sembol
	BarData []DUBar `mapstructure:"sds_1"` // Bar verileri
}

// DUBar represents an individual bar (OHLC data)
type DUBar struct {
	Timestamp float64 `mapstructure:"i"`    // Zaman damgası
	Open      float64 `mapstructure:"v[0]"` // Açılış fiyatı
	High      float64 `mapstructure:"v[1]"` // En yüksek fiyat
	Low       float64 `mapstructure:"v[2]"` // En düşük fiyat
	Close     float64 `mapstructure:"v[3]"` // Kapanış fiyatı
	Volume    float64 `mapstructure:"v[4]"` // Hacim
}

// Flags ...
type Flags struct {
	Flags []string `json:"flags"`
}

// OnReceiveDataCallback ...
type OnReceiveDataCallback func(symbol string, data *QuoteData)

// OnErrorCallback ...
type OnErrorCallback func(err error, context string)
