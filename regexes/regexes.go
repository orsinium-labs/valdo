package regexes

const (
	ASCII                 = "^[\x00-\x7F]*$"
	ASCIIAlpha            = "^[a-zA-Z]+$"
	ASCIIALphaNum         = "^[a-zA-Z0-9]+$"
	Base32                = "^(?:[A-Z2-7]{8})*(?:[A-Z2-7]{2}={6}|[A-Z2-7]{4}={4}|[A-Z2-7]{5}={3}|[A-Z2-7]{7}=|[A-Z2-7]{8})$"
	Base64                = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	Base64RawURL          = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2,4})$"
	Base64URL             = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	BIC                   = `^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`
	BtcAddress            = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$` // bitcoin address
	BtcAddressLowerBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`       // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	BtcAddressUpperBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`       // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	Cron                  = `(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})`
	CVE                   = `^CVE-(1999|2\d{3})-(0[^0]\d{2}|0\d[^0]\d{1}|0\d{2}[^0]|[1-9]{1}\d{3,})$` // CVE Format Id https://cve.mitre.org/cve/identifiers/syntaxchange.html
	DataURI               = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	DnsRFC1035Label       = "^[a-z]([-a-z0-9]*[a-z0-9]){0,62}$"
	E164                  = "^\\+[1-9]?[0-9]{7,14}$"
	Email                 = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	EthAddress            = `^0x[0-9a-fA-F]{40}$`
	EthAddressLower       = `^0x[0-9a-f]{40}$`
	EthAddressUpper       = `^0x[0-9A-F]{40}$`
	FQDNRFC1123           = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // same as hostnameRFC1123 but must contain a non numerical TLD (possibly ending with '.')
	Hexadecimal           = "^(0[xX])?[0-9a-fA-F]+$"
	HexColor              = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	HostnameRFC1123       = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?$` // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	HostnameRFC952        = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                   // https://tools.ietf.org/html/rfc952
	HSL                   = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	HSLA                  = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	HTML                  = `<[/]?([a-zA-Z]+).*?>`
	HTMLEncoded           = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	ISBN10                = "^(?:[0-9]{9}X|[0-9]{10})$"
	ISBN13                = "^(?:(?:97(?:8|9))[0-9]{10})$"
	ISSN                  = "^(?:[0-9]{4}-[0-9]{3}[0-9X])$"
	JWT                   = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"
	Latitude              = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	Longitude             = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	MD4                   = "^[0-9a-f]{32}$"
	MD5                   = "^[0-9a-f]{32}$"
	MongoDB               = "^[a-f\\d]{24}$"
	Multibyte             = "[^\x00-\x7F]"
	Number                = "^[0-9]+$"
	Numeric               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	PrintableASCII        = "^[\x20-\x7E]*$"
	RGB                   = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	RGBA                  = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	RipeMD128             = "^[0-9a-f]{32}$"
	RipeMD160             = "^[0-9a-f]{40}$"
	SemVer                = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$` // numbered capture groups https://semver.org/
	Sha256                = "^[0-9a-f]{64}$"
	Sha384                = "^[0-9a-f]{96}$"
	Sha512                = "^[0-9a-f]{128}$"
	SpiceDBID             = `^(([a-zA-Z0-9/_|\-=+]{1,})|\*)$`
	SpiceDBPermission     = "^([a-z][a-z0-9_]{1,62}[a-z0-9])?$"
	SpiceDBType           = "^([a-z][a-z0-9_]{1,61}[a-z0-9]/)?[a-z][a-z0-9_]{1,62}[a-z0-9]$"
	SplitParams           = `'[^']*'|\S+`
	SSN                   = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	Tiger128              = "^[0-9a-f]{32}$"
	Tiger160              = "^[0-9a-f]{40}$"
	Tiger192              = "^[0-9a-f]{48}$"
	ULID                  = "^(?i)[A-HJKMNP-TV-Z0-9]{26}$"
	UnicodeAlpha          = "^[\\p{L}]+$"
	UnicodeAlphaNum       = "^[\\p{L}\\p{N}]+$"
	URLEncoded            = `^(?:[^%]|%[0-9A-Fa-f]{2})*$`
	UUID                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UUID3                 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UUID3RFC4122          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	UUID4                 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID4RFC4122          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	UUID5                 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID5RFC4122          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	UUIDRFC4122           = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
)
