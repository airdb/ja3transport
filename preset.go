package ja3transport

// Browser represents a browser JA3 and User-Agent string
type Browser struct {
	JA3       string
	UserAgent string
}

// ChromeAuto mocks Chrome 78
var ChromeAuto = Browser{
	JA3:       "769,47–53–5–10–49161–49162–49171–49172–50–56–19–4,0–10–11,23–24–25,0",
	UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36",
}

// SafariAuto mocks Safari 604.1
var SafariAuto = Browser{
	JA3:       "771,4865-4866-4867-49196-49195-49188-49187-49162-49161-52393-49200-49199-49192-49191-49172-49171-52392-157-156-61-60-53-47-49160-49170-10,65281-0-23-13-5-18-16-11-51-45-43-10-21,29-23-24-25,0",
	UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1",
}

var ChromeVersion103 = Browser{
	JA3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27,29-23-24,1",
	UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
}

var Ja3WithoutExtension = Browser{
	JA3:       "771,4865-4866-4867-49195-49196-52393-49199-49200-52392-49171-49172-156-157-47-53,,,",
	UserAgent: "without Extensions",
}

// Extension must contain 10 and 11 if group and format exist.
var Ja3WithPadding = Browser{
	JA3:       "771,4865-4866-4867-49195-49196-52393-49199-49200-53,0-51-43-13-45-10-11-21,29-23-30-25-24,0-1-2",
	UserAgent: "with padding",
}

// Extension must contain 10 and 11 if group and format exist.
var Ja3WithoutPadding = Browser{
	JA3:       "771,4865-4866-4867-49195-49196-52393-49199-49200-53,0-51-43-13-45-10-11,29-23-30-25-24,0-1-2",
	UserAgent: "without padding",
}

// Extension must contain 10 and 11 if group and format exist.
// So if extension 10 and 11 are not setting,  group and format are meaningless.
var Ja3WithoutGroupAndFormat = Browser{
	JA3:       "771,4865-4866-4867-49195-49196-52393-49199-49200-53,0-51-43-13-45,29-23-30-25-24,0-1-2",
	UserAgent: "without group and format",
}

var Ja3WithoutGroupAndFormatWrong = Browser{
	JA3:       "771,4865-4866-4867-49195-49196-52393-49199-49200-53,0-51-43-13-45-10-11,,",
	UserAgent: "without padding",
}

var BadJA3 = Browser{
	JA3:       "771,4865-4867-4866-49196-49162-49195-52393-49161-49200-49172-49199-52392-49171-157-53-156-47-10,0-23-65281-10-11-35-16-5-51-43-13-45-28",
	UserAgent: "curl/xx",
}
