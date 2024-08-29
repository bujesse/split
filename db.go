package main

import (
	"log"
	"os"
	"split/config/logger"
	"split/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}

	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,       // Slow SQL threshold
			LogLevel:      gormLogger.Silent, // Log level
			// LogLevel:                  gormLogger.Info,   // Log level
			IgnoreRecordNotFoundError: true,  // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,  // Don't include params in the SQL log
			Colorful:                  false, // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("split.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Fatal("🔥 failed to connect to the database: %s", err.Error())
	}

	logger.Debug.Println("🚀 Connected Successfully to the Database")

	return db
}

func MakeMigrations() error {
	db := GetConnection()

	err := db.AutoMigrate(
		&models.Category{},
		&models.Currency{},
		&models.Settlement{},
		&models.Expense{},
		&models.ExpenseOwed{},
		&models.User{},
		&models.FxRate{},
	)
	if err != nil {
		logger.Fatal("failed to migrate database schema: %v", err)
	}

	seedCurrencies(db)

	return nil
}

func seedCurrencies(db *gorm.DB) {
	var count int64
	db.Model(&models.Currency{}).Count(&count)

	if count == 0 {
		currencies := []models.Currency{
			{Code: "ADA", Name: "Cardano", LatestFxRateUSD: 2.7848696578},
			{Code: "AED", Name: "United Arab Emirates Dirham", LatestFxRateUSD: 3.6719906741},
			{Code: "AFN", Name: "Afghan Afghani", LatestFxRateUSD: 70.9078290632},
			{Code: "ALL", Name: "Albanian Lek", LatestFxRateUSD: 89.4135964048},
			{Code: "AMD", Name: "Armenian Dram", LatestFxRateUSD: 387.2598943342},
			{Code: "ANG", Name: "Netherlands Antillean Guilder", LatestFxRateUSD: 1.7906703374},
			{Code: "AOA", Name: "Angolan Kwanza", LatestFxRateUSD: 904.4885337335},
			{Code: "ARS", Name: "Argentine Peso", LatestFxRateUSD: 950.6361911163},
			{Code: "AUD", Name: "Australian Dollar", LatestFxRateUSD: 1.4750001665},
			{Code: "AWG", Name: "Aruban Florin", LatestFxRateUSD: 1.79},
			{Code: "AZN", Name: "Azerbaijani Manat", LatestFxRateUSD: 1.7},
			{
				Code:            "BAM",
				Name:            "Bosnia-Herzegovina Convertible Mark",
				LatestFxRateUSD: 1.7562802586,
			},
			{Code: "BBD", Name: "Barbadian Dollar", LatestFxRateUSD: 2},
			{Code: "BDT", Name: "Bangladeshi Taka", LatestFxRateUSD: 119.6123974291},
			{Code: "BGN", Name: "Bulgarian Lev", LatestFxRateUSD: 1.7511202874},
			{Code: "BHD", Name: "Bahraini Dinar", LatestFxRateUSD: 0.376},
			{Code: "BIF", Name: "Burundian Franc", LatestFxRateUSD: 2887.4009588629},
			{Code: "BMD", Name: "Bermudian Dollar", LatestFxRateUSD: 1},
			{Code: "BND", Name: "Brunei Dollar", LatestFxRateUSD: 1.3025902247},
			{Code: "BOB", Name: "Bolivian Boliviano", LatestFxRateUSD: 6.940600843},
			{Code: "BRL", Name: "Brazilian Real", LatestFxRateUSD: 5.5075905621},
			{Code: "BSD", Name: "Bahamian Dollar", LatestFxRateUSD: 1},
			{Code: "BTN", Name: "Bhutanese Ngultrum", LatestFxRateUSD: 84.1010814152},
			{Code: "BWP", Name: "Botswana Pula", LatestFxRateUSD: 13.2782418289},
			{Code: "BYN", Name: "Belarusian Ruble", LatestFxRateUSD: 3.2702341953},
			{Code: "BZD", Name: "Belize Dollar", LatestFxRateUSD: 2},
			{Code: "CAD", Name: "Canadian Dollar", LatestFxRateUSD: 1.3466401488},
			{Code: "CDF", Name: "Congolese Franc", LatestFxRateUSD: 2801.6365572201},
			{Code: "CHF", Name: "Swiss Franc", LatestFxRateUSD: 0.8434201645},
			{Code: "CLP", Name: "Chilean Peso", LatestFxRateUSD: 910.9512938298},
			{Code: "CNY", Name: "Chinese Yuan", LatestFxRateUSD: 7.1258713654},
			{Code: "COP", Name: "Colombian Peso", LatestFxRateUSD: 4032.2542326529},
			{Code: "CRC", Name: "Costa Rican Colón", LatestFxRateUSD: 525.1028424602},
			{Code: "CUP", Name: "Cuban Peso", LatestFxRateUSD: 24},
			{Code: "CVE", Name: "Cape Verdean Escudo", LatestFxRateUSD: 99.1047579044},
			{Code: "CZK", Name: "Czech Koruna", LatestFxRateUSD: 22.5105039984},
			{Code: "DJF", Name: "Djiboutian Franc", LatestFxRateUSD: 177.721},
			{Code: "DKK", Name: "Danish Krone", LatestFxRateUSD: 6.7032712486},
			{Code: "DOP", Name: "Dominican Peso", LatestFxRateUSD: 59.4223674467},
			{Code: "DZD", Name: "Algerian Dinar", LatestFxRateUSD: 134.0401654643},
			{Code: "EGP", Name: "Egyptian Pound", LatestFxRateUSD: 48.6943557047},
			{Code: "ERN", Name: "Eritrean Nakfa", LatestFxRateUSD: 15},
			{Code: "ETB", Name: "Ethiopian Birr", LatestFxRateUSD: 111.3553035278},
			{Code: "EUR", Name: "Euro", LatestFxRateUSD: 0.8986501054},
			{Code: "FJD", Name: "Fijian Dollar", LatestFxRateUSD: 2.2033902849},
			{Code: "FKP", Name: "Falkland Islands Pound", LatestFxRateUSD: 0.7568269555},
			{Code: "GBP", Name: "British Pound Sterling", LatestFxRateUSD: 0.7567701029},
			{Code: "GEL", Name: "Georgian Lari", LatestFxRateUSD: 2.7049103874},
			{Code: "GHS", Name: "Ghanaian Cedi", LatestFxRateUSD: 15.639352684},
			{Code: "GIP", Name: "Gibraltar Pound", LatestFxRateUSD: 0.7568270064},
			{Code: "GMD", Name: "Gambian Dalasi", LatestFxRateUSD: 56.05686774},
			{Code: "GNF", Name: "Guinean Franc", LatestFxRateUSD: 8614.9657353233},
			{Code: "GTQ", Name: "Guatemalan Quetzal", LatestFxRateUSD: 7.7369612221},
			{Code: "GYD", Name: "Guyanese Dollar", LatestFxRateUSD: 208.7687009967},
			{Code: "HKD", Name: "Hong Kong Dollar", LatestFxRateUSD: 7.7957609828},
			{Code: "HNL", Name: "Honduran Lempira", LatestFxRateUSD: 24.8302938213},
			{Code: "HRK", Name: "Croatian Kuna", LatestFxRateUSD: 6.4064111421},
			{Code: "HTG", Name: "Haitian Gourde", LatestFxRateUSD: 134.2725038572},
			{Code: "HUF", Name: "Hungarian Forint", LatestFxRateUSD: 353.4550586598},
			{Code: "IDR", Name: "Indonesian Rupiah", LatestFxRateUSD: 15410.737485805},
			{Code: "ILS", Name: "Israeli New Shekel", LatestFxRateUSD: 3.6624003772},
			{Code: "INR", Name: "Indian Rupee", LatestFxRateUSD: 83.9096237085},
			{Code: "IQD", Name: "Iraqi Dinar", LatestFxRateUSD: 1311.1072907009},
			{Code: "IRR", Name: "Iranian Rial", LatestFxRateUSD: 42097.462809666},
			{Code: "ISK", Name: "Icelandic Króna", LatestFxRateUSD: 137.0278961392},
			{Code: "JMD", Name: "Jamaican Dollar", LatestFxRateUSD: 157.0148804008},
			{Code: "JOD", Name: "Jordanian Dinar", LatestFxRateUSD: 0.71},
			{Code: "JPY", Name: "Japanese Yen", LatestFxRateUSD: 144.2905964449},
			{Code: "KES", Name: "Kenyan Shilling", LatestFxRateUSD: 128.7513188621},
			{Code: "KGS", Name: "Kyrgyzstani Som", LatestFxRateUSD: 85.1174209361},
			{Code: "KHR", Name: "Cambodian Riel", LatestFxRateUSD: 4056.5008097807},
			{Code: "KMF", Name: "Comorian Franc", LatestFxRateUSD: 442.4986271866},
			{Code: "KPW", Name: "North Korean Won", LatestFxRateUSD: 900.0051632123},
			{Code: "KRW", Name: "South Korean Won", LatestFxRateUSD: 1332.3270548667},
			{Code: "KWD", Name: "Kuwaiti Dinar", LatestFxRateUSD: 0.306190043},
			{Code: "KYD", Name: "Cayman Islands Dollar", LatestFxRateUSD: 0.83333},
			{Code: "KZT", Name: "Kazakhstani Tenge", LatestFxRateUSD: 480.8160313315},
			{Code: "LAK", Name: "Lao Kip", LatestFxRateUSD: 22118.801478198},
			{Code: "LBP", Name: "Lebanese Pound", LatestFxRateUSD: 89762.460178053},
			{Code: "LKR", Name: "Sri Lankan Rupee", LatestFxRateUSD: 300.3186415391},
			{Code: "LRD", Name: "Liberian Dollar", LatestFxRateUSD: 195.7749977456},
			{Code: "LSL", Name: "Lesotho Loti", LatestFxRateUSD: 17.7668621957},
			{Code: "LTL", Name: "Lithuanian Litas", LatestFxRateUSD: 3.1035636234},
			{Code: "LVL", Name: "Latvian Lats", LatestFxRateUSD: 0.6317146098},
			{Code: "LYD", Name: "Libyan Dinar", LatestFxRateUSD: 4.7691306499},
			{Code: "MAD", Name: "Moroccan Dirham", LatestFxRateUSD: 9.652041577},
			{Code: "MDL", Name: "Moldovan Leu", LatestFxRateUSD: 17.5491331246},
			{Code: "MGA", Name: "Malagasy Ariary", LatestFxRateUSD: 4565.2302310972},
			{Code: "MKD", Name: "Macedonian Denar", LatestFxRateUSD: 55.0697497254},
			{Code: "MMK", Name: "Myanmar Kyat", LatestFxRateUSD: 2099.6893561992},
			{Code: "MNT", Name: "Mongolian Tugrik", LatestFxRateUSD: 3384.3225416266},
			{Code: "MOP", Name: "Macanese Pataca", LatestFxRateUSD: 8.0433810671},
			{Code: "MRO", Name: "Mauritanian Ouguiya", LatestFxRateUSD: 356.999828},
			{Code: "MUR", Name: "Mauritian Rupee", LatestFxRateUSD: 46.1905169497},
			{Code: "MVR", Name: "Maldivian Rufiyaa", LatestFxRateUSD: 15.4850529017},
			{Code: "MWK", Name: "Malawian Kwacha", LatestFxRateUSD: 1738.0148931997},
			{Code: "MXN", Name: "Mexican Peso", LatestFxRateUSD: 19.6518136887},
			{Code: "MYR", Name: "Malaysian Ringgit", LatestFxRateUSD: 4.3447004457},
			{Code: "MZN", Name: "Mozambican Metical", LatestFxRateUSD: 63.7254208987},
			{Code: "NAD", Name: "Namibian Dollar", LatestFxRateUSD: 17.7299218717},
			{Code: "NGN", Name: "Nigerian Naira", LatestFxRateUSD: 1590.4679603308},
			{Code: "NIO", Name: "Nicaraguan Córdoba", LatestFxRateUSD: 36.7902770887},
			{Code: "NOK", Name: "Norwegian Krone", LatestFxRateUSD: 10.5282216768},
			{Code: "NPR", Name: "Nepalese Rupee", LatestFxRateUSD: 134.2732618466},
			{Code: "NZD", Name: "New Zealand Dollar", LatestFxRateUSD: 1.602150244},
			{Code: "OMR", Name: "Omani Rial", LatestFxRateUSD: 0.3850500718},
			{Code: "PAB", Name: "Panamanian Balboa", LatestFxRateUSD: 0.9991401095},
			{Code: "PEN", Name: "Peruvian Nuevo Sol", LatestFxRateUSD: 3.7537204919},
			{Code: "PGK", Name: "Papua New Guinean Kina", LatestFxRateUSD: 3.890620769},
			{Code: "PHP", Name: "Philippine Peso", LatestFxRateUSD: 56.2719497018},
			{Code: "PKR", Name: "Pakistani Rupee", LatestFxRateUSD: 278.2498026441},
			{Code: "PLN", Name: "Polish Zloty", LatestFxRateUSD: 3.8616904152},
			{Code: "PYG", Name: "Paraguayan Guarani", LatestFxRateUSD: 7656.216485634},
			{Code: "QAR", Name: "Qatari Riyal", LatestFxRateUSD: 3.6478204653},
			{Code: "RON", Name: "Romanian Leu", LatestFxRateUSD: 4.4709508435},
			{Code: "RSD", Name: "Serbian Dinar", LatestFxRateUSD: 104.7918684174},
			{Code: "RUB", Name: "Russian Ruble", LatestFxRateUSD: 91.2540498904},
			{Code: "RWF", Name: "Rwandan Franc", LatestFxRateUSD: 1330.4805246363},
			{Code: "SAR", Name: "Saudi Riyal", LatestFxRateUSD: 3.7459205048},
			{Code: "SBD", Name: "Solomon Islands Dollar", LatestFxRateUSD: 8.3183242893},
			{Code: "SCR", Name: "Seychellois Rupee", LatestFxRateUSD: 14.8399123622},
			{Code: "SDG", Name: "Sudanese Pound", LatestFxRateUSD: 601.5},
			{Code: "SEK", Name: "Swedish Krona", LatestFxRateUSD: 10.1918418606},
			{Code: "SGD", Name: "Singapore Dollar", LatestFxRateUSD: 1.3034501871},
			{Code: "SHP", Name: "Saint Helena Pound", LatestFxRateUSD: 0.7567701057},
			{Code: "SLL", Name: "Sierra Leonean Leone", LatestFxRateUSD: 22533.962718626},
			{Code: "SOS", Name: "Somali Shilling", LatestFxRateUSD: 571.3042132791},
			{Code: "SRD", Name: "Surinamese Dollar", LatestFxRateUSD: 28.7205349904},
			{Code: "STD", Name: "São Tomé and Príncipe Dobra", LatestFxRateUSD: 22104.781525979},
			{Code: "SVC", Name: "Salvadoran Colón", LatestFxRateUSD: 8.75},
			{Code: "SYP", Name: "Syrian Pound", LatestFxRateUSD: 13027.036533812},
			{Code: "SZL", Name: "Swazi Lilangeni", LatestFxRateUSD: 17.762312647},
			{Code: "THB", Name: "Thai Baht", LatestFxRateUSD: 34.0030753682},
			{Code: "TJS", Name: "Tajikistani Somoni", LatestFxRateUSD: 10.6843414529},
			{Code: "TMT", Name: "Turkmenistani Manat", LatestFxRateUSD: 3.5},
			{Code: "TND", Name: "Tunisian Dinar", LatestFxRateUSD: 3.0299904339},
			{Code: "TOP", Name: "Tongan Paʻanga", LatestFxRateUSD: 2.3234204034},
			{Code: "TRY", Name: "Turkish Lira", LatestFxRateUSD: 34.0277434116},
			{Code: "TTD", Name: "Trinidad and Tobago Dollar", LatestFxRateUSD: 6.7833110244},
			{Code: "TWD", Name: "New Taiwan Dollar", LatestFxRateUSD: 31.9707737785},
			{Code: "TZS", Name: "Tanzanian Shilling", LatestFxRateUSD: 2717.7868273129},
			{Code: "UAH", Name: "Ukrainian Hryvnia", LatestFxRateUSD: 41.5748353771},
			{Code: "UGX", Name: "Ugandan Shilling", LatestFxRateUSD: 3721.9414868223},
			{Code: "USD", Name: "United States Dollar", LatestFxRateUSD: 1},
			{Code: "UYU", Name: "Uruguayan Peso", LatestFxRateUSD: 40.3702148632},
			{Code: "UZS", Name: "Uzbekistani Som", LatestFxRateUSD: 12676.007117065},
			{Code: "VND", Name: "Vietnamese Dong", LatestFxRateUSD: 24869.195251626},
			{Code: "VUV", Name: "Vanuatu Vatu", LatestFxRateUSD: 118.2436867563},
			{Code: "WST", Name: "Samoan Tala", LatestFxRateUSD: 2.7065744428},
			{Code: "XAF", Name: "Central African CFA Franc", LatestFxRateUSD: 589.2872853481},
			{Code: "XAG", Name: "Silver Ounce", LatestFxRateUSD: 0.0340498766},
			{Code: "XAU", Name: "Gold Ounce", LatestFxRateUSD: 0.0003995997},
			{Code: "XCD", Name: "East Caribbean Dollar", LatestFxRateUSD: 2.7},
			{Code: "XDR", Name: "Special Drawing Rights", LatestFxRateUSD: 0.7408701216},
			{Code: "XOF", Name: "West African CFA Franc", LatestFxRateUSD: 589.2872800705},
			{Code: "XPD", Name: "Palladium Ounce", LatestFxRateUSD: 0.0010383669},
			{Code: "XPF", Name: "CFP Franc", LatestFxRateUSD: 107.1128754341},
			{Code: "XPT", Name: "Platinum Ounce", LatestFxRateUSD: 0.0010629616},
			{Code: "YER", Name: "Yemeni Rial", LatestFxRateUSD: 250.3614001095},
			{Code: "ZAR", Name: "South African Rand", LatestFxRateUSD: 17.7748832091},
			{Code: "ZMW", Name: "Zambian Kwacha", LatestFxRateUSD: 26.0035138049},
			{Code: "ZWL", Name: "Zimbabwean Dollar", LatestFxRateUSD: 34997.477850009},
		}

		for _, currency := range currencies {
			db.Create(&currency)
		}
	}
}
